package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 根据id删除分类
// @Description 根据id删除分类
// @Tags	category
// @Accept  json
// @Produce  json
// @Param id path int true "分类的数据库id"
// @Success 200 {object} v1.Response "{"code":200,"message":"OK","data":null}"
// @Router /v1/category/{id} [delete]
func DeleteCategory(c *gin.Context) {
	CategoryId, _ := strconv.Atoi(c.Param("id"))
	
	code := model.DeleteCategory(CategoryId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_DELETE_CATENAME, nil)
		return
	}

	// Delete redis store 
	if v1.Redis().Exists("categoryies") || v1.Redis().Exists("categoryiesTotal") {
		v1.Redis().Del("categoryiesTotal")
		v1.Redis().Del("categoryies")
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
