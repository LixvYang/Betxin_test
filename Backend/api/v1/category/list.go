package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type ListResponse struct {
	TotalCount int              `json:"totalCount"`
	List       []model.Category `json:"list"`
}

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// @Summary 获得category列表
// @Description 获取category列表
// @Tags category
// @Accept  json
// @Produce  json
// @Param   offset      query    int     true     "Offset"
// @Param   limit      query    int     true      "Limit"
// @Success 200 {object} category.ListResponse "{"code":200,"message":"OK","data":{"totalCount":1,"list":[]}"
// @Router /v1/category [get]
func ListCategories(c *gin.Context) {
	var categoryies string
	var total int
	var code int
	var err error
	var data []model.Category

	total, _ = v1.Redis().Get("categoryiesTotal").Int()
	categoryies, err = v1.Redis().Get("categoryies").Result()
	_ = json.Unmarshal([]byte(categoryies), &data)
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

		data, total, code = model.ListCategories(r.Offset, r.Limit)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR_LIST_CATEGORY, nil)
			return
		}
		categoryies, _ := json.Marshal(data)
		fmt.Println("设值")
		v1.Redis().Set("categoryiesTotal", total, time.Hour*12)
		v1.Redis().Set("categoryies", categoryies, time.Hour*12)
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	} else if err != nil {
		log.Panicln(err)
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	} else {
		fmt.Println("从redis拿数据")
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	}
}
