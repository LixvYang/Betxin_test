package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/fox-one/mixin-sdk-go"
)

type WorkerQueue struct {
	SubClients []*mixin.Client
}

func NewWorkderQueue(ctx context.Context, client *mixin.Client) WorkerQueue {
	var WorkerQueue WorkerQueue
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
			sub, subStore, err := client.CreateUser(ctx, privateKey, fmt.Sprintf("betxin Bot", strconv.Itoa(i)))

			if err != nil {
				log.Printf("CreateUser: %v", err)
				return
			}

			newPin := mixin.RandomPin()
			subClient, _ := mixin.NewFromKeystore(subStore)
			if err := subClient.ModifyPin(ctx, "", newPin); err != nil {
				log.Printf("ModifyPin: %v", err)
				return
			}

			fmt.Println(sub.FullName, sub.PinToken, subClient.ClientID)

			if err := subClient.VerifyPin(ctx, newPin); err != nil {
				log.Printf("sub user VerifyPin: %v", err)
				return
			}
			WorkerQueue.SubClients = append(WorkerQueue.SubClients, subClient)
		}(i)
	}

	wg.Wait()
	return WorkerQueue
}

func (w *WorkerQueue) Pop() *mixin.Client {
	subClient := w.SubClients[0]
	w.SubClients = w.SubClients[1:]
	w.SubClients = append(w.SubClients, subClient)
	return subClient
}
