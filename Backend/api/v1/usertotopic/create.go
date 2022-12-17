package usertotopic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"sync"

	"github.com/gin-gonic/gin"
)

var userToTopicPool = sync.Pool{
	New: func() any {
		return new(model.UserToTopic)
	},
}

func CreateUserToTopic(c *gin.Context) {
	// var r model.UserToTopic
	r := userToTopicPool.Get().(*model.UserToTopic)
	if err := c.ShouldBindJSON(r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	if code := model.CreateUserToTopic(r); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	userToTopicPool.Put(r)
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
