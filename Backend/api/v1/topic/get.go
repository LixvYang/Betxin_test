package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// GetArtInfo 查询单个话题信息
func GetTopicInfoById(c *gin.Context) {
	id := c.Param("id")
	var data model.Topic
	var code int
	var topic string
	var err error

	topic, err = v1.Redis().Get(v1.TOPIC_GET + id).Result()
	convert.Unmarshal(topic, &data)
	if err == redis.Nil {
		uuid, _ := uuid.FromString(id)
		fmt.Println(uuid)
		data, code = model.GetTopicById(uuid)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_GET_TOPIC, nil)
			return
		}
		topic = convert.Marshal(&data)
		v1.Redis().Set(v1.TOPIC_GET+id, topic, v1.REDISEXPIRE)
		v1.SendResponse(c, errmsg.SUCCSE, data)
		
	} else if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
	} else {
		v1.SendResponse(c, errmsg.SUCCSE, data)
	}
}
