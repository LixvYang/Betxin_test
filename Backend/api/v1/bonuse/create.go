package bonuse

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func CreateBonuse(c *gin.Context) {
	var r model.Bonuse
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	
	if code := model.CreateBonuse(&r); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_CREATE_BONUSE, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, r)
}
