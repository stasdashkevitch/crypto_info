package entity

import (
	"time"
)

type (
	User struct {
		ID           string    `json:"id"`
		Username     string    `json:"username"`
		Email        string    `json:"email"`
		PasswordHash string    `json:"-"`
		CreatedAt    time.Time `json:"created_at"`
	}
)
