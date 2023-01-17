package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devem-tech/echo/internal/color"
	"github.com/devem-tech/echo/internal/config"
	"github.com/devem-tech/echo/internal/handler"
	"github.com/devem-tech/echo/internal/logger"
	"github.com/devem-tech/echo/internal/routing"
)

const shutdownTimeout = 10 * time.Second

func main() {
	cfg := config.New()
	clr := color.New(cfg.IsOutputColored)
	log := logger.New(clr)

	routes, err := routing.Parse(cfg.Path)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Starting server at port %d...", cfg.Port)
	log.Debug("Creating routes")

	for _, route := range routes {
		method := route.Method
		if route.Path == "*" {
			method = "*"
		}

		log.Debug("%-6s %s", method, route.Path)
	}

	srv := &http.Server{ //nolint:gosec
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler.New(log, clr, routes, cfg.ResponseLatency, cfg.Verbose),
	}

	go func() {
		if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	log.Debug("Server started (version %s)", config.Version)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Debug("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal(err) //nolint:gocritic
	}

	log.Debug("Graceful shutdown complete")
}
