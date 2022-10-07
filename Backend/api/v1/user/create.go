package user

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

// @Summary 创建分类
// @Description 创建分类
// @Tags category
// @Accept  json
// @Produce  json
// @Param category body category.model true "创建分类"
// @Success 200 {object} v1.Response "{"code":200,"message":"OK","data":null}"
// @Router /v1/category/add [post]
func CreateUser(c *gin.Context) {
	var r model.User
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	code := model.CheckUser(r.UserId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_CATENAME_USED, nil)
		return
	}
	
	if code = model.CreateUser(&r); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.Redis().DelKeys(v1.USER_LIST, v1.USER_TOTAL)
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
