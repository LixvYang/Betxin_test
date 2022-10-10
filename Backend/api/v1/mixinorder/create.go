package mixinorder

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func CreateMixinOrder(c *gin.Context) {
	var r model.MixinOrder
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	if code := model.CreateMixinOrder(&r); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	

	v1.SendResponse(c, errmsg.SUCCSE, r)
}
