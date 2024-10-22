package main

import (
	"github.com/stasdashkevitch/crypto_info/internal/config"
	database "github.com/stasdashkevitch/crypto_info/internal/database/postgres"
	v1 "github.com/stasdashkevitch/crypto_info/internal/server/http/v1"
	"github.com/stasdashkevitch/crypto_info/pkg/env"
	"github.com/stasdashkevitch/crypto_info/pkg/logger"
)

func main() {
	env.LoadEnv()

	l := logger.GetLogger()
	defer l.Sync()
	l.Info("create config")

	cfg := config.NewConfig()
	l.Info(cfg.Server.Port)

	db := database.NewPostgresDatabase(cfg)

	v1.NewStandartHTTPServer(cfg, l, db).Start()

}
