package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategoryInfo(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))
	category, code := model.GetCategoryById(categoryId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, &model.Category{
		Id:           category.Id,
		CategoryName: category.CategoryName,
	})
}
