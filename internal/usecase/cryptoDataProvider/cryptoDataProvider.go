package cryptodataprovider

import (
	"context"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

type CryptoDataProvider interface {
	GetCryptoData(ctx context.Context, id string) (*entity.CryptoData, error)
}
