package usertotopic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type DeleteUserToTopicRequest struct {
	UserId    string `json:"user_id"`
	TopicUuid string `json:"topic_uuid"`
}

func DeleteUserToTopic(c *gin.Context) {
	var d DeleteUserToTopicRequest
	if err := c.ShouldBindJSON(&d); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	code := model.DeleteUserToTopic(d.UserId, d.TopicUuid)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_DELETE_CATENAME, nil)
		return
	}

	v1.Redis().DelKeys(
		v1.USERTOTOPIC_LIST,
		v1.USERTOTOPIC_TOTAL,
		v1.USERTOTOPIC_TOPIC_TOTAL+d.TopicUuid,
		v1.USERTOTOPIC_TOPIC_LIST+d.TopicUuid,
		v1.USERTOTOPIC_USER_LIST+d.UserId,
		v1.USERTOTOPIC_USER_TOTAL+d.UserId,
	)
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
