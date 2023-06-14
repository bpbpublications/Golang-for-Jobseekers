package main

import (
	"context"
	"fmt"

	nats "github.com/nats-io/nats.go"
)

type Nats struct {
	Conn         *nats.Conn
	Topic        string
	Subscription *nats.Subscription
}

func NewNats(natsEndpoint string, topic string) (Nats, error) {
	conn, err := nats.Connect(natsEndpoint)
	if err != nil {
		return Nats{}, fmt.Errorf("Error with connecting to Nats. Err: %v", err)
	}
	s, err := conn.SubscribeSync(topic)
	if err != nil {
		return Nats{}, fmt.Errorf("Error with creating the subscriber. Err: %v", err)
	}
	return Nats{
		Conn:         conn,
		Topic:        topic,
		Subscription: s,
	}, nil
}

func (n Nats) Add(ctx context.Context, message []byte) error {
	// Function not defined here as this is not the main focus for this section of the chapter
	return nil
}

func (n Nats) Pop(ctx context.Context) ([]byte, error) {
	// Function not defined here as this is not the main focus for this section of the chapter
	return []byte{}, nil
}

func main() {}
