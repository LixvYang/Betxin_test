package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	TotalCount int           `json:"totalCount"`
	List       []model.Topic `json:"list"`
}

func ListTopics(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	data, total, code := model.ListTopics(pageSize, pageNum)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_LIST_TOPIC, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: int(total),
		List:       data,
	})
}
