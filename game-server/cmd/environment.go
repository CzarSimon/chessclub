package main

import (
	"database/sql"
	"io"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/dbutil"
	"github.com/czarsimon/chessclub/game-server/internal/config"
	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

type env struct {
	cfg         config.Config
	db          *sql.DB
	traceCloser io.Closer
}

func setupEnvironment() *env {
	cfg := config.GetConfig()
	db := connectDB(cfg)

	return &env{
		cfg:         cfg,
		db:          db,
		traceCloser: setupTracer(),
	}
}

func (e *env) checkHealth() error {
	err := dbutil.Connected(e.db)
	if err != nil {
		return httputil.ServiceUnavailableError(err)
	}

	return nil
}

func (e *env) close() {
	err := e.db.Close()
	if err != nil {
		log.Error("failed to close database connection", zap.Error(err))
	}

	err = e.traceCloser.Close()
	if err != nil {
		log.Error("failed to close tracer", zap.Error(err))
	}
}

func setupTracer() io.Closer {
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		log.Fatal("failed to create jaeger configuration", zap.Error(err))
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Fatal("failed to create tracer", zap.Error(err))
	}

	opentracing.SetGlobalTracer(tracer)

	return closer
}

func connectDB(cfg config.Config) *sql.DB {
	log.Info("connecting to database", zap.String("driver", cfg.DB.Driver()))
	db := dbutil.MustConnect(cfg.DB)

	log.Info("applying database migrations", zap.String("migrationsPath", cfg.MigrationsPath), zap.String("driver", cfg.DB.Driver()))
	err := dbutil.Upgrade(cfg.MigrationsPath, cfg.DB.Driver(), db)
	if err != nil {
		log.Fatal("failed to apply database migrations", zap.Error(err))
	}

	return db
}
