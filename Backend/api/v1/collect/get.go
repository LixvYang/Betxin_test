package collect

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"
	betxinredis "betxin/utils/redis"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type GetRequest struct {
	UserId string `json:"user_id"`
}

func GetCollect(c *gin.Context) {
	var total int
	var data []model.Collect
	var err error
	var collect string
	var code int
	userId := c.Param("userId")

	total, _ = betxinredis.Get(v1.COLLECT_GET_USER_TOTAL + userId).Int()
	collect, err = betxinredis.Get(v1.COLLECT_GET_USER_LIST + userId).Result()
	convert.Unmarshal(collect, &data)
	if err == redis.Nil {
		data, total, code = model.GetCollectByUserId(userId)
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR, nil)
			return
		}

		collect = convert.Marshal(&data)
		betxinredis.Set(v1.COLLECT_GET_USER_TOTAL, total, v1.REDISEXPIRE)
		betxinredis.Set(v1.COLLECT_GET_USER_LIST, collect, v1.REDISEXPIRE)
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	} else if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	} else {
		fmt.Println("从redis拿数据")
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	}
}
