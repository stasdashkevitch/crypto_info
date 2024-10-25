package repository

import "github.com/stasdashkevitch/crypto_info/internal/entity"

type CryptoDatarepository interface {
	SetCryptoData(*entity.CryptoData) error
	GetCrpytoData(string) (*entity.CryptoData, error)
	UpdateCryptoData(*entity.CryptoData) error
}
