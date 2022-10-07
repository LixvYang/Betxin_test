package usertotopic

import (
	"betxin/model"
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

// func ListUserToTopics(c *gin.Context) {
// 	var data []model.UserToTopic
// 	var code int
// 	var total int
// 	val, err := v1.Redis().RCGet("usertotopics").Result()

// 	if err == redis.Nil {
// 		var r ListRequest
// 		if err := c.ShouldBindJSON(&r); err != nil {
// 			v1.SendResponse(c, errmsg.ERROR_BIND, nil)
// 			return
// 		}
// 		switch {
// 		case r.Offset >= 100:
// 			r.Offset = 100
// 		case r.Limit <= 0:
// 			r.Limit = 10
// 		}

// 		if r.Limit == 0 {
// 			r.Limit = 10
// 		}

// 		if r.UserId == "" && r.TopicId == "" {
// 			data, total, code = model.ListUserToTopics(r.Offset, r.Limit)
// 			if code != errmsg.SUCCSE {
// 				v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
// 				return
// 			}
// 		} else if r.UserId != "" && r.TopicId == "" {
// 			data, total, code = model.ListUserToTopicsByUserId(r.UserId, r.Offset, r.Limit)
// 			if code != errmsg.SUCCSE {
// 				v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
// 				return
// 			}
// 		} else if r.UserId == "" && r.TopicId != "" {
// 			data, total, code = model.ListUserToTopicsByTopicId(r.UserId, r.Offset, r.Limit)
// 			if code != errmsg.SUCCSE {
// 				v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
// 				return
// 			}
// 		}
// 		v1.Redis().RCSet("usertotopics", data, 0)
// 		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
// 			TotalCount: total,
// 			List:       data,
// 		})
// 	} else if err != nil {
// 		v1.SendResponse(c, errmsg.ERROR, nil)
// 	} else {
// 		log.Println("Get usertotopic from Redis")
// 		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
// 			TotalCount: total,
// 		})
// 	}
// }
