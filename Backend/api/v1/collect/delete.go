package collect

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type DeleteRequest struct {
	UserId  string `json:"user_id"`
	TopicUuid uuid.UUID `json:"topic_uuid"`
}

func DeleteCollect(c *gin.Context) {
	var r DeleteRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	code := model.DeleteCollect(r.UserId, r.TopicUuid)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_DELETE_CATENAME, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
