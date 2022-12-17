package mq

import (
	"fmt"
	"sync"
	"testing"
)

type Message struct {
	ID   int
	Name string
}

func TestClient(t *testing.T) {
	var m Message
	m.ID = 12
	m.Name = "llllllsssssss"

	b := NewMQClient()
	b.SetConditions(100)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		topic := fmt.Sprintf("Golang梦工厂%d", i)
		payload := m

		ch, err := b.Subscribe(topic)
		if err != nil {
			t.Fatal(err)
		}

		wg.Add(1)
		go func() {
			e := b.GetPayLoad(ch).(Message)
			fmt.Println(e)
			if e != payload {
				t.Fatal(topic, " expected ", payload, " but get", e)
			}
			if err := b.Unsubscribe(topic, ch); err != nil {
				t.Fatal(err)
			}
			wg.Done()
		}()

		if err := b.Publish(topic, payload); err != nil {
			t.Fatal(err)
		}
	}

	wg.Wait()
}
