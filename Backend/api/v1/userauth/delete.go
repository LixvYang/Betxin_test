package userauth

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

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
func DeleteUserAUth(c *gin.Context) {
	userId := c.Param("userId")

	code := model.DeleteUserAuthorization(userId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
