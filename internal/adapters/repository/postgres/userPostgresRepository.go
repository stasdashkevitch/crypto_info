package postgres

import (
	"github.com/stasdashkevitch/crypto_info/internal/database"
	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

type userPostgresRepository struct {
	db database.Database
}

func NewUserPostgresRepository(db database.Database) *userPostgresRepository {
	return &userPostgresRepository{
		db: db,
	}
}

func (r *userPostgresRepository) GetByUsername(username string) (*entity.User, error) {
	var user entity.User
	query := "select id, username, email, password_hash, created_at from users where username = $1"
	err := r.db.GetDB().QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userPostgresRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := "select id, username, email, password_hash, created_at from users where email = $1"
	err := r.db.GetDB().QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userPostgresRepository) Create(user *entity.User) error {
	_, err := r.db.GetDB().Exec("insert into users (id, username, email, password_hash, created_at) values ($1, $2, $3, $4, $5)", user.ID, user.Username, user.Email, user.PasswordHash, user.CreatedAt)

	return err
}
