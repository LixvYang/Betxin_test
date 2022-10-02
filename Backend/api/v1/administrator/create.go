package administrator

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type CreateResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateAdministrator(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}

	data := &model.Administrator{
		Username: r.Password,
		Password: r.Password,
	}
	if code := model.CreateAdministrator(data); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	rsp := CreateResponse(r)
	v1.SendResponse(c, errmsg.SUCCSE, rsp)
}
