package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteTopic(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if code := model.DeleteTopic(id); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_DELETE_TOPIC, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, nil)
}