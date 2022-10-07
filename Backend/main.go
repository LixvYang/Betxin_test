package main

import (
	"context"
	"log"
	"betxin/model"
	"betxin/router"
	"betxin/service"
	"betxin/utils"

	"github.com/fox-one/mixin-sdk-go"
)

func main() {

	model.InitDb()

	ctx := context.Background()
	store := &mixin.Keystore{
		ClientID:   utils.ClientId,
		SessionID:  utils.SessionId,
		PrivateKey: utils.PrivateKey,
		PinToken:   utils.PinToken,
	}

	// Create a Mixin Client from the keystore, which is the instance to invoke Mixin APIs
	client, err := mixin.NewFromKeystore(store)
	if err != nil {
		log.Panicln(err)
	}

	go service.Worker(ctx, client)
	// dailycurrency.DailyCurrency(ctx)

	router.InitRouter()
}
