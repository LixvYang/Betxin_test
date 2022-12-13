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
	go service.Worker(ctx)
	// go timewheel.At(time.Now().Add(time.Second), "", func() {
	// 	fmt.Println("nihao")
	// })
	// go timewheel.Every(time.Second, func() {
	// 	fmt.Println("nihao")
	// })
	router.InitRouter()
}
