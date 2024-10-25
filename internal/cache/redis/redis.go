package redis

import (
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/stasdashkevitch/crypto_info/internal/config"
)

type RedisDatabase struct {
	DB *redis.Client
}

var (
	once       sync.Once
	dbInstance *RedisDatabase
)

func NewRedisDatabase(cfg *config.Config) *RedisDatabase {
	once.Do(func() {
		addr := fmt.Sprintf("%s:%s", cfg.Cache.Host, cfg.Cache.Port)

		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: cfg.Cache.Password,
			DB:       cfg.Cache.DB,
		})

		dbInstance = &RedisDatabase{
			DB: client,
		}
	})

	return nil
}

func (db *RedisDatabase) GetDB() *redis.Client {
	return dbInstance.DB
}
