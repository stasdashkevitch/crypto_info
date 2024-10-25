package cryptotrackerusecase

import (
	"context"
	"errors"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/repository"
)

type CryptoTrackerUsecase struct {
	repository repository.CryptoDatarepository
}

func NewCryptoTrackerUsecase(repository repository.CryptoDatarepository) *CryptoTrackerUsecase {
	return &CryptoTrackerUsecase{
		repository: repository,
	}
}

func (u *CryptoTrackerUsecase) GetCryptoData(ctx context.Context, id string) (*entity.CryptoData, error) {
	data, err := u.repository.GetCrpytoData(id)

	if err == nil && data != nil {
		return data, nil
	}

	// TODO
	// externalData, err
}
