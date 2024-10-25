package cryptodataprovider

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

const coinGeckoURL = "https://api.coingecko.com/api/v3/coins/markets"

type CryptoDataFromCoinGecko struct{}

func (p *CryptoDataFromCoinGecko) GetCryptoData(ctx context.Context, id string) (*entity.CryptoData, error) {
	url := fmt.Sprintf("%s?vs_currency=usd&ids=%s", coinGeckoURL, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error while making HTTP request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recieved non-OK http status: %s", err)
	}

	defer resp.Body.Close()

}
