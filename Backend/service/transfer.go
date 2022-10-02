package service

import (
	"context"
	"log"
	"betxin/utils"
	"sort"

	fswap "github.com/fox-one/4swap-sdk-go"
	"github.com/fox-one/4swap-sdk-go/mtg"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

func Transfer(ctx context.Context, client *mixin.Client, AssetID string, OpponentID string, Amount decimal.Decimal, Memo string) (*mixin.Snapshot, error) {
	transferInput := &mixin.TransferInput{
		AssetID:    AssetID,
		OpponentID: OpponentID,
		Amount:     Amount,
		TraceID:    mixin.RandomTraceID(),
		Memo:       Memo,
	}
	tx, err := client.Transfer(ctx, transferInput, utils.Pin)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// 输入数量和输入资产id 输出交易单
func Transaction(ctx context.Context, client *mixin.Client, Amount decimal.Decimal, InputAssetID string) (*mixin.RawTransaction, error) {
	fswap.UseEndpoint(fswap.MtgEndpoint)
	// read the mtg group
	// the group information would change frequently
	// it's recommended to save it for later use
	group, err := fswap.ReadGroup(ctx)
	if err != nil {
		log.Println("读取组失败")
		return nil, err
	}
	pairs, _ := fswap.ListPairs(ctx)
	sort.Slice(pairs, func(i, j int) bool {
		aLiquidity := pairs[i].BaseValue.Add(pairs[i].QuoteValue)
		bLiquidity := pairs[j].BaseValue.Add(pairs[j].QuoteValue)
		return aLiquidity.GreaterThan(bLiquidity)
	})

	preOrder, err := fswap.Route(pairs, InputAssetID, utils.PUSD, Amount)
	if err != nil {
		log.Println("路由失败")
		return nil, err
	}

	followID, _ := uuid.NewV4()
	action := mtg.SwapAction(
		client.ClientID,
		followID.String(),
		utils.PUSD,
		preOrder.Routes,
		decimal.NewFromFloat(0.00000001),
	)

	// 生成 memo
	memo, err := action.Encode(group.PublicKey)
	if err != nil {
		log.Println("生成memo失败")
		return nil, err
	}

	tx, err := client.Transaction(ctx, &mixin.TransferInput{
		AssetID: InputAssetID,
		Amount:  Amount,
		TraceID: mixin.RandomTraceID(),
		Memo:    memo,
		OpponentMultisig: struct {
			Receivers []string `json:"receivers,omitempty"`
			Threshold uint8    `json:"threshold,omitempty"`
		}{
			Receivers: group.Members,
			Threshold: uint8(group.Threshold),
		},
	}, utils.Pin)
	if err != nil {
		log.Println("生成交易失败")
		return nil, err
	}
	return tx, nil
}

// 根据输入的资产id和资产数目计算出资产总价格
func CalculateTotalPriceByAssetId(ctx context.Context, AssedId string, amount decimal.Decimal)( decimal.Decimal, error) {
	decimal.DivisionPrecision = 2 // 保留两位小数，如有更多位，则进行四舍五入保留两位小数 
	asset, err := mixin.ReadNetworkAsset(ctx, AssedId)
	if err != nil {
		return decimal.NewFromFloat(0), err
	}
	return asset.PriceUSD.Mul(amount), nil
}