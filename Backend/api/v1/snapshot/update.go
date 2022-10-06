package snapshot

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateMixinNetworkSnapshot(c *gin.Context) {
	var mixinNetworkSnapshot *model.MixinNetworkSnapshot
	traceId := c.Param("traceId")
	if err := c.ShouldBindJSON(&mixinNetworkSnapshot); err != nil {
		log.Panicln(err)
	}

	code := model.UpdateMixinNetworkSnapshot(traceId, mixinNetworkSnapshot)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_CATENAME, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
