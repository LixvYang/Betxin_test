package mixpayorder

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Tid      string `json:"tid"`
	YesRatio bool   `json:"yes_ratio"`
	NoRatio  bool   `json:"no_ratio"`
	OrderId string `json:"orderId"`
	PayeeId string `json:"payeeId"`
}

// 在用户点击时创建
func CreateMixinpayOrder(c *gin.Context) {
	var r CreateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	fmt.Println(r)

	data := &model.MixpayOrder{
		Tid: r.Tid,
		YesRatio: r.YesRatio,
		NoRatio: r.NoRatio,
		OrderId: r.OrderId,
		PayeeId: r.PayeeId,
	}

	if code := model.CreateMixpayOrder(data); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, r)
}