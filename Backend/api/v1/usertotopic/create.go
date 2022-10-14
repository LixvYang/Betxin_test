package usertotopic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func CreateUserToTopic(c *gin.Context) {
	var r model.UserToTopic
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	if code := model.CreateUserToTopic(&r); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	// betxinredis.DelKeys(
	// 	v1.USERTOTOPIC_LIST,
	// 	v1.USERTOTOPIC_TOTAL,
	// 	v1.USERTOTOPIC_TOPIC_TOTAL+r.Tid,
	// 	v1.USERTOTOPIC_TOPIC_LIST+r.Tid,
	// 	v1.USERTOTOPIC_USER_LIST+r.UserId,
	// 	v1.USERTOTOPIC_USER_TOTAL+r.UserId,
	// )
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
