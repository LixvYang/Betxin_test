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
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	model.InitDb()
	var ctx = context.Background()
	betxinredis.NewRedisClient(ctx)
	service.NewMixinClient()
	go dailycurrency.DailyCurrency(ctx)
	go service.Worker(ctx, service.MixinClient())
	router.InitRouter()
}
