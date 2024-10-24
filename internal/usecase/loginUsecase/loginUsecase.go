package loginusecase

import (
	"errors"

	"github.com/stasdashkevitch/crypto_info/internal/dtos"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/auth"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/repository"
	"github.com/stasdashkevitch/crypto_info/pkg/password"
)

type LoginUsecase struct {
	auth       auth.Auth
	repository repository.UserRepository
}

func NewLoginUsecase(auth auth.Auth, repository repository.UserRepository) *LoginUsecase {
	return &LoginUsecase{
		auth:       auth,
		repository: repository,
	}
}

func (s *LoginUsecase) Login(dto dtos.LoginUserDTO) (string, error) {
	user, err := s.repository.GetByEmail(dto.Email)

	if err != nil || user == nil {
		return "", errors.New("User with this email does not exists")
	}

	if err := password.ValidatePassword(user.PasswordHash, dto.Password); err != nil {
		return "", errors.New("Invalid password")
	}

	token, err := s.auth.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
