package mixpay

import (
	"betxin/model"
	"betxin/service"
	"betxin/utils/errmsg"
	"context"
	"log"

	"github.com/shopspring/decimal"
)

func Worker(mixpayorder model.MixpayOrder, mixpayRes MixpayResult) error {
	// Mixpay支付成功
	// 将用户信息加入到user to topic 里面
	var data model.UserToTopic
	var selectWin string
	var err error
	var userTotalPrice decimal.Decimal
	data.UserId = mixpayorder.Uid
	data.Tid = mixpayorder.Tid

	if mixpayorder.YesRatio {
		selectWin = "yes_win"
		payAmount, _ := decimal.NewFromString(mixpayRes.Data.PaymentAmount)
		userTotalPrice, _ = service.CalculateTotalPriceBySymbol(context.Background(), mixpayRes.Data.PaymentSymbol, payAmount)
		data.YesRatioPrice = userTotalPrice
	} else {
		selectWin = "no_win"
		payAmount, _ := decimal.NewFromString(mixpayRes.Data.PaymentAmount)
		userTotalPrice, _ = service.CalculateTotalPriceBySymbol(context.Background(), mixpayRes.Data.PaymentSymbol, payAmount)
		data.NoRatioPrice = userTotalPrice
	}

	// 已经买过了
	if code := model.CheckUserToTopic(data.UserId, data.Tid); code != errmsg.ERROR {
		code = model.UpdateUserToTopic(&data)
		if code != errmsg.SUCCSE {
			log.Println("CreateUserToTopic错误")
			return err
		}
	} else {
		code = model.CreateUserToTopic(&data)
		if code != errmsg.SUCCSE {
			log.Println("CreateUserToTopic错误")
			return err
		}
	}

	if code := model.UpdateTopicTotalPrice(data.Tid, selectWin, userTotalPrice); code != errmsg.SUCCSE {
		log.Println("UpdateTopicTotalPrice错误")
		return err
	}
	return nil
}
