package topic

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/errmsg"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

// GetArtInfo 查询单个话题信息
func GetTopicInfoById(c *gin.Context) {
	id, _ := uuid.FromString(c.Param("id"))
	data, code := model.GetTopicById(id)
	if code != errmsg.SUCCSE {
		v1.SendResponse(c, errmsg.ERROR_GET_TOPIC, nil)
		return
	}
	v1.SendResponse(c, errmsg.SUCCSE, data)
}
