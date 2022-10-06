package bonuse

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func DeleteBonuse(c *gin.Context) {
	id := c.Param("id")

	code := model.DeleteBonuse(id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	
	v1.SendResponse(c, errmsg.SUCCSE, id)
}