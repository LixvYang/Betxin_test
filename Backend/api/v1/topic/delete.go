package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	betxinredis "betxin/utils/redis"

	"github.com/gin-gonic/gin"
)

func DeleteTopic(c *gin.Context) {
	tid := c.Param("id")

	if code := model.DeleteTopic(tid); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_DELETE_TOPIC, nil)
		return
	}

	betxinredis.DelKeys(v1.TOPIC_LIST, v1.TOPIC_TOTAL, v1.TOPIC_GET+tid)
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
