package service

import (
	"betxin/model"
	"betxin/utils"
	"betxin/utils/errmsg"
	"context"
	"sync"

	"github.com/shopspring/decimal"
)

func EndOfTopic(c context.Context, tid string, win string) {
	type UserBounse struct {
		percentage decimal.Decimal
		UserId     string
	}

	var code int
	var userTotopics []model.UserToTopic
	var totalPrice int
	var userBounses []UserBounse
	var wg sync.WaitGroup
	var data model.Bonuse

	totalPrice, code = model.GetTopicTotalPrice(tid)
	if code != errmsg.SUCCSE {
		return
	}

	// add win user to userTotopics
	userTotopics, _, _ = model.ListUserToTopicsWin(tid, win)
	for _, userToTopic := range userTotopics {
		wg.Add(1)
		go func(userToTopic model.UserToTopic) {
			defer wg.Done()
			data.Amount = userToTopic.YesRatioPrice.Div(decimal.NewFromInt(int64(totalPrice)))
			data.AssetId = utils.PUSD
			data.Memo = "bounse from betxin"
			data.UserId = userToTopic.UserId
			if win == "yes_win" {
				userBounses = append(userBounses, UserBounse{percentage: data.Amount, UserId: data.UserId})
			} else {
				userBounses = append(userBounses, UserBounse{percentage: data.Amount, UserId: data.UserId})
			}
			model.CreateBonuse(&data)
		}(userToTopic)
	}
	wg.Wait()

	// 给用户转账
	for _, userBounse := range userBounses {
		wg.Add(1)
		go func(c context.Context, userBounse UserBounse) {
			defer wg.Done()
			Transfer(c, mixinClient, utils.PUSD, userBounse.UserId, userBounse.percentage.Mul(decimal.NewFromInt(int64(totalPrice))), "")
		}(c, userBounse)
	}
	wg.Wait()
}
