package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func StopTopic(c *gin.Context) {
	id := c.Param("id")
	uuid, _ := uuid.FromString(id)
	fmt.Println(uuid)
	if code := model.StopTopic(uuid); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, id)
}
