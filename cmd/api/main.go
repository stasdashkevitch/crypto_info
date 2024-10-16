package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/stasdashkevitch/crypto_info/internal/config"
	"github.com/stasdashkevitch/crypto_info/pkg/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	l := logger.GetLogger()
	defer l.Sync()

	l.Info("start")
	cfg := config.NewConfig()
	fmt.Println(cfg.Env)
}
