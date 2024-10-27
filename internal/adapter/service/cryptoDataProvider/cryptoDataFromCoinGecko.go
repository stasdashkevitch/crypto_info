package cryptodataprovider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

const coinGeckoURL = "https://api.coingecko.com/api/v3/"

type CryptoDataFromCoinGeckoProvider struct{}

func NewCryptoDataFromCoinGeckoProvide() *CryptoDataFromCoinGeckoProvider {
	return &CryptoDataFromCoinGeckoProvider{}
}

func (p *CryptoDataFromCoinGeckoProvider) GetCryptoDataPrice(ctx context.Context, id string) (*entity.CryptoDataPrice, error) {
	url := fmt.Sprintf("%ssimple/price?ids=%s&vs_currencies=usd&&include_market_cap=true&include_24hr_vol=true&include_24hr_change=true&include_last_updated_at=true", coinGeckoURL, id)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.New("Failed to create request")
	}
	req.Header.Add("accept", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while making HTTP request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recieved non-OK http status: %d", resp.StatusCode)
	}

	var result map[string]struct {
		USD          float64 `json:"usd"`
		USDMarketCap float64 `json:"usd_market_cap"`
		USD24Vol     float64 `json:"usd_24h_vol"`
		USD24hChange float64 `json:"usd_24h_change"`
		LastUpdateAt int64   `json:"last_updated_at"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error while decoding response: %v", err)
	}

	data, exists := result[id]
	if !exists {
		return nil, fmt.Errorf("no data found for ID: %s", id)
	}

	defer resp.Body.Close()

	return &entity.CryptoDataPrice{
		ID:        id,
		Name:      id,
		PriceUSD:  data.USD,
		MarketCap: data.USDMarketCap,
		Volume24h: data.USD24hChange,
		Timestamp: time.Unix(data.LastUpdateAt, 0),
	}, nil

}
