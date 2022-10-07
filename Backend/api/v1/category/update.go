package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(c *gin.Context) {
	var cate *model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&cate); err != nil {
		log.Panicln(err)
	}

	code := model.CheckCategory(cate.CategoryName)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_CATENAME_USED, nil)
		return
	}
	
	code = model.UpdateCate(id, cate.CategoryName)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_CATENAME, nil)
		return
	}

	// Delete redis store
	if v1.Redis().Exists("categoryies") || v1.Redis().Exists("categoryiesTotal") {
		v1.Redis().Del("categoryiesTotal")
		v1.Redis().Del("categoryies")
	}

	v1.SendResponse(c, errmsg.SUCCSE, cate.CategoryName)
}
