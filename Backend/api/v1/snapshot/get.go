package snapshot

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func GetMixinNetworkSnapshot(c *gin.Context) {
	trace_id := c.Param("trace_id")

	if code := model.CheckMixinNetworkSnapshot(trace_id); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	mixinNetworkSnapshot, code := model.GetMixinNetworkSnapshot(trace_id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, mixinNetworkSnapshot)
}
