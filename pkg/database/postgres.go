package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/stasdashkevitch/crypto_info/internal/config"
)

type postgresDatabase struct {
	Db *sql.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(cfg *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
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

		dbInstance = &postgresDatabase{
			Db: db,
		}
	})

	return dbInstance
}

func (db *postgresDatabase) GetDB() *sql.DB {
	return dbInstance.Db
}