package message

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context) {
	messageId := c.Param("id")

	message, code := model.GetMixinMessageById(convert.StrToNum(messageId))
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, message)
}
