package postgres

import (
	"fmt"

	"github.com/pelletier/go-toml/query"
	"github.com/stasdashkevitch/crypto_info/internal/database"
	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

type userPortfolioPostgresRepository struct {
	db database.Database
}

func NewUserPortfolioPostgresRepository(db database.Database) *userPortfolioPostgresRepository {
	return &userPortfolioPostgresRepository{
		db: db,
	}
}

func (r *userPortfolioPostgresRepository) Create(userPortfolio *entity.UserPortfolio) error {
	query := `
  INSERT INTO  user_portfolio (
    user_id,
    crypto_id,
    increase_threshold,
    decrease_threshold,
    notify_increase,
    notify_decrease,
    notification_method)
  VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.GetDB().Exec(
		query,
		userPortfolio.UserID,
		userPortfolio.CryptoID,
		userPortfolio.IncreaseThreshold,
		userPortfolio.DecreaseThreshold,
		userPortfolio.NotifyIncrease,
		userPortfolio.NotifyDecrease,
		userPortfolio.NotificationMethod)

	if err != nil {
		return err
	}

	return nil
}

func (r *userPortfolioPostgresRepository) GetAllByUserID(userID string) ([]*entity.UserPortfolio, error) {
	var allUserPortfolio []*entity.UserPortfolio
	query := `
  SELECT 
    user_id, 
    crypto_id,
    increase_threshold,
    decrease_threshold,
    notify_increase,
    notify_decrease,
    notification_method
  FROM user_portfolio where
  WHERE user_id = $1`

	rows, err := r.db.GetDB().Query(query, userID)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userPortfolio entity.UserPortfolio

		err := rows.Scan(
			&userPortfolio.UserID,
			&userPortfolio.CryptoID,
			&userPortfolio.IncreaseThreshold,
			&userPortfolio.DecreaseThreshold,
			&userPortfolio.NotifyIncrease,
			&userPortfolio.NotifyDecrease,
			&userPortfolio.NotificationMethod,
		)
		if err != nil {
			return nil, err
		}

		allUserPortfolio = append(allUserPortfolio, &userPortfolio)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return allUserPortfolio, nil
}

func (r *userPortfolioPostgresRepository) GetByCryptoID(userID, cryptoID string) (*entity.UserPortfolio, error) {
	var userPortfolio entity.UserPortfolio
	query := `
  SELECT
    user_id, 
    crypto_id,
    increase_threshold,
    decrease_threshold,
    notify_increase,
    notify_decrease,
    notification_method
  FROM user_portfolio
  WHERE user_id = $1 AND crypto_id = $2`

	err := r.db.GetDB().QueryRow(query, userID, cryptoID).Scan(
		&userPortfolio.UserID,
		&userPortfolio.CryptoID,
		&userPortfolio.IncreaseThreshold,
		&userPortfolio.DecreaseThreshold,
		&userPortfolio.NotifyIncrease,
		&userPortfolio.NotifyDecrease,
		&userPortfolio.NotificationMethod)

	if err != nil {
		return nil, err
	}

	return &userPortfolio, nil
}

func (r *userPortfolioPostgresRepository) Update(userPortfolio *entity.UserPortfolio) error {
	query := `
  UPDATE user_portfolio SET 
		increase_threshold = $1,
		decrease_threshold = $2,
		notify_increase = $3,
		notify_decrease = $4,
		notification_method = $5
	WHERE user_id = $6 AND crypto_id = $7`

	_, err := r.db.GetDB().Exec(
		query,
		userPortfolio.IncreaseThreshold,
		userPortfolio.DecreaseThreshold,
		userPortfolio.NotifyIncrease,
		userPortfolio.NotifyDecrease,
		userPortfolio.NotificationMethod,
		userPortfolio.UserID,
		userPortfolio.CryptoID)

	if err != nil {
		return err
	}

	return nil
}

func (r *userPortfolioPostgresRepository) Delete(userID, cryptoID string) error {
	query := `
  DELETE FROM user_portfolio
  WHERE user_id = $1 AND crypto_id = $2
  `

	_, err := r.db.GetDB().Exec(query, userID, cryptoID)

	if err != nil {
		return err
	}

	return nil
}
