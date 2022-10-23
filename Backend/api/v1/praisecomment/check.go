package praisecomment

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type CheckPraiseRequest struct {
	Cid  int    `json:"cid"`
	Uid string `json:"uid"`
}

func CheckPraise(c *gin.Context) {
	var r CheckPraiseRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	if code := model.CheckPraiseComment(r.Cid, r.Uid); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
