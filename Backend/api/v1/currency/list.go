package currency

import (
	v1 "betxin/api/v1"
	"betxin/model"
	"betxin/utils/convert"
	"betxin/utils/errmsg"
	betxinredis "betxin/utils/redis"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type ListResponse struct {
	TotalCount int              `json:"totalCount"`
	List       []model.Currency `json:"list"`
}

func ListCurrencies(c *gin.Context) {
	var data []model.Currency
	var total int
	var code int
	var err error
	var currencies string
	total, _ = betxinredis.Get(v1.CURRENCY_TOTAL).Int()
	currencies, err = betxinredis.Get(v1.CURRENCY_LIST).Result()
	convert.Unmarshal(currencies, &data)
	if err == redis.Nil {
		data, total, code = model.ListCurrencies()
		if code != errmsg.SUCCSE {
			v1.SendResponse(c, errmsg.ERROR, nil)
			return
		}

		currencies = convert.Marshal(&data)
		betxinredis.Set(v1.CURRENCY_TOTAL, total, time.Minute/2)
		betxinredis.Set(v1.CURRENCY_LIST, currencies, time.Minute/2)

		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	} else if err != nil {
		v1.SendResponse(c, errmsg.ERROR, nil)
		return
	} else {
		v1.SendResponse(c, errmsg.SUCCSE, ListResponse{
			TotalCount: total,
			List:       data,
		})
	}
}
