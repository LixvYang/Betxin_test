package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 用分类id获取分类
// @Description 用分类id获取分类信息
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.Category "{"code":200,"message":"OK","data":{}}"
// @Router /v1/category/{id} [get]
func GetCategoryInfo(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("获取参数id错误")
	}

	category, code := model.GetCategoryById(categoryId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	a, _ := json.Marshal(category)
	v1.Redis().Set("category"+c.Param("id"), string(a), time.Hour*2)

	v1.SendResponse(c, errmsg.SUCCSE, category)
}
