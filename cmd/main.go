package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devemio/mockio/internal/config"
	"github.com/devemio/mockio/internal/handler"
	"github.com/devemio/mockio/internal/logger"
	"github.com/devemio/mockio/internal/types"
)

func main() {
	log, cfg := logger.New(), config.New()

	routes, err := parse(cfg.Path)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Starting server at port %s...", cfg.Port)

	for route := range routes {
		log.Debug("Creating route %q", route)
	}

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler.New(log, routes, cfg.Latency),
	}

	go func() {
		if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	log.Debug("Server started (version 0.2.0)")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Debug("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Debug("Graceful shutdown complete")
}

func parse(path string) (types.Routes, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var res types.Routes

	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
