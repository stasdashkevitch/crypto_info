package repositories

import "github.com/stasdashkevitch/crypto_info/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByEmail(email string) (*entities.User, error)
	GetByUsername(username string) (*entities.User, error)
}
