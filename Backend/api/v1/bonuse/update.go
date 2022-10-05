package bonuse

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateBonuse(c *gin.Context) {
	var bonuse *model.Bonuse
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&bonuse); err != nil {
		log.Panicln(err)
	}
	code := model.CheckBonuse(bonuse.TraceId)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_CATENAME_USED, nil)
		return
	}
	code = model.UpdateBonuse(id, bonuse)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_CATENAME, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, bonuse.TraceId)
}
