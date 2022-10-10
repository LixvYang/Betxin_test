package category

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"
	betxinredis "betxin/utils/redis"

	"github.com/gin-gonic/gin"
)

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var cate *model.Category
	if err := c.ShouldBindJSON(&cate); err != nil {
		v1.SendResponse(c, errmsg.ERROR_BIND, nil)
	}

	code := model.CheckCategory(cate.CategoryName)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_CATENAME_USED, nil)
		return
	}

	code = model.UpdateCate(convert.StrToNum(id), cate.CategoryName)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_UPDATE_CATENAME, nil)
		return
	}

	// Delete redis store
	betxinredis.DelKeys(v1.CATEGORY_GET+id, v1.CATEGORY_LIST, v1.CATEGORY_TOTAL)

	v1.SendResponse(c, errmsg.SUCCSE, cate.CategoryName)
}
