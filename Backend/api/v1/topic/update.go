package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func UpdateTopic(c *gin.Context) {
	id := c.Param("id")
	var topic model.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
	}

	uuid, _ := uuid.FromString(id)
	if code := model.UpdateTopic(uuid, &topic); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_TOPIC, nil)
		return
	}

	// Delete redis store
	v1.Redis().DelKeys(v1.TOPIC_TOTAL, v1.TOPIC_LIST, v1.TOPIC_GET+id)

	v1.SendResponse(c, errmsg.SUCCSE, id)
}
