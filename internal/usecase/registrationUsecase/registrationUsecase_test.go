package registrationusecase

import (
	"testing"

	"github.com/stasdashkevitch/crypto_info/internal/dto"
	"github.com/stasdashkevitch/crypto_info/internal/entity"
	repositorymock "github.com/stasdashkevitch/crypto_info/internal/usecase/mock/repositoryMock"
)

func TestRegistrationUsecase(t *testing.T) {
	user := &entity.User{
		Email: "exists@gmail.com",
	}

	mockRepo := &repositorymock.MockUserRepository{
		Users: make(map[string]*entity.User),
	}

	mockRepo.Create(user)

	registrationUsecase := NewRegistrationUsecase(mockRepo)

	cases := []struct {
		Description string
		UserData    dto.RegisterUserDTO
		Error       string
	}{
		{
			"Succesfull user creation",
			dto.RegisterUserDTO{
				Email:    "example@gmail.com",
				Username: "Ivan",
				Password: "aaa",
			},
			"",
		},
		{
			"Without one of the fields",
			dto.RegisterUserDTO{
				Email: "example@gmail.com",
			},
			"Username, email, password are required",
		},
		{
			"User already exists",
			dto.RegisterUserDTO{
				Email:    "exists@gmail.com",
				Username: "Ivan",
				Password: "aaa",
			},
			"User with this email already exists",
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			err := registrationUsecase.Register(test.UserData)

			if err != nil && (err.Error() != test.Error) {
				t.Errorf("Expected %v, got %v", test.Error, err)
			}

		})
	}
}
