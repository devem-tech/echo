package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"syscall"
	"time"

	"github.com/devem-tech/echo/internal/color"
	"github.com/devem-tech/echo/internal/config"
	"github.com/devem-tech/echo/internal/handler"
	"github.com/devem-tech/echo/internal/logger"
	"github.com/devem-tech/echo/internal/model"
)

const (
	shutdownTimeout   = 10 * time.Second
	readHeaderTimeout = 30 * time.Second
)

func main() {
	cfg := config.New()
	clr := color.New(cfg.IsOutputColored)
	log := logger.New(clr, cfg.IsVerbose)

	routes, err := config.ReadRouteFile(cfg.Filepath)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Starting server at port %d...", cfg.Port)
	log.Debug("Creating routes...")

	for _, route := range sorted(routes) {
		log.Debug("%-6s %s", route.Request.Method, route.Request.Endpoint)
	}

	//nolint:exhaustivestruct,exhaustruct
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           handler.New(log, clr, routes, cfg.ResponseLatency, cfg.Print),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	go func() {
		if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	log.Info("Started server (port=%s version=%s)", clr.Purple(strconv.Itoa(cfg.Port)), clr.Purple(config.Version))

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

func sorted(routes model.RouteMap) []model.Route {
	res := make([]model.Route, 0, len(routes))
	for _, route := range routes {
		res = append(res, route)
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].Request.Endpoint != res[j].Request.Endpoint {
			return res[i].Request.Endpoint < res[j].Request.Endpoint
		}

		return res[i].Request.Method < res[j].Request.Method
	})

	return res
}
