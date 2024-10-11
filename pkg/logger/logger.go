package logger

import (
	"os"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	config := getConfig()
	logger := zap.Must(config.Build())
	defer logger.Sync()

	logger.Sugar()

	return logger
}

func getConfig() *zap.Config {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	if os.Getenv("CRYPTO_INFO_ENV") == "prod" {
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	return &config
}
