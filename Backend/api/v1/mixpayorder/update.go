package mixpayorder

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/service/mixpay"
	"betxin/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	OrderId string `json:"orderId"`
	PayeeId string `json:"payeeId"`
	TraceId string `json:"traceId"`
}

func UpdateMixpayOrder(c *gin.Context) {
	var mixpayorder model.MixpayOrder
	var u UpdateRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	mixpayorder = model.MixpayOrder{
		OrderId: u.OrderId,
		PayeeId: u.PayeeId,
		TraceId: u.TraceId,
	}

	if code := model.UpdateMixpayOrder(&mixpayorder); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	mixpayRes, err := mixpay.GetMixpayResult(u.TraceId)
	if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	// 查询Mixpay支付信息　比如
	mixpayorder, code := model.GetMixpayOrder(u.TraceId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	if err := mixpay.Worker(mixpayorder, mixpayRes); err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "SUCCESS",
	})
}
