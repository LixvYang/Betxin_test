package service

import (
	"betxin/utils"
	"errors"

	"github.com/fox-one/mixin-sdk-go"
)

var (
	mixinClient *mixin.Client
	err         error
)

func NewMixinClient() {
	store := &mixin.Keystore{
		ClientID:   utils.ClientId,
		SessionID:  utils.SessionId,
		PrivateKey: utils.PrivateKey,
		PinToken:   utils.PinToken,
	}

	mixinClient, err = mixin.NewFromKeystore(store)
	if err != nil {
		errors.New("mixin bot client error...")
	}
}

func MixinClient() *mixin.Client {
	return mixinClient
}
