package snapshot

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func GetMixinNetworkSnapshot(c *gin.Context) {
	traceId := c.Param("traceId")

	if code := model.CheckMixinNetworkSnapshot(traceId); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	mixinNetworkSnapshot, code := model.GetMixinNetworkSnapshot(traceId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, mixinNetworkSnapshot)
}
