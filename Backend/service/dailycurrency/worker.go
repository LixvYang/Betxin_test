package dailycurrency

import (
	"betxin/utils"
	betxinredis "betxin/utils/redis"
	"context"
	"sync"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/jasonlvhit/gocron"
)

var AllCurrency = [...]string{
	utils.PUSD,
	utils.BTC,
	utils.BOX,
	utils.XIN,
	utils.ETH,
	utils.MOB,
	utils.USDC,
	utils.USDT,
	utils.EOS,
	utils.SOL,
	utils.UNI,
	utils.DOGE,
	utils.RUM,
	utils.DOT,
	utils.WOO,
	utils.ZEC,
	utils.LTC,
	utils.SHIB,
	utils.BCH,
	utils.MANA,
	utils.FIL,
	utils.BNB,
	utils.XRP,
	utils.SC,
	utils.MATIC,
	utils.ETC,
	utils.XMR,
	utils.DCR,
	utils.TRX,
	utils.ATOM,
	utils.CKB,
	utils.LINK,
	utils.GTC,
	utils.HNS,
	utils.DASH,
	utils.XLM,
}

func updateRedisCurrency(ctx context.Context) {
	var wg sync.WaitGroup

	for _, currency := range AllCurrency {
		wg.Add(1)
		go func(currency string) {
			defer wg.Done()
			asset, err := mixin.ReadNetworkAsset(ctx, currency)
			if err != nil {
				return
			}
			betxinredis.Del(asset.Name + "_" + currency + "_" + "price")
			betxinredis.Set(asset.Name+"_"+currency+"_"+"price", asset.PriceUSD, time.Minute)
		}(currency)
	}
	wg.Wait()
}

func DailyCurrency(ctx context.Context) {
	updateRedisCurrency(ctx)
	gocron.Every(1).Minute().Do(updateRedisCurrency, ctx)
}
