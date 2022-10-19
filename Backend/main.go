package main

import (
	"betxin/model"
	"betxin/router"
	"betxin/service"
	betxinredis "betxin/utils/redis"
	"context"
)

func main() {
	model.InitDb()
	var ctx = context.Background()
	betxinredis.NewRedisClient(ctx)
	service.NewMixinClient()
	// go dailycurrency.DailyCurrency(ctx)
	go service.Worker(ctx, service.MixinClient())
	router.InitRouter()
}
