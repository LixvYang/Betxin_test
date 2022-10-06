package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	TotalCount int           `json:"totalCount"`
	List       []model.Topic `json:"list"`
}

type TitleListRequest struct {
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Intro   string `json:"intro"`
}

func ListTopics(c *gin.Context) {
	var r TitleListRequest
	var data []model.Topic
	var total int
	var code int
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	fmt.Println(r)
	if r.Title == "" && r.Content == "" && r.Intro == "" {
		data, total, code = model.ListTopics(r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
			return
		}
	} else {
		data, total, code = model.SearchTopic(r.Title, r.Content, r.Intro, r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
			return
		}
	}
	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: int(total),
		List:       data,
	})
}

// GetTopicByCid 通过种类id获取信息
func GetTopicByCid(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
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
	data, total, code := model.GetTopicByCid(cid, r.Limit, r.Offset)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_GET_TOPIC, nil)
		return
	}
	fmt.Println(data)
	v1.SendResponse(c, errmsg.SUCCSE, &ListResponse{
		TotalCount: total,
		List:       data,
	})
}
