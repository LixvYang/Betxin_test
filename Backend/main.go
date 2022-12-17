package main

import "betxin/service"

func main() {
	// signalch := make(chan os.Signal, 1)
	// ctx := context.Background()
	service.InitMixin.Do(service.NewmixinClient)

	// utils.Setting.Do(utils.Init)
	// model.InitDb()
	// betxinredis.NewRedisClient(ctx)
	// go dailycurrency.DailyCurrency(ctx)
	// go service.Worker(ctx)
	// go router.InitRouter(signalch)

	// //attach signal
	// signal.Notify(signalch, os.Interrupt, os.Kill, syscall.SIGTERM)
	// signalType := <-signalch
	// signal.Stop(signalch)

	// //cleanup before exit
	// log.Println("On Signal <%s>", signalType)
	// log.Println("Exit command received. Exiting...")
}
