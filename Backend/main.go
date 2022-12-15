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
	ctx := context.Background()
	model.InitDb()
	service.NewMixinClient()
	betxinredis.NewRedisClient(ctx)
	go dailycurrency.DailyCurrency(ctx)
	go service.Worker(ctx)
	router.InitRouter()
}
