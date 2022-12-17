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
	ctx := context.Background()
	utils.Setting.Do(utils.Init)
	model.InitDb()
	service.NewMixinClient()
	betxinredis.NewRedisClient(ctx)
	go dailycurrency.DailyCurrency(ctx)
	go service.Worker(ctx)
	router.InitRouter()
}
