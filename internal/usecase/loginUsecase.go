package usecase

import (
	"errors"

	"github.com/stasdashkevitch/crypto_info/internal/dtos"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/auth"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/repository"
	"github.com/stasdashkevitch/crypto_info/pkg/password"
)

type LoginServis struct {
	auth       auth.Auth
	repository repository.UserRepository
}

func NewLoginService(auth auth.Auth, repository repository.UserRepository) *LoginServis {
	return &LoginServis{
		auth:       auth,
		repository: repository,
	}
}

func (s *LoginServis) Login(dto dtos.LoginUserDTO) (string, error) {
	user, err := s.repository.GetByEmail(dto.Email)

	if err != nil || user == nil {
		return "", err
	}

	if err := password.ValidatePassword(user.PasswordHash, dto.Password); err != nil {
		return "", errors.New("invalid passsword")
	}

	token, err := s.auth.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
