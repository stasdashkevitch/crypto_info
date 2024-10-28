package cryptotrackerusecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stasdashkevitch/crypto_info/internal/usecase/pubsub"
	cryptodataprovider "github.com/stasdashkevitch/crypto_info/internal/usecase/service/cryptoDataProvider"
)

type CryptoTrackerUsecase struct {
	cryptoDataProvider cryptodataprovider.CryptoDataProvider
	pubsub             pubsub.CryptoTrackerPubSub
}

func NewCryptoTrackerUsecase(cryptoDataProvider cryptodataprovider.CryptoDataProvider, pubsub pubsub.CryptoTrackerPubSub) *CryptoTrackerUsecase {
	return &CryptoTrackerUsecase{
		cryptoDataProvider: cryptoDataProvider,
		pubsub:             pubsub,
	}
}

func (u *CryptoTrackerUsecase) UpdateCryptoDataPrice(ctx context.Context, id string) error {
	data, err := u.cryptoDataProvider.GetCryptoDataPrice(ctx, id)
	if err != nil {
		return fmt.Errorf("Failed to retriev data: %v", err)
	}

	jsonData, _ := json.Marshal(data)

	return u.pubsub.Publish(ctx, "crypto_price_updates", jsonData)
}

func (u *CryptoTrackerUsecase) SubscribeCryptoDataPrice(ctx context.Context) (<-chan []byte, error) {
	return u.pubsub.Subscribe(ctx, "crypto_price_updates")
}
