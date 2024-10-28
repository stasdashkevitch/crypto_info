package repository

import "github.com/stasdashkevitch/crypto_info/internal/entity"

type UserPortfolioRepository interface {
	Create(userPortfolio *entity.UserPortfolio) error
	GetAllByUserID(id string) ([]*entity.UserPortfolio, error)
	GetByCryptoID(userID, cryptoID string) (*entity.UserPortfolio, error)
	Update(userPortfolio *entity.UserPortfolio) error
	Delete(userID, cryptoID string) error
}
