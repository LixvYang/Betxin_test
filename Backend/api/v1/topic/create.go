package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func CreateTopic(c *gin.Context) {
	var data model.Topic
	if err := c.ShouldBindJSON(&data); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	code := model.CreateTopic(&data)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.Redis().DelKeys(v1.TOPIC_TOTAL, v1.TOPIC_LIST)

	v1.SendResponse(c, errmsg.SUCCSE, data)
}
