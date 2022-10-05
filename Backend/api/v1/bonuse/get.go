package bonuse

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBonuseByTraceId(c *gin.Context) {
	trace_id := c.Param("trace_id")
	bonuse, code := model.GetBonuseByTraceId(trace_id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, bonuse)
}

func GetBonuseById(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.Param("user_id"))

	bonuse, code := model.GetBonusesByUserId(user_id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	
	v1.SendResponse(c, errmsg.SUCCSE, bonuse)
}
