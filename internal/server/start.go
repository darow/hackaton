package server

import (
	"hackaton/store"
	"net"
	"net/http"

	"go.uber.org/zap"
)

func Start(cfg *Config, logger *zap.SugaredLogger) error {
	s := store.New()

	serv := newServer(s, logger)

	return http.ListenAndServe(net.JoinHostPort(cfg.Host, cfg.Port), serv)
}
