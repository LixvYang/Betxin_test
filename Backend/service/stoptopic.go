package service

import (
	"betxin/model"
	"betxin/utils"
	"betxin/utils/errmsg"
	"context"
	"fmt"
	"log"
	"time"

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
	var totalPrice decimal.Decimal
	var winTotalPrice decimal.Decimal
	var userBounses []UserBounse
	data := &model.Bonuse{}

	totalPrice, code = model.GetTopicTotalPrice(tid)
	if code != errmsg.SUCCSE {
		return
	}
	log.Println(totalPrice)

	userTotopics, _, code = model.ListUserToTopicsWin(tid, win)
	if code != errmsg.SUCCSE {
		log.Println("列出赢了的用户失败")
		return
	}

	for _, userToTopic := range userTotopics {
		data.Tid = tid
		data.AssetId = utils.PUSD
		data.Memo = fmt.Sprintln("bonuse from betxin" + userToTopic.Topic.Intro)
		data.UserId = userToTopic.UserId
		data.TraceId = uuid.NewV4().String()
		winTotalPrice, code = model.SearchTopicWinTopic(tid, win)
		if code != errmsg.SUCCSE {
			log.Println("计算赢了总价格失败")
		}

		if win == "yes_win" {
			// 占赢了的百分比
			percentage := userToTopic.YesRatioPrice.Div(winTotalPrice)
			log.Println("percentage: ", percentage)
			data.Amount = percentage.Mul(totalPrice)
			userBounses = append(userBounses, UserBounse{percentage: percentage, UserId: data.UserId, TraceId: data.TraceId, Memo: data.Memo})
		} else {
			percentage := userToTopic.NoRatioPrice.Div(winTotalPrice)
			data.Amount = percentage.Mul(totalPrice)
			userBounses = append(userBounses, UserBounse{percentage: percentage, UserId: data.UserId, TraceId: data.TraceId, Memo: data.Memo})
		}

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
		fmt.Println("转账给用户")
		TransferWithRetry(c, mixinClient, userBounse.TraceId, utils.PUSD, userBounse.UserId, userBounse.percentage.Mul(totalPrice), userBounse.Memo)
		time.Sleep(1 * time.Second)
	}
}
