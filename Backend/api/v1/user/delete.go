package user

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	code := model.DeleteUser(userId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_DELETE_CATENAME, nil)
		return
	}
	v1.Redis().DelKeys(v1.USER_INFO+userId, v1.USER_LIST, v1.USER_TOTAL)
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
