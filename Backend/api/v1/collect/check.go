package collect

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type CheckCollectRequest struct {
	UserId string `json:"user_id"`
	Tid    string `json:"tid"`
}

func CheckCollect(c *gin.Context) {
	var r CheckCollectRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	if code := model.CheckCollect(r.UserId, r.Tid); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, r.Tid)
}
