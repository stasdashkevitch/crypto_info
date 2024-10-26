package v1

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/stasdashkevitch/crypto_info/internal/adapters/controller/http/v1/standart"
	"github.com/stasdashkevitch/crypto_info/internal/adapters/controller/websocket"
	redispubsub "github.com/stasdashkevitch/crypto_info/internal/adapters/pubsub/redisPubSub"
	"github.com/stasdashkevitch/crypto_info/internal/adapters/repository/postgres"
	"github.com/stasdashkevitch/crypto_info/internal/cache/redis"
	"github.com/stasdashkevitch/crypto_info/internal/config"
	"github.com/stasdashkevitch/crypto_info/internal/database"
	"github.com/stasdashkevitch/crypto_info/internal/usecase/auth"
	cryptodataprovider "github.com/stasdashkevitch/crypto_info/internal/usecase/cryptoDataProvider"
	cryptotrackerusecase "github.com/stasdashkevitch/crypto_info/internal/usecase/cryptoTrackerUsecase"
	loginusecase "github.com/stasdashkevitch/crypto_info/internal/usecase/loginUsecase"
	registrationusecase "github.com/stasdashkevitch/crypto_info/internal/usecase/registrationUsecase"
	"go.uber.org/zap"
)

type standartHTTPServer struct {
	server *http.Server
	db     database.Database
	cfg    *config.Config
	l      *zap.SugaredLogger
	cache  *redis.RedisDatabase
}

func (s *standartHTTPServer) Start() {
	go func() {
		s.l.Infof("Start tcp server in port %s", s.cfg.Server.Port)
		s.l.Fatal(s.server.ListenAndServe())
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill, os.Interrupt)

	sig := <-sigChan

	s.l.Info("Recieved terminate, gracefull shutdown ", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.db.GetDB().Close()
	if err != nil {
		s.l.Error("Database closing error: ", err)
	}

	err = s.cache.GetDB().Close()
	if err != nil {
		s.l.Error("Cache database closing error: ", err)
	}

	s.l.Info("Shutdown")
	s.server.Shutdown(tc)
}

func (s *standartHTTPServer) getServer() *http.Handler {
	return &s.server.Handler
}

func NewStandartHTTPServer(cfg *config.Config, l *zap.SugaredLogger, db database.Database, cache *redis.RedisDatabase) Server {
	sm := http.NewServeMux()

	regisеringRoute(sm, l, db, cache)

	s := &http.Server{
		Handler:      sm,
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}

	return &standartHTTPServer{
		server: s,
		db:     db,
		cfg:    cfg,
		l:      l,
		cache:  cache,
	}
}

func regisеringRoute(sm *http.ServeMux, l *zap.SugaredLogger, db database.Database, cache *redis.RedisDatabase) {
	userRepository := postgres.NewUserPostgresRepository(db)

	// login route
	auth := auth.NewJWTAuth()
	loginUsecase := loginusecase.NewLoginUsecase(auth, userRepository)
	standart.NewLoginHandler(sm, l, *loginUsecase)

	// registration route
	registrationUsecase := registrationusecase.NewRegistrationUsecase(userRepository)
	standart.NewRegistrationHandler(sm, l, *registrationUsecase)

	// crypto price info
	cryptoTrackerRedisPubSub := redispubsub.NewCryptoTrackerRedisPubSub(cache.GetDB())
	cryptoDataProvider := cryptodataprovider.NewCryptoDataFromCoinGeckoProvide()
	cryptoTrackerUsecase := cryptotrackerusecase.NewCryptoTrackerUsecase(cryptoDataProvider, cryptoTrackerRedisPubSub)
	websocket.NewCryptoTrackerWebsocketHandler(sm, l, *cryptoTrackerUsecase)

}
