package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/service"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type StopRequest struct {
	Tid    string `json:"tid"`
	YesWin bool   `json:"yes_win"`
	NoWin  bool   `json:"no_win"`
}

func StopTopic(c *gin.Context) {
	var r StopRequest
	var win string

	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	if r.YesWin {
		win = "yes_win"
	} else {
		win = "no_win"
	}
	service.EndOfTopic(c, r.Tid, win)

	if code := model.StopTopic(r.Tid); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, r.Tid)
}
