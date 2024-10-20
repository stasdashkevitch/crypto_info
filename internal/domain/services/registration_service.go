package services

import (
	"errors"
	"time"

	"github.com/stasdashkevitch/crypto_info/internal/domain/entities"
	"github.com/stasdashkevitch/crypto_info/internal/domain/repositories"
	"github.com/stasdashkevitch/crypto_info/internal/dtos"
	"github.com/stasdashkevitch/crypto_info/pkg"
	"github.com/stasdashkevitch/crypto_info/pkg/password"
)

type RegistrationService struct {
	userRepository repositories.UserRepository
}

func NewRegistrationUsecase(userRepository repositories.UserRepository) *RegistrationService {
	return &RegistrationService{
		userRepository: userRepository,
	}
}

func (u *RegistrationService) Register(registerUserDTO dtos.RegisterUserDTO) error {
	if registerUserDTO.Username == "" || registerUserDTO.Email == "" || registerUserDTO.Password == "" {
		return errors.New("username, email, password are required")
	}

	existingUser, _ := u.userRepository.GetByEmail(registerUserDTO.Email)
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := password.HashPassword(registerUserDTO.Password)
	if err != nil {
		return err
	}

	user := &entities.User{
		ID:           pkg.GenerateID(),
		Username:     registerUserDTO.Username,
		PasswordHash: hashedPassword,
		Email:        registerUserDTO.Email,
		CreatedAt:    time.Now(),
	}

	return u.userRepository.Create(user)
}
