package userportfoliousecase

import (
	"testing"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
	repositorymock "github.com/stasdashkevitch/crypto_info/internal/usecase/mock/repositoryMock"
)

func TestUserPortfolioUsecase(t *testing.T) {
	goodUserPortfolio := &entity.UserPortfolio{
		UserID:             "1",
		CryptoID:           "first",
		IncreaseThreshold:  10,
		DecreaseThreshold:  10,
		NotifyIncrease:     true,
		NotifyDecrease:     true,
		NotificationMethod: "message",
	}
	mockRepo := &repositorymock.MockUserPortfolioRepository{
		UsersPortfolio: make(map[string]*entity.UserPortfolio),
	}

	mockRepo.Create(goodUserPortfolio)
	usecase := NewUserPortfolioUsecase(mockRepo)

	t.Run("Test for create and update functions", func(t *testing.T) {
		cases := []struct {
			Description   string
			UserPortfolio *entity.UserPortfolio
			Expected      string
		}{
			{
				"Succesfull example",
				&entity.UserPortfolio{
					UserID:             "2",
					CryptoID:           "succesfull",
					IncreaseThreshold:  10,
					DecreaseThreshold:  10,
					NotifyIncrease:     true,
					NotifyDecrease:     true,
					NotificationMethod: "message",
				},
				"",
			},
			{
				"With flag, but without amount of threshold",
				&entity.UserPortfolio{
					UserID:             "3",
					CryptoID:           "without amount",
					IncreaseThreshold:  0,
					DecreaseThreshold:  1,
					NotifyIncrease:     true,
					NotifyDecrease:     true,
					NotificationMethod: "message",
				},
				"If you specified a flag to monitor the price, then you must provide an amount that is greater than zero",
			},
			{
				"Without any flag",
				&entity.UserPortfolio{
					UserID:             "4",
					CryptoID:           "without any flag",
					IncreaseThreshold:  0,
					DecreaseThreshold:  1,
					NotifyIncrease:     false,
					NotifyDecrease:     false,
					NotificationMethod: "message",
				},
				"You must select one of the checkboxes and indicate an amount greater than zero",
			},
			{
				"Without notification method",
				&entity.UserPortfolio{
					UserID:             "4",
					CryptoID:           "without any flag",
					IncreaseThreshold:  1,
					DecreaseThreshold:  1,
					NotifyIncrease:     true,
					NotifyDecrease:     true,
					NotificationMethod: "",
				},
				"You must provide on of the notification methods",
			},
		}

		for _, test := range cases {
			t.Run(test.Description, func(t *testing.T) {
				gotCreate := usecase.CreateUserPortfolio(test.UserPortfolio)
				gotUpdate := usecase.UpdateUserPortfolio(test.UserPortfolio)

				if gotCreate != nil && (gotCreate.Error() != test.Expected) {
					t.Errorf("Create: expected %v, but got: %v", test.Expected, gotCreate)
					return
				}

				if gotUpdate != nil && (gotUpdate.Error() != test.Expected) {
					t.Errorf("Create: expected %v, but got: %v", test.Expected, gotUpdate)
					return
				}
			})
		}
	})
}
