package administrator

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func CreateAdministratorLixin() {
	d := &model.Administrator{
		Username: "Lixin",
		Password: "123123123",
	}
	if code := model.CreateAdministrator(d); code != errmsg.SUCCSE {
		return
	}
}

// @Summary 创建管理员
// @Description 创建管理员
// @Tags administrator
// @Accept  json
// @Produce  json
// @Param category body administrator.model true "创建管理员"
// @Success 200 {object} v1.Response "{"code":200,"message":"OK","data":null}"
// @Router /v1/administrator/add [post]
func CreateAdministrator(c *gin.Context) {
	var r model.Administrator
	if err := c.ShouldBindJSON(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	code := model.CheckAdministrator(r.Username)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_USERNAME_USED, nil)
		return
	}

	if code := model.CreateAdministrator(&r); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	
	v1.SendResponse(c, errmsg.SUCCSE, r.Id)
}
