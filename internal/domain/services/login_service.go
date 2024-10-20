package services

import (
	"errors"

	"github.com/stasdashkevitch/crypto_info/internal/domain/repositories"
	"github.com/stasdashkevitch/crypto_info/internal/domain/services/auth"
	"github.com/stasdashkevitch/crypto_info/internal/dtos"
	"github.com/stasdashkevitch/crypto_info/pkg/password"
)

type LoginServis struct {
	authService    auth.AuthService
	userRepository repositories.UserRepository
}

func NewLoginService(authService *auth.AuthService) *LoginServis {
	return &LoginServis{
		authService: *authService,
	}
}

func (s *LoginServis) Login(loginUserDTO *dtos.LoginUserDTO) (string, error) {
	user, err := s.userRepository.GetByUsername(loginUserDTO.Username)

	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	if err := password.ValidatePassword(user.PasswordHash, loginUserDTO.Password); err != nil {
		return "", errors.New("invalid passsword")
	}

	token, err := s.authService.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
