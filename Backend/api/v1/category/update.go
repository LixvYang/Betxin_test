package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(c *gin.Context) {
	var data *model.Category
	_ = c.ShouldBindJSON(&data)

	code := model.CheckCategory(data.CategoryName)
	if code == errmsg.SUCCSE {
		model.UpdateCate(data)
	}

	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	v1.SendResponse(c, errmsg.SUCCSE, nil)
}
