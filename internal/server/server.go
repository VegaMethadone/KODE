package server

import (
	"fmt"
	"kode/config"
	"kode/internal/server/handlers"
	"net/http"
	"time"
)

func NewServer(conf *config.Config) *http.Server {

	srv := &http.Server{
		Handler:      handlers.GetRoutes(),
		Addr:         fmt.Sprintf("%s:%s", conf.Network.Ip, conf.Network.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}
