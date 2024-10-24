package loginusecase

import (
	"testing"

	"github.com/stasdashkevitch/crypto_info/internal/dtos"
	"github.com/stasdashkevitch/crypto_info/internal/entity"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/mock"
	repositorymock "github.com/stasdashkevitch/crypto_info/internal/usecase/mock/repositoryMock"
	"github.com/stasdashkevitch/crypto_info/pkg/password"
)

func TestLoginUsecase(t *testing.T) {
	password, _ := password.HashPassword("123")

	user := &entity.User{
		ID:           "1",
		Email:        "example@gmail.com",
		PasswordHash: password,
	}
	mockRepo := &repositorymock.MockUserRepository{
		Users: make(map[string]*entity.User),
	}

	mockRepo.Create(user)

	mockAuth := &mock.MockAuth{}
	loginUsecase := NewLoginUsecase(mockAuth, mockRepo)

	cases := []struct {
		Description string
		UserData    dtos.LoginUserDTO
		Error       string
		Expected    string
	}{
		{
			"Succesfull login",
			dtos.LoginUserDTO{
				Email:    "example@gmail.com",
				Password: "123",
			},
			"",
			"token",
		},
		{
			"Invalid password",
			dtos.LoginUserDTO{
				Email:    "example@gmail.com",
				Password: "12",
			},
			"Invalid password",
			"",
		},
		{
			"Non-existent user",
			dtos.LoginUserDTO{
				Email:    "",
				Password: "123",
			},
			"User with this email does not exists",
			"",
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got, err := loginUsecase.Login(test.UserData)

			if got == "" {
				if err.Error() != test.Error {
					t.Errorf("Expected %v, got %v", test.Error, err)
				}
			}

			if got != test.Expected {
				t.Errorf("Expected token: %v, got %v", test.Expected, got)
			}
		})
	}
}
