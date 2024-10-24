package entity

import "time"

type CryptoData struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	PriceUSD         float64   `json:"price_usd"`
	MarketCap        float64   `json:"market_cap"`
	Volume24h        float64   `json:"volume24h"`
	PercentChange24h float64   `json:"percent_change24h"`
	AvailableSupply  float64   `json:"available_supply"`
	TotalSupply      float64   `json:"total_supply"`
	MaxSupply        float64   `json:"max_supply"`
	Timestamp        time.Time `json:"timestamp"`
}
