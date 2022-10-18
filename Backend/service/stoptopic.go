package service

import (
	"betxin/model"
	"betxin/utils"
	"betxin/utils/errmsg"
	"context"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type UserBounse struct {
	percentage decimal.Decimal
	TraceId    string
	UserId     string
	Memo       string
}

func EndOfTopic(c context.Context, tid string, win string) {
	var code int
	var userTotopics []model.UserToTopic
	var totalPrice int64
	var userBounses []UserBounse
	data := &model.Bonuse{}

	totalPrice, code = model.GetTopicTotalPrice(tid)
	if code != errmsg.SUCCSE {
		return
	}

	userTotopics, _, code = model.ListUserToTopicsWin(tid, win)
	if code != errmsg.SUCCSE {
		return
	}

	for _, userToTopic := range userTotopics {
		data.Tid = tid
		if win == "yes_win" {
			data.Amount = userToTopic.YesRatioPrice.Div(decimal.NewFromInt(totalPrice))
		} else {
			data.Amount = userToTopic.NoRatioPrice.Div(decimal.NewFromInt(totalPrice))
		}
		data.AssetId = utils.PUSD
		data.Memo = fmt.Sprintln("bonuse from betxin" + userToTopic.Topic.Intro)
		data.UserId = userToTopic.UserId
		data.TraceId = uuid.NewV4().String()
		userBounses = append(userBounses, UserBounse{percentage: data.Amount, UserId: data.UserId, TraceId: data.TraceId, Memo: data.Memo})
		if code = model.CreateBonuse(data); code != errmsg.SUCCSE {
			log.Println("创建奖金出错")
			return
		}
		snapShot := &model.MixinNetworkSnapshot{
			TraceId: data.TraceId,
		}
		model.CreateMixinNetworkSnapshot(snapShot)
	}

	// send for users
	for _, userBounse := range userBounses {
		TransferWithRetry(c, mixinClient, userBounse.TraceId, utils.PUSD, userBounse.UserId, userBounse.percentage.Mul(decimal.NewFromInt(int64(totalPrice))), userBounse.Memo)

	}
}
