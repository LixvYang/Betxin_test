package swaporder

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func UpdateMessage(c *gin.Context) {
	var swapOrder *model.MixinOrder
	traceId := c.Param("traceId")
	if err := c.ShouldBindJSON(&swapOrder); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	code := model.UpdateMixinOrder(traceId, swapOrder)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_CATENAME, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, traceId)
}
