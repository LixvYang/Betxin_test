package service

import (
	"betxin/model"
	"betxin/utils"
	"betxin/utils/errmsg"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/jasonlvhit/gocron"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Memo struct {
	Tid      string `json:"tid"`
	YesRatio bool   `json:"yes_ratio"`
	NoRatio  bool   `json:"no_ratio"`
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
		return
	}
	var wg sync.WaitGroup

	for _, snapshot := range snapshots {
		wg.Add(1)
		if snapshot.CreatedAt.After(preCreatedAt) {
			stats.updatePrevSnapshotCreatedAt(snapshot.CreatedAt)
			if snapshot.Amount.Cmp(decimal.NewFromInt(0)) == 1 && snapshot.Type == "transfer" {
				go func(ctx context.Context, client *mixin.Client, snapshot *mixin.Snapshot) {
					log.Println("又有新的订单了")
					defer wg.Done()
					HandlerNewMixinSnapshot(ctx, client, snapshot)
				}(ctx, client, snapshot)
			}
		}
	}
	wg.Wait()
}

func Worker(ctx context.Context, client *mixin.Client) error {
	// subclients := subclient.NewWorkderQueue(ctx, client)
	createdAt, err := getTopSnapshotCreatedAt(client, ctx)
	if err != nil {
		return nil
	}
	stats := &Stats{createdAt}
	gocron.Every(2).Second().Do(sendTopCreatedAtToChannel, ctx, stats, client)
	<-gocron.Start()
	return nil
}

func HandlerNewMixinSnapshot(ctx context.Context, client *mixin.Client, snapshot *mixin.Snapshot) error {
	r := model.MixinOrder{
		Type:       snapshot.Type,
		AssetId:    snapshot.AssetID,
		Amount:     snapshot.Amount,
		TraceId:    snapshot.TraceID,
		Memo:       snapshot.Memo,
		SnapshotId: snapshot.SnapshotID,
	}

	if code := model.CreateMixinOrder(&r); code != errmsg.SUCCSE {
		log.Println("创建CreateMixinOrder错误")
		return errors.New("")
	}

	// 用户传过来的memo是经过base64加密的  yes或no  再加上trace_id 的json
	///  memo  traceId:不应该是随机id 应该是把userid和买的topic id yesorno放在一起
	tx := &mixin.RawTransaction{}
	if snapshot.AssetID != utils.PUSD {
		tx = SwapOrderToPusd(ctx, client, snapshot.Amount, snapshot.AssetID, snapshot)
	} else {
		tx.Amount = snapshot.Amount.String()
		tx.AssetID = snapshot.AssetID
	}
	amount, err := decimal.NewFromString(tx.Amount)
	if err != nil {
		log.Println(err)
		log.Println("计算失败")
	}

	// 用户投入的总价格
	userTotalPrice, err := CalculateTotalPriceByAssetId(ctx, tx.AssetID, amount.Abs())
	if err != nil {
		log.Println("计算失败")
	}

	memoMsg, err := base64.StdEncoding.DecodeString(snapshot.Memo)
	if err != nil {
		return errors.New("解码memo失败")
	}
	memo := &Memo{}
	if err := json.Unmarshal(memoMsg, &memo); err != nil {
		return errors.New("解构memo失败")
	}

	var data model.UserToTopic
	var selectWin string
	data.UserId = snapshot.OpponentID
	if memo.YesRatio {
		selectWin = "yes_win"
		data.YesRatioPrice = userTotalPrice
	} else {
		selectWin = "no_win"
		data.NoRatioPrice = userTotalPrice
	}
	data.Tid = memo.Tid

	if code := model.CreateUserToTopic(&data); code != errmsg.SUCCSE {
		log.Println("CreateUserToTopic错误")
		return err
	}

	if code := model.UpdateTopicTotalPrice(data.Tid, selectWin, userTotalPrice); code != errmsg.SUCCSE {
		log.Println("UpdateTopicTotalPrice错误")
		return err
	}
	return nil
}

func SwapOrderToPusd(ctx context.Context, client *mixin.Client, Amount decimal.Decimal, InputAssetId string, snapshot *mixin.Snapshot) *mixin.RawTransaction {
	tx, err := TransactionWithRetry(ctx, client, Amount, InputAssetId)
	if err != nil {
		uuid := uuid.NewV4()
		model.CreateSendBack(&model.SendBack{TraceId: uuid.String()})
		err := TransferReturnWithRetry(ctx, client, uuid.String(), InputAssetId, snapshot.OpponentID, snapshot.Amount, "Swap 失败")
		switch {
		case mixin.IsErrorCodes(err, mixin.InsufficientBalance):
			log.Println("insufficient balance")
		default:
			log.Printf("transfer: %v", err)
		}
	}
	amount, _ := decimal.NewFromString(tx.Amount)
	data := &model.SwapOrder{
		Type:       tx.Type,
		SnapshotId: tx.SnapshotID,
		AssetID:    tx.AssetID,
		Amount:     amount,
		TraceId:    tx.TraceID,
		Memo:       tx.Memo,
		State:      tx.State,
	}

	if code := model.CreateSwapOrder(data); code != errmsg.SUCCSE {
		return nil
	}
	return tx
}
