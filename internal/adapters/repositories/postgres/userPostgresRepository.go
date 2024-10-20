package repositories

import (
	"github.com/stasdashkevitch/crypto_info/internal/database"
	"github.com/stasdashkevitch/crypto_info/internal/domain/entities"
)

type userPostgresRepository struct {
	db database.Database
}

func NewUserPostgresRepository(db database.Database) *userPostgresRepository {
	return &userPostgresRepository{
		db: db,
	}
}

func (r *userPostgresRepository) GetByUserName(username string) (*entities.User, error) {
	var user entities.User
	query := "select id, username, email, password_hash from users where username = ?"
	err := r.db.GetDB().QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userPostgresRepository) GetByEmail(email string) (*entities.User, error) {
	var user entities.User
	query := "select id, username, email, password_hash from users where email = ?"
	err := r.db.GetDB().QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userPostgresRepository) Create(user *entities.User) error {
	_, err := r.db.GetDB().Exec("insert into users (id, username, email, password_hash, created_at) values (?, ?, ?, ?, ?)", user.Username, user.Email, user.PasswordHash, user.CreatedAt)

	return err
}
