package main

import (
	"fmt"

	"github.com/stasdashkevitch/crypto_info/internal/config"
	"github.com/stasdashkevitch/crypto_info/pkg/logger"
)

func main() {
	l := logger.GetLogger()
	defer l.Sync()

	l.Info("start")
	cfg := config.NewConfig()
	fmt.Println(cfg.Env)
}
