package collect

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type GetRequest struct {
	UserId string `json:"user_id"`
}

func GetCollect(c *gin.Context) {
	var r GetRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	category, total, code := model.GetCollectByUserId(r.UserId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List: category,
	})
}
