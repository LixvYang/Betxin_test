package main

import (
	"betxin/model"
	"betxin/router"
	"betxin/service"
	"betxin/service/dailycurrency"
	"betxin/utils"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	betxinredis "betxin/utils/redis"
)

func main() {
	signalch := make(chan os.Signal, 1)
	utils.Setting.Do(utils.Init)
	service.InitMixin.Do(service.InitMixinClient)

	ctx := context.Background()
	model.InitDb()
	betxinredis.NewRedisClient(ctx)
	go dailycurrency.DailyCurrency(ctx)
	go service.Worker(ctx)
	go router.InitRouter(signalch)

	//attach signal
	signal.Notify(signalch, os.Interrupt, syscall.SIGTERM)
	signalType := <-signalch
	signal.Stop(signalch)
	//cleanup before exit
	log.Printf("On Signal <%s>", signalType)
	log.Println("Exit command received. Exiting...")
}
