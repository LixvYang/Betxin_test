package message

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func DeleteCollect(c *gin.Context) {
	messageId := c.Param("messageId")
	code := model.DeleteMixinMessageByMessageId(messageId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
