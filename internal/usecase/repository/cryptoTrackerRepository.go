package repository

import "github.com/stasdashkevitch/crypto_info/internal/entity"

type CryptoDatarepository interface {
	SetCryptoData(*entity.CryptoData) error
	GetCrpytoData(id string) (*entity.CryptoData, error)
	DeleteCryptoData(id string) error
	UpdateCryptoData(id string) error
}
