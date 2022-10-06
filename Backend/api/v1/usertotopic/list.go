package usertotopic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
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

	var data []model.UserToTopic
	var code int
	var total int
	if r.UserId == "" && r.TopicId == "" {
		data, total, code = model.ListUserToTopics(r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
			return
		}
	} else if r.UserId != "" && r.TopicId == "" {
		data, total, code = model.ListUserToTopicsByUserId(r.UserId, r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
			return
		}
	} else if r.UserId == "" && r.TopicId != "" {
		data, total, code = model.ListUserToTopicsByTopicId(r.UserId, r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
			return
		}
	}
	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List:       data,
	})
}
