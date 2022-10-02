package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	TotalCount int               `json:"totalCount"`
	List       []model.Category `json:"list"`
}

func ListCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.ListCategories(pageSize, pageNum)
	v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
		TotalCount: total,
		List: data,
	})
}