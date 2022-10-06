package user

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var user *model.User
	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Panicln(err)
	}
	code := model.CheckUser(userId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	code = model.UpdateUser(userId, user)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, userId)
}
