package service

import (
	"betxin/model"
	"betxin/utils/errmsg"
	"context"
	"log"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/jasonlvhit/gocron"
	"github.com/shopspring/decimal"
)

type CreateRequest struct {
	Type       string          `json:"type"`
	AssetId    string          `json:"asset_id"`
	Amount     decimal.Decimal `gorm:"type: decimal(10, 10)" json:"amount"`
	TraceId    string          `json:"trace_id"`
	SnapshotId string          `json:"snapshot_id"`
	Memo       string          `json:"memo"`
}

type Memo struct {
	UserId   string `json:"user_id"`
	YesRatio bool   `json:"yes_ratio"`
	NoRatio  bool   `json:"no_ratio"`
	TraceId  string `json:"trace_id"`
}

type Stats struct {
	preCreatedAt time.Time
}

func (s *Stats) getPrevSnapshotCreatedAt() time.Time {
	return s.preCreatedAt
}

func (s *Stats) updatePrevSnapshotCreatedAt(time time.Time) {
	s.preCreatedAt = time
}

func getTopSnapshotCreatedAt(client *mixin.Client, c context.Context) (time.Time, error) {
	snapshots, err := client.ReadSnapshots(c, "", time.Now(), "", 1)
	if err != nil {
		return time.Now(), err
	}
	return snapshots[0].CreatedAt, nil
}

func getTopHundredCreated(client *mixin.Client, c context.Context) ([]*mixin.Snapshot, error) {
	snapshots, err := client.ReadSnapshots(c, "", time.Now(), "", 50)
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}

func sendTopCreatedAtToChannel(ctx context.Context, stats *Stats, client *mixin.Client) {
	preCreatedAt := stats.getPrevSnapshotCreatedAt()
	snapshots, err := getTopHundredCreated(client, ctx)
	if err != nil {
		log.Printf("getTopHundredCreated error")
		log.Printf(err.Error())
		return
	}
	for _, snapshot := range snapshots {
		if snapshot.CreatedAt.After(preCreatedAt) {
			log.Println("又有新的订单了")
			stats.updatePrevSnapshotCreatedAt(snapshot.CreatedAt)
			go HandlerNewMixinSnapshot(ctx, client, snapshot)
		}
	}
}

func Worker(ctx context.Context, client *mixin.Client) error {
	createdAt, err := getTopSnapshotCreatedAt(client, ctx)
	if err != nil {
		return nil
	}
	stats := &Stats{createdAt}
	gocron.Every(2).Second().Do(sendTopCreatedAtToChannel, ctx, stats, client)
	<-gocron.Start()
	return nil
}

func HandlerNewMixinSnapshot(ctx context.Context, client *mixin.Client, snapshot *mixin.Snapshot) {
	r := &CreateRequest{
		Type:       snapshot.Type,
		AssetId:    snapshot.AssetID,
		Amount:     snapshot.Amount,
		TraceId:    snapshot.TraceID,
		SnapshotId: snapshot.SnapshotID,
		Memo:       snapshot.Memo,
	}
	// 用户传过来的memo是经过base64加密的  yes或no  再加上trace_id 的json
	if code := CreateMixinOrder(r); code != errmsg.SUCCSE {
		log.Println("创建订单错误")
		return
	}

	// TODO:::: 解码memo后做
	// var tx *mixin.RawTransaction
	// if snapshot.AssetID != utils.PUSD {
	// 	tx = SwapOrderToPusd(ctx, client, snapshot.Amount, snapshot.AssetID, snapshot)
	// } else {
	// 	tx.Amount = snapshot.Amount.String()
	// 	tx.AssetID = snapshot.AssetID
	// }
	// amount, _ := decimal.NewFromString(tx.Amount)
	// // 用户投入的总价格
	// totalPrice, err := CalculateTotalPriceByAssetId(ctx, tx.AssetID, amount)
	// if err != nil {
	// 	log.Println("计算失败")
	// }

	// memoMsg, err := base64.StdEncoding.DecodeString(snapshot.Memo)
	// if err != nil {
	// 	log.Println("解码memo失败")
	// }
	// var memo Memo
	// if err := json.Unmarshal(memoMsg, &memo); err != nil {
	// 	log.Println("解构memo失败")
	// }
	// var data *model.UserToTopic
	// data.UserId = snapshot.OpponentID
	// if memo.YesRatio {
	// 	data.YesRatioPrice = totalPrice
	// } else {
	// 	data.NoRatioPrice = totalPrice
	// }

	// model.CreateUserToTopic(data)
}

func SwapOrderToPusd(ctx context.Context, client *mixin.Client, Amount decimal.Decimal, InputAssetId string, snapshot *mixin.Snapshot) *mixin.RawTransaction {
	tx, err := Transaction(ctx, client, Amount, InputAssetId)
	if err != nil {
		log.Println("swap交易失败")
		Transfer(ctx, client, InputAssetId, snapshot.OpponentID, snapshot.Amount, "交易失败")
		return nil
	}
	amount, _ := decimal.NewFromString(tx.Amount)
	date := &model.SwapOrder{
		Type:       tx.Type,
		SnapshotId: tx.SnapshotID,
		AssetID:    tx.AssetID,
		Amount:     amount,
		TraceId:    tx.TraceID,
		Memo:       tx.Memo,
		State:      tx.State,
	}

	// 加入数据库
	log.Println("加入数据库")
	if code := model.CreateSwapOrder(date); code != errmsg.SUCCSE {
		return nil
	}
	return tx
}

func CreateMixinOrder(r *CreateRequest) int {
	data := &model.MixinOrder{
		Type:       r.Type,
		AssetId:    r.AssetId,
		Amount:     r.Amount,
		TraceId:    r.TraceId,
		SnapshotId: r.SnapshotId,
		Memo:       r.Memo,
	}
	code := model.CreateMixinOrder(data)
	return code
}
