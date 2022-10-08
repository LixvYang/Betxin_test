package administrator

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateAdministrator(c *gin.Context) {
	var admin model.Administrator
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Panicln(err)
	}

	code := model.UpdateAdministrator(id, &admin)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_USER, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, admin.Id)
}
