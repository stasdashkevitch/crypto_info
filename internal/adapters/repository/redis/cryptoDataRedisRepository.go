package redis

import (
	"context"
	"encoding/json"

	"github.com/stasdashkevitch/crypto_info/internal/cache/redis"
	"github.com/stasdashkevitch/crypto_info/internal/entity"
)

type cryptoDataRedisRepository struct {
	db *redis.RedisDatabase
}

func NewCryptoDataRedisRepository(db *redis.RedisDatabase) *cryptoDataRedisRepository {
	return &cryptoDataRedisRepository{
		db: db,
	}
}

func (r *cryptoDataRedisRepository) SetCryptoData(ctx context.Context, data *entity.CryptoData) error {
	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return r.db.GetDB().Set(ctx, data.ID, json, 0).Err()
}

func (r *cryptoDataRedisRepository) GetCrpytoData(ctx context.Context, id string) (*entity.CryptoData, error) {
	val, err := r.db.GetDB().Get(ctx, id).Result()

	if err != nil {
		return nil, err
	}

	var data entity.CryptoData

	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *cryptoDataRedisRepository) UpdateCryptoData(ctx context.Context, data *entity.CryptoData) error {
	json, err := json.Marshal(data)

	if err != nil {
		return err
	}
	return r.db.GetDB().Set(ctx, data.ID, json, 0).Err()
}
