package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	TotalCount int           `json:"totalCount"`
	List       []model.Topic `json:"list"`
}

type ListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Intro  string `json:"intro"`
	Cid    string `json:"cid"`
}

func ListTopics(c *gin.Context) {
	var r ListRequest
	var data []model.Topic
	var total int
	var code int
	var err error

	if err = c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	data, total, code = model.ListTopics(r.Offset, r.Limit)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List:       data,
	})
}

func ListTopicsNoLimit(c *gin.Context) {
	var r ListRequest
	var data []model.Topic
	var total int
	var code int
	var err error

	if err = c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	data, total, code = model.ListTopics(r.Offset, r.Limit)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List:       data,
	})
}

// GetTopicByCid 通过种类id获取信息
func GetTopicByCid(c *gin.Context) {
	// var topics string
	var data []model.Topic
	var total int
	var code int
	var err error

	cid := c.Param("cid")
	// total, _ = betxinredis.Get(v1.TOPIC_LIST_FROMCATE_TOTAL + cid).Int()
	// topics, err = betxinredis.Get(v1.TOPIC_LIST_FROMCATE + cid).Result()
	// convert.Unmarshal(topics, &data)
	// if err == redis.Nil {
	var r ListRequest
	if err = c.ShouldBindJSON(&r); err != nil {
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
	data, total, code = model.GetTopicByCid(convert.StrToNum(cid), r.Limit, r.Offset)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_GET_TOPIC, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List:       data,
	})
	// }
}

// GetTopicByCid 通过种类id获取信息
func GetTopicByTitle(c *gin.Context) {
	var data []model.Topic
	var total int
	var code int
	var err error

	var r ListRequest
	if err = c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	fmt.Println(r)

	switch {
	case r.Offset >= 100:
		r.Offset = 100
	case r.Limit <= 0:
		r.Limit = 10
	}

	if r.Limit == 0 {
		r.Limit = 10
	}
	data, total, code = model.SearchTopic(r.Offset, r.Limit, "intro LIKE  ?", "%"+r.Intro+"%")
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_GET_TOPIC, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List:       data,
	})
}
