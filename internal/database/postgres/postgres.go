package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/stasdashkevitch/crypto_info/internal/config"
	"github.com/stasdashkevitch/crypto_info/internal/database"
)

type postgresDatabase struct {
	Db *sql.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(cfg *config.Config) database.Database {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			cfg.DB.Host,
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.DBName,
			cfg.DB.Port,
			cfg.DB.SSLMode,
			cfg.DB.Timezone,
		)

		db, err := sql.Open("postgres", dsn)

		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY, 
    username TEXT NOT NULL, 
    email TEXT UNIQUE NOT NULL, 
    password_hash TEXT NOT NULL, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)`)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS portfolio_items (
    UserID UUID,
    CryptoID TEXT NOT NULL,
    IncreaseThreshold FLOAT,
    DecreaseThreshold FLOAT,
    NotifyIncrease BOOLEAN,
    NotifyDecrease BOOLEAN,
    NotificationMethod TEXT,
    PRIMARY KEY (UserID, CryptoID),
    FOREIGN KEY (UserID) REFERENCES users(id)
);`)

		if err != nil {
			panic(err)
		}

		dbInstance = &postgresDatabase{
			Db: db,
		}
	})

	return dbInstance
}

func (db *postgresDatabase) GetDB() *sql.DB {
	return dbInstance.Db
}
