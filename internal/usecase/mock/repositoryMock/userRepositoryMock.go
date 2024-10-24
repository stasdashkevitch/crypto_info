package repositorymock

import (
	"errors"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

type MockUserRepository struct {
	Users map[string]*entity.User
}

func (m *MockUserRepository) Create(user *entity.User) error {
	_, exists := m.Users[user.Email]
	if exists {
		return errors.New("User already exists")
	}
	m.Users[user.Email] = user
	return nil
}

func (m *MockUserRepository) GetByEmail(email string) (*entity.User, error) {
	user, exists := m.Users[email]
	if !exists {
		return nil, errors.New("User with this email does not exists")
	}

	return user, nil
}
