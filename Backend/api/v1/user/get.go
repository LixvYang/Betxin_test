package user

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func GetUserInfoByUserId(c *gin.Context) {
	userId := c.Param("userId")

	user, code := model.GetUserById(userId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, user)
}

func GetUserInfoByUserFullName(c *gin.Context) {
	fullName := c.Param("fullName")

	user, code := model.GetUserByName(fullName)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, user)
}
