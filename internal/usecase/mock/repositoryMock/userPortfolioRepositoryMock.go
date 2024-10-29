package repositorymock

import (
	"errors"
	"strings"

	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

type MockUserPortfolioRepository struct {
	UsersPortfolio map[string]*entity.UserPortfolio
}

func (m *MockUserPortfolioRepository) Create(userPortfolio *entity.UserPortfolio) error {
	key := userPortfolio.UserID + userPortfolio.CryptoID
	if _, exists := m.UsersPortfolio[key]; exists {
		return errors.New("User portfolio already exists")
	}

	m.UsersPortfolio[key] = userPortfolio

	return nil
}

func (m *MockUserPortfolioRepository) GetAllByUserID(id string) ([]*entity.UserPortfolio, error) {
	var allUserPortfolio []*entity.UserPortfolio
	for key, value := range m.UsersPortfolio {
		if strings.HasPrefix(key, id) {
			allUserPortfolio = append(allUserPortfolio, value)
		}
	}

	if len(allUserPortfolio) == 0 {
		return nil, errors.New("Cannot find all user portfolio with this id")
	}

	return allUserPortfolio, nil
}

func (m *MockUserPortfolioRepository) GetByCryptoID(userID, cryptoID string) (*entity.UserPortfolio, error) {
	key := userID + cryptoID
	if value, exists := m.UsersPortfolio[key]; exists {
		return value, nil
	}

	return nil, errors.New("Cannot find user portfolio with this id")
}

func (m *MockUserPortfolioRepository) Update(userPortfolio *entity.UserPortfolio) error {
	key := userPortfolio.UserID + userPortfolio.CryptoID
	if _, exists := m.UsersPortfolio[key]; !exists {
		return errors.New("User portfolio doesn't exists")
	}

	m.UsersPortfolio[key] = userPortfolio

	return nil
}

func (m *MockUserPortfolioRepository) Delete(userID, cryptoID string) error {
	key := userID + cryptoID
	if _, exists := m.UsersPortfolio[key]; !exists {
		return errors.New("User portfolio doesn't exists")
	}

	delete(m.UsersPortfolio, userID+cryptoID)

	return nil
}
