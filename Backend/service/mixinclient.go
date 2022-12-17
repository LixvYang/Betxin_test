package service

import (
	"betxin/utils"
	"log"
	"sync"

	"github.com/fox-one/mixin-sdk-go"
)

var (
	InitMixin   sync.Once
	mixinClient *mixin.Client
	err         error
)

func NewmixinClient() {
	store := &mixin.Keystore{
		ClientID:   utils.ClientId,
		SessionID:  utils.SessionId,
		PrivateKey: utils.PrivateKey,
		PinToken:   utils.PinToken,
	}

	mixinClient, err = mixin.NewFromKeystore(store)
	if err != nil {
		log.Fatal(err)
	}
}

func MixinClient() *mixin.Client {
	return mixinClient
}