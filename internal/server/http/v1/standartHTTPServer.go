package v1

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/stasdashkevitch/crypto_info/internal/config"
	"github.com/stasdashkevitch/crypto_info/internal/database"
	"go.uber.org/zap"
)

type standartHTTPServer struct {
	server *http.Server
	db     database.Database
	cfg    *config.Config
	l      *zap.SugaredLogger
}

func (s *standartHTTPServer) Start() {
	go func() {
		s.l.Infof("start tcp server in port %s", s.cfg.Server.Port)
		s.l.Fatal(s.server.ListenAndServe())
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill, os.Interrupt)

	sig := <-sigChan

	s.l.Info("recieved terminate, gracefull shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.db.GetDB().Close()
	if err != nil {
		s.l.Error("database closing error", err)
	}

	s.server.Shutdown(tc)
}

func NewStandartHTTPServer(cfg *config.Config, l *zap.SugaredLogger, db database.Database) Server {
	sm := http.NewServeMux()
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
	}
}
