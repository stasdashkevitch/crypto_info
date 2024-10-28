package entity

type UserPortfolio struct {
	UserID             string  `json:"user_id"`
	CryptoID           string  `json:"crypto_id"`
	IncreaseThreshold  float64 `json:"increase_threshold"`
	DecreaseThreshold  float64 `json:"decrease_threshold"`
	NotifyIncrease     bool    `json:"notify_increase"`
	NotifyDecrease     bool    `json:"notify_decrease"`
	NotificationMethod string  `json:"notification_method"`
}
