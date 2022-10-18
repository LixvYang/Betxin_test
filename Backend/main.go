package main

import (
	"betxin/model"
	"betxin/router"
	"betxin/service"
	"betxin/service/dailycurrency"
	"betxin/utils"
	betxinredis "betxin/utils/redis"
	"context"
)

func main() {
	utils.InitIni()	
	model.InitDb()
	var ctx = context.Background()
	betxinredis.NewRedisClient(ctx)
	service.NewMixinClient()
	dailycurrency.DailyCurrency(ctx)
	go service.Worker(ctx, service.MixinClient())
	router.InitRouter()
}
