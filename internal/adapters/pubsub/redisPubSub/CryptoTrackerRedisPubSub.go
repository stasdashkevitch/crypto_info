package redispubsub

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CryptoTrackerRedisPubSub struct {
	client *redis.Client
}

func NewCryptoTrackerRedisPubSub(client *redis.Client) *CryptoTrackerRedisPubSub {
	return &CryptoTrackerRedisPubSub{
		client: client,
	}
}

func (ps *CryptoTrackerRedisPubSub) Publish(ctx context.Context, channel string, message []byte) error {
	return ps.client.Publish(ctx, channel, message).Err()
}

func (ps *CryptoTrackerRedisPubSub) Subscribe(ctx context.Context, channel string) (<-chan []byte, error) {
	pubsub := ps.client.Subscribe(ctx, channel)

	ch := make(chan []byte)

	go func() {
		defer pubsub.Close()

		for {
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				close(ch)
				return
			}

			ch <- []byte(msg.Payload)
		}
	}()

	return ch, nil
}
