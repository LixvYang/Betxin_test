package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"
	betxinredis "betxin/utils/redis"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateTopic(c *gin.Context) {
	tid := c.Param("id")
	var r CreateReqeust
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
	}

	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", r.EndTime, time.Local)
	if err != nil {
		log.Println(err)
	}
	
	topic := &model.Topic{
		Tid:     r.Tid,
		Cid:     convert.StrToNum(r.Cid),
		Title:   r.Title,
		Intro:   r.Intro,
		ImgUrl:  r.ImgUrl,
		EndTime: endTime,
	}

	if code := model.UpdateTopic(tid, topic); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_TOPIC, nil)
		return
	}

	// Delete redis store
	betxinredis.DelKeys(v1.TOPIC_TOTAL, v1.TOPIC_LIST, v1.TOPIC_GET+tid)

	v1.SendResponse(c, errmsg.SUCCSE, tid)
}
