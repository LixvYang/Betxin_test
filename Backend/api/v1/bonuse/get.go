package bonuse

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func GetBonuse(c *gin.Context) {
	trace_id := c.Param("trace_id")
	bonuse, code := model.GetBonuseByTraceId(trace_id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, bonuse)
}
