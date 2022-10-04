package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func UpdateTopic(c *gin.Context) {
	var data model.Topic
	_ = c.ShouldBindJSON(&data)

	if code := model.UpdateTopic(&data); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_TOPIC, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, data)
}
