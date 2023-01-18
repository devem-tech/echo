package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devemio/mockio/internal/color"
	"github.com/devemio/mockio/internal/logger"
	"github.com/devemio/mockio/internal/types"
)

type Handler struct {
	log     *logger.Log
	routes  types.Routes
	latency int
	c       *color.Color
}

func New(log *logger.Log, routes types.Routes, latency int, noColors bool) *Handler {
	return &Handler{
		log:     log,
		routes:  routes,
		latency: latency,
		c:       color.New(!noColors),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path

	res, ok := h.routes[path]
	if !ok {
		h.log.Info(h.c.Yellow("[404]")+" %s %s", r.Method, path)
		http.Error(w, fmt.Sprintf("route=%q not found", path), http.StatusNotFound)

		return
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		h.log.Info(h.c.LightRed("[500]")+" %s %s", r.Method, path)
		http.Error(w, fmt.Sprintf("failed to build response: %v", err), http.StatusInternalServerError)

		return
	}

	if h.latency > 0 {
		time.Sleep(time.Duration(h.latency) * time.Millisecond)
	}

	h.log.Info(h.c.LightGreen("[200]")+" %s %s", r.Method, path)
}
