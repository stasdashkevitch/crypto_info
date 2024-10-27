package cryptodataprovider

import (
	"context"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

type CryptoDataProvider interface {
	GetCryptoDataPrice(ctx context.Context, id string) (*entity.CryptoDataPrice, error)
}
