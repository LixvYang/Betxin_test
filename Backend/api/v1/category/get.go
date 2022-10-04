package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	v1.SendResponse(c, errmsg.SUCCSE, category)
}
