package service

import (
	"betxin/model"
	"betxin/utils"
	"betxin/utils/errmsg"
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type UserBounse struct {
	percentage decimal.Decimal
	UserId     string
}

func EndOfTopic(c context.Context, tid string, win string) {
	// var wg sync.WaitGroup
	var code int
	var userTotopics []model.UserToTopic
	var totalPrice int
	// var userBounses []UserBounse
	// var wg sync.WaitGroup
	data := &model.Bonuse{}

	totalPrice, code = model.GetTopicTotalPrice(tid)
	if code != errmsg.SUCCSE {
		return
	}
	userTotopics, _, code = model.ListUserToTopicsWin(tid, win)
	if code != errmsg.SUCCSE {
		return
	}

	for i := 0; i < len(userTotopics); i++ {
		// wg.Add(1)
		// go func(userToTopic model.UserToTopic) {
		// 	defer wg.Done()
		// data.Tid = userTotopics[i].Tid
		data.Amount = userTotopics[i].YesRatioPrice.Div(decimal.NewFromInt(int64(totalPrice)))
		data.AssetId = utils.PUSD
		data.Memo = "bounse from betxin"
		data.UserId = userTotopics[i].UserId
		data.TraceId = uuid.NewV4().String()
		// if win == "yes_win" {
		// 	userBounses = append(userBounses, UserBounse{percentage: data.Amount, UserId: data.UserId})
		// } else {
		// 	userBounses = append(userBounses, UserBounse{percentage: data.Amount, UserId: data.UserId})
		// }
		if code = model.CreateBonuse(data); code != errmsg.SUCCSE {
			log.Println("创建奖金出错")
			return
		}

		// 	}(userToTopic)
	}

	// wg.Wait()

	// 给用户转账
	// for _, userBounse := range userBounses {
		// wg.Add(1)
		// go func(c context.Context, userBounse UserBounse) {
		// defer wg.Done()
		// Transfer(c, mixinClient, utils.PUSD, userBounse.UserId, userBounse.percentage.Mul(decimal.NewFromInt(int64(totalPrice))), "")
		// }(c, userBounse)
	// }
	// wg.Wait()
}
