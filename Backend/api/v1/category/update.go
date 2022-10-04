package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(c *gin.Context) {
	var cate *model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&cate)
	fmt.Println(cate)
	code := model.CheckCategory(cate.CategoryName)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_CATENAME_USED, nil)
		return
	}
	code = model.UpdateCate(id, cate.CategoryName)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_CATENAME, nil)
	}

	v1.SendResponse(c, errmsg.SUCCSE, cate.CategoryName)
}
