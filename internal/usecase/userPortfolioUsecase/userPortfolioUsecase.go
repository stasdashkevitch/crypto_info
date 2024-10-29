package userportfoliousecase

import (
	"errors"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/repository"
)

type UserPortfolioUsecase struct {
	repository repository.UserPortfolioRepository
}

func NewUserPortfolioUsecase(repository repository.UserPortfolioRepository) *UserPortfolioUsecase {
	return &UserPortfolioUsecase{
		repository: repository,
	}
}

func (u *UserPortfolioUsecase) CreateUserPortfolio(userPortfolio *entity.UserPortfolio) error {
	if !userPortfolio.NotifyIncrease && !userPortfolio.NotifyDecrease {
		return errors.New("You must select one of the checkboxes and indicate an amount greater than zero")
	}

	if (userPortfolio.NotifyIncrease && userPortfolio.IncreaseThreshold <= 0) || (userPortfolio.NotifyDecrease && userPortfolio.DecreaseThreshold <= 0) {
		return errors.New("If you specified a flag to monitor the price, then you must provide an amount that is greater than zero")
	}

	if userPortfolio.NotificationMethod == "" {
		return errors.New("You must provide on of the notification methods")
	}

	return u.repository.Create(userPortfolio)
}

func (u *UserPortfolioUsecase) GetAllUserPortfolio(userID string) ([]*entity.UserPortfolio, error) {
	allUserPortfolio, err := u.GetAllUserPortfolio(userID)
	if err != nil {
		return nil, err
	}

	return allUserPortfolio, nil
}

func (u *UserPortfolioUsecase) GetUserPortfolio(userID, cryptoID string) (*entity.UserPortfolio, error) {
	userPortfolio, err := u.repository.GetByCryptoID(userID, cryptoID)
	if err != nil {
		return nil, err
	}

	return userPortfolio, nil
}

func (u *UserPortfolioUsecase) UpdateUserPortfolio(userPortfolio *entity.UserPortfolio) error {
	if (userPortfolio.NotifyIncrease && userPortfolio.IncreaseThreshold <= 0) || (userPortfolio.NotifyDecrease && userPortfolio.DecreaseThreshold <= 0) {
		return errors.New("If you specified a flag to monitor the price, then you must provide an amount that is greater than zero")
	}

	if !userPortfolio.NotifyIncrease && !userPortfolio.NotifyDecrease {
		return errors.New("You must select one of the checkboxes and indicate an amount greater than zero")
	}

	if userPortfolio.NotificationMethod == "" {
		return errors.New("You must provide on of the notification methods")
	}

	return u.repository.Update(userPortfolio)
}

func (u *UserPortfolioUsecase) DeleteUserPortfolio(userID, cryptoID string) error {
	return u.repository.Delete(userID, cryptoID)
}
