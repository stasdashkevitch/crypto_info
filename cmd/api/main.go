package main

import (
	"github.com/stasdashkevitch/crypto_info/internal/config"
	v1 "github.com/stasdashkevitch/crypto_info/internal/server/http/v1"
	"github.com/stasdashkevitch/crypto_info/pkg/database"
	"github.com/stasdashkevitch/crypto_info/pkg/logger"
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
