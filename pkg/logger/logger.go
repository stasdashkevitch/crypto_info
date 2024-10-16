package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
)

const (
	logFilePath = "logs/out.log"
)

var loggerInstance *zap.SugaredLogger
var once sync.Once

func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		config := getConfig()
		logger := zap.Must(config.Build())

		loggerInstance = logger.Sugar()
	})

	return loggerInstance
}

func getConfig() *zap.Config {
	var config zap.Config
	if os.Getenv("CRYPTO_INFO_ENV") == "prod" {
		config = zap.NewProductionConfig()
		config.OutputPaths = []string{"stdout", logFilePath}
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	} else {
		config = zap.NewDevelopmentConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	return &config
}
