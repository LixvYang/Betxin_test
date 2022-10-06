package userauth

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"log"

	"github.com/gin-gonic/gin"
)

func UpdateUserAuth(c *gin.Context) {
	var userAuth *model.UserAuthorization
	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userAuth); err != nil {
		log.Panicln(err)
	}

	code := model.CheckUserAuthorization(userId)
	if code != errmsg.ERROR {
		v1.SendResponse(c, errmsg.ERROR_CATENAME_USED, nil)
		return
	}

	code = model.UpdateUserAuthorization(userId, userAuth.AccessToken)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_CATENAME, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, userId)
}
