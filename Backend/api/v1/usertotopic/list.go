package usertotopic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type ListResponse struct {
	TotalCount int                 `json:"totalCount"`
	List       []model.UserToTopic `json:"list"`
}

type ListRequest struct {
	UserId  string `json:"user_id"`
	TopicId string `json:"topic_id"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
}

func ListUserToTopics(c *gin.Context) {
	var data []model.UserToTopic
	var code int
	var total int
	var usertotopic string
	var err error

	total, _ = v1.Redis().Get(v1.USERTOTOPIC_TOTAL).Int()
	usertotopic, err = v1.Redis().Get(v1.USERTOTOPIC_LIST).Result()
	convert.Unmarshal(usertotopic, &data)
	if err == redis.Nil {
		var r ListRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			v1.SendResponse(c, errmsg.ERROR_BIND, nil)
			return
		}
		switch {
		case r.Offset >= 100:
			r.Offset = 100
		case r.Limit <= 0:
			r.Limit = 10
		}

		if r.Limit == 0 {
			r.Limit = 10
		}

		data, total, code = model.ListUserToTopics(r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
			return
		}

		//
		usertotopic = convert.Marshal(&data)
		v1.Redis().Set(v1.USERTOTOPIC_TOTAL, total, v1.REDISEXPIRE)
		v1.Redis().Set(v1.USERTOTOPIC_LIST, usertotopic, v1.REDISEXPIRE)
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	} else if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	} else {
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	}
}

func ListUserToTopicsByUserId(c *gin.Context) {
	var data []model.UserToTopic
	var code int
	var total int
	var usertotopic string
	var err error
	userId := c.Param("userId")

	total, _ = v1.Redis().Get(v1.USERTOTOPIC_USER_TOTAL + userId).Int()
	usertotopic, err = v1.Redis().Get(v1.USERTOTOPIC_USER_LIST + userId).Result()
	convert.Unmarshal(usertotopic, &data)
	if err == redis.Nil {
		var r ListRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			v1.SendResponse(c, errmsg.ERROR_BIND, nil)
			return
		}
		switch {
		case r.Offset >= 100:
			r.Offset = 100
		case r.Limit <= 0:
			r.Limit = 10
		}

		if r.Limit == 0 {
			r.Limit = 10
		}

		data, total, code = model.ListUserToTopicsByUserId(r.UserId, r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
			return
		}

		//
		usertotopic = convert.Marshal(&data)
		v1.Redis().Set(v1.USERTOTOPIC_USER_TOTAL+userId, total, v1.REDISEXPIRE)
		v1.Redis().Set(v1.USERTOTOPIC_USER_LIST+userId, usertotopic, v1.REDISEXPIRE)
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	} else if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	} else {
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	}
}

func ListUserToTopicsByTopicId(c *gin.Context) {
	var data []model.UserToTopic
	var code int
	var total int
	var usertotopic string
	var err error
	topicId := c.Param("topicId")

	total, _ = v1.Redis().Get(v1.USERTOTOPIC_TOPIC_TOTAL + topicId).Int()
	usertotopic, err = v1.Redis().Get(v1.USERTOTOPIC_TOPIC_LIST + topicId).Result()
	convert.Unmarshal(usertotopic, &data)
	if err == redis.Nil {
		var r ListRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			v1.SendResponse(c, errmsg.ERROR_BIND, nil)
			return
		}
		switch {
		case r.Offset >= 100:
			r.Offset = 100
		case r.Limit <= 0:
			r.Limit = 10
		}

		if r.Limit == 0 {
			r.Limit = 10
		}

		data, total, code = model.ListUserToTopicsByTopicId(r.TopicId, r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
			return
		}

		//
		usertotopic = convert.Marshal(&data)
		v1.Redis().Set(v1.USERTOTOPIC_TOPIC_TOTAL+topicId, total, v1.REDISEXPIRE)
		v1.Redis().Set(v1.USERTOTOPIC_TOPIC_LIST+topicId, usertotopic, v1.REDISEXPIRE)
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	} else if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	} else {
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	}
}