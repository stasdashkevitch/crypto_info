package main

import (
	"github.com/stasdashkevitch/crypto_info/internal/config"
	"github.com/stasdashkevitch/crypto_info/internal/database"
	"github.com/stasdashkevitch/crypto_info/internal/logger"
	v1 "github.com/stasdashkevitch/crypto_info/internal/server/http/v1"
	"github.com/stasdashkevitch/crypto_info/pkg/util/env"
)

func main() {
	env.LoadEnv()

	l := logger.GetLogger()
	defer l.Sync()

	cfg := config.NewConfig()
	l.Info(cfg.Server.Port)

	db := database.NewPostgresDatabase(cfg)

	v1.NewStandartHTTPServer(cfg, l, db).Start()
}
