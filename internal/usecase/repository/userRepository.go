package repository

import "github.com/stasdashkevitch/crypto_info/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByEmail(email string) (*entity.User, error)
	GetByUsername(username string) (*entity.User, error)
}
