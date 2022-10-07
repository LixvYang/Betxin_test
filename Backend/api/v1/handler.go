package v1

import (
	"betxin/utils/errmsg"
	"betxin/utils/redis"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ctx context.Context

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code int, data interface{}) {
	message := errmsg.GetErrMsg(code)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Redis() *redis.RedisClient {
	ctx = context.Background()
	return redis.NewRedisClient(ctx)
}
