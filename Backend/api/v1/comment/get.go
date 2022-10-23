package comment

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

func GetCommentById(c *gin.Context) {
	id := c.Param("id")

	comment, code := model.GetCommentById(convert.StrToNum(id))
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	}

	v1.SendResponse(c, errmsg.SUCCSE, comment)
}
