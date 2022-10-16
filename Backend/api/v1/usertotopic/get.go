package usertotopic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserToTopic(c *gin.Context) {
	tid := c.Param("id")
	session := sessions.Default(c)
	user := session.Get("userId")
	userId := fmt.Sprintf("%v", user)
	data, code := model.GetUserToTopic(userId, tid)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, data)
}
