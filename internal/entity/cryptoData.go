package entity

import "time"

type CryptoDataPrice struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	PriceUSD         float64   `json:"price_usd"`
	MarketCap        float64   `json:"market_cap"`
	Volume24h        float64   `json:"volume24h"`
	PercentChange24h float64   `json:"percent_change24h"`
	LastUpdatedAt    float64   `json:"last_updated_at"`
	Timestamp        time.Time `json:"timestamp"`
}
