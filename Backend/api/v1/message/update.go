package message

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func UpdateMessage(c *gin.Context) {
	var msg *model.MixinMessage
	msgId := c.Param("id")
	if err := c.ShouldBindJSON(&msg); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	code := model.UpdateMixinMessageByMsgId(msgId, msg)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_CATENAME, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, msgId)
}
