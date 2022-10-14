package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	TotalCount int           `json:"totalCount"`
	List       []model.Topic `json:"list"`
}

type ListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Title  string `json:"title"`
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

	if r.Title == "" && r.Cid == "" {
		data, total, code = model.ListTopics(r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
			return
		}
	} else if r.Title != "" {
		data, total, code = model.SearchTopic(r.Offset, r.Limit, "title LIKE ?", r.Title+"%")
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
		}
	} else if r.Cid != "" {
		data, total, code = model.SearchTopic(r.Offset, r.Limit, "cid = ?", r.Cid)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
		}
	}

	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List:       data,
	})
}

// func ListTopics(c *gin.Context) {
// 	var topics string
// 	var r ListRequest
// 	var data []model.Topic
// 	var total int
// 	var code int
// 	var err error

// 	total, _ = betxinredis.Get(v1.TOPIC_TOTAL).Int()
// 	topics, err = betxinredis.Get(v1.TOPIC_LIST).Result()
// 	convert.Unmarshal(topics, &data)
// 	if err == redis.Nil {
// 		if err = c.ShouldBindJSON(&r); err != nil {
// 			v1.SendResponse(c, errmsg.ERROR_BIND, nil)
// 			return
// 		}
// 		if r.Title == "" && r.Content == "" && r.Intro == "" {
// 			data, total, code = model.ListTopics(r.Offset, r.Limit)
// 			if code != errmsg.SUCCSE {
// 				v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
// 				return
// 			}

// 			//
// 			topics = convert.Marshal(&data)
// 			fmt.Println("设值")
// 			betxinredis.Set(v1.TOPIC_TOTAL, total, v1.REDISEXPIRE)
// 			betxinredis.Set(v1.TOPIC_LIST, topics, v1.REDISEXPIRE)
// 		} else {
// 			data, total, code = model.SearchTopic(r.Title, r.Content, r.Intro, r.Offset, r.Limit)
// 			if code != errmsg.SUCCSE {
// 				v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
// 				return
// 			}
// 		}

// 		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
// 			TotalCount: total,
// 			List:       data,
// 		})
// 	} else if err != nil {
// 		v1.SendResponse(c, errmsg.ERROR, nil)
// 		return
// 	} else {
// 		fmt.Println("从redis拿数据")
// 		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
// 			TotalCount: total,
// 			List:       data,
// 		})
// 	}
// }

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

	//
	// topics = convert.Marshal(&data)
	// betxinredis.Set(v1.TOPIC_LIST_FROMCATE_TOTAL+cid, total, v1.REDISEXPIRE)
	// betxinredis.Set(v1.TOPIC_LIST_FROMCATE+cid, topics+cid, v1.REDISEXPIRE)

	// v1.SendResponse(c, errmsg.SUCCSE, &ListResponse{
	// 	TotalCount: total,
	// 	List:       data,
	// })
	// } else if err != nil {
	// 	v1.SendResponse(c, errmsg.ERROR, nil)
	// 	return
	// } else {
	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List:       data,
	})
	// }
}
