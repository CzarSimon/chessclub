package main

import (
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/logger"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

var log = logger.GetDefaultLogger("game-server/main")

func main() {
	e := setupEnvironment()
	defer e.close()

	server := newServer(e)
	log.Info("Started game-server listening on port: " + e.cfg.Port)

	err := server.ListenAndServe()
	if err != nil {
		log.Error("Unexpected error stoped server.", zap.Error(err))
	}
}

func newServer(e *env) *http.Server {
	r := httputil.NewRouter("game-server", e.checkHealth)
	r.Use(httputil.AllowJSON())

	return &http.Server{
		Addr:    ":" + e.cfg.Port,
		Handler: r,
	}
}
