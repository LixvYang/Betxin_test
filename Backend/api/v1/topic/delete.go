package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func DeleteTopic(c *gin.Context) {
	id := c.Param("id")

	if code := model.DeleteTopic(convert.StrToNum(id)); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_DELETE_TOPIC, nil)
		return
	}

	v1.Redis().DelKeys(v1.TOPIC_LIST, v1.TOPIC_TOTAL, v1.TOPIC_GET+id)
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
