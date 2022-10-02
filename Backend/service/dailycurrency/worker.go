package dailycurrency

import (
	"context"
	"fmt"
	"log"
	"betxin/model"
	"betxin/utils"
	"betxin/utils/errmsg"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/jasonlvhit/gocron"
)

func updateDailyCurrency(ctx context.Context, update bool) {

	AllCurrency := [...]string{
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
	for _, currency := range AllCurrency {
		asset, err := mixin.ReadNetworkAsset(ctx, currency)
		if err != nil {
			return
		}
		data := model.Currency{
			AssetId:  asset.AssetID,
			PriceUsd: asset.PriceUSD,
			PriceBtc: asset.PriceBTC,
			ChainId:  asset.ChainID,
			IconUrl:  asset.IconURL,
			Symbol:   asset.Symbol,
		}
		if update {
			if code := model.UpdateCurrency(&data); code != errmsg.SUCCSE {
				log.Fatalln(code)
			}
		} else {
			if code := model.CreateCurrency(&data); code != errmsg.SUCCSE {
				log.Fatalln(code)
			}
		}
	}
}

func DailyCurrency(ctx context.Context) {
	fmt.Println("调用")
	updateDailyCurrency(ctx, false)
	gocron.Every(1).Day().Do(updateDailyCurrency, ctx, true)
}
