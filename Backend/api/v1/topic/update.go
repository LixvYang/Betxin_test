package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	betxinredis "betxin/utils/redis"

	"github.com/gin-gonic/gin"
)

func UpdateTopic(c *gin.Context) {
	tid := c.Param("tid")
	var topic model.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
	}

	if code := model.UpdateTopic(tid, &topic); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_TOPIC, nil)
		return
	}

	// Delete redis store
	betxinredis.DelKeys(v1.TOPIC_TOTAL, v1.TOPIC_LIST, v1.TOPIC_GET+tid)

	v1.SendResponse(c, errmsg.SUCCSE, tid)
}
