package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	CategoryName string `json:"category_name"`
}
type CreateResponse struct {
	CategoryName string `json:"category_name"`
}

func CreateCatrgory(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
		return
	}
	data := &model.Category{
		CategoryName: r.CategoryName,
	}

	code := model.CheckCategory(data.CategoryName)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}
	if code = model.CreateCatrgory(data); code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	rsp := CreateResponse(r)
	v1.SendResponse(c, errmsg.SUCCSE, rsp)
}
