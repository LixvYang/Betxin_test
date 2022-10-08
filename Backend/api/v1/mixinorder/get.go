package mixinorder

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func GetMixinOrderById(c *gin.Context) {
	traceId := c.Param("traceId")
	mixinOrder, code := model.GetMixinOrderByTraceId(traceId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, mixinOrder)
}
