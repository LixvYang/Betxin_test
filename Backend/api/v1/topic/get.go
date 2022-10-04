package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetArtInfo 查询单个话题信息
func GetTopicInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetTopicById(id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_GET_TOPIC, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, data)
}
