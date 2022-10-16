package main

import (
	"betxin/model"
	"betxin/router"
	"betxin/service"
	"betxin/service/dailycurrency"
	betxinredis "betxin/utils/redis"
	"context"
)

func main() {

	model.InitDb()
	var ctx = context.Background()
	dailycurrency.DailyCurrency(ctx)
	service.NewMixinClient()
	go betxinredis.NewRedisClient(ctx)
	go service.Worker(ctx, service.MixinClient())
	router.InitRouter()
}
