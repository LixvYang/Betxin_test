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

	v1.Redis().DelKeys(
		v1.USERTOTOPIC_LIST,
		v1.USERTOTOPIC_TOTAL,
		v1.USERTOTOPIC_TOPIC_TOTAL+r.TopicUuid,
		v1.USERTOTOPIC_TOPIC_LIST+r.TopicUuid,
		v1.USERTOTOPIC_USER_LIST+r.UserId,
		v1.USERTOTOPIC_USER_TOTAL+r.UserId,
	)
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
