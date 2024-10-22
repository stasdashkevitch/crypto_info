package usecase

import (
	"errors"
	"time"

	"github.com/stasdashkevitch/crypto_info/internal/dtos"
	"github.com/stasdashkevitch/crypto_info/internal/entity"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/repository"
	"github.com/stasdashkevitch/crypto_info/pkg/id"
	"github.com/stasdashkevitch/crypto_info/pkg/password"
)

type RegistrationService struct {
	repository repository.UserRepository
}

func NewRegistrationUsecase(repository repository.UserRepository) *RegistrationService {
	return &RegistrationService{
		repository: repository,
	}
}

func (u *RegistrationService) Register(dto dtos.RegisterUserDTO) error {
	if dto.Username == "" || dto.Email == "" || dto.Password == "" {
		return errors.New("username, email, password are required")
	}

	existingUser, _ := u.repository.GetByEmail(dto.Email)

	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := password.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	user := &entity.User{
		ID:           id.GenerateID(),
		Username:     dto.Username,
		PasswordHash: hashedPassword,
		Email:        dto.Email,
		CreatedAt:    time.Now(),
	}

	return u.repository.Create(user)
}
