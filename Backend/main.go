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
	betxinredis.NewRedisClient(ctx)
	service.NewMixinClient()
	dailycurrency.DailyCurrency(ctx)
	go service.Worker(ctx, service.MixinClient())
	router.InitRouter()
}
