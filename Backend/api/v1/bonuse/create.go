package bonuse

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type CreateRequest struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AssetId     string `json:"asset_id"`
	Amount      decimal.Decimal    `json:"amount"`
	Memo        string `json:"memo"`
	TraceId     string `json:"trace_id"`
}
type CreateResponse struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AssetId     string `json:"asset_id"`
	Amount      decimal.Decimal    `json:"amount"`
	Memo        string `json:"memo"`
	TraceId     string `json:"trace_id"`
}

func CreateBonuse(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	data := &model.Bonuse{
		UserId:      r.UserId,
		Title:       r.Title,
		Description: r.Description,
		AssetId:     r.AssetId,
		Amount:      r.Amount,
		TraceId:     r.TraceId,
	}

	code := model.CheckBonuse(data.TraceId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_BONUSE_EXIST, nil)
		return
	}

	model.CreateBonuse(data)
	rsp := CreateResponse(r)
	v1.SendResponse(c, errmsg.SUCCSE, rsp)
}
