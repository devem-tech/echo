package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devemio/mockio/internal/logger"
	"github.com/devemio/mockio/internal/types"
	"github.com/devemio/mockio/pkg/color"
)

type Handler struct {
	log     *logger.Log
	routes  types.Routes
	latency int
}

func New(log *logger.Log, routes types.Routes, latency int) *Handler {
	return &Handler{
		log:     log,
		routes:  routes,
		latency: latency,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path

	res, ok := h.routes[path]
	if !ok {
		h.log.Info(color.Yellow("[404]")+" %s %s", r.Method, path)
		http.Error(w, fmt.Sprintf("route=%q not found", path), http.StatusNotFound)

		return
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		h.log.Info(color.LightRed("[500]")+" %s %s", r.Method, path)
		http.Error(w, fmt.Sprintf("failed to build response: %v", err), http.StatusInternalServerError)

		return
	}

	if h.latency > 0 {
		time.Sleep(time.Duration(h.latency) * time.Millisecond)
	}

	h.log.Info(color.LightGreen("[200]")+" %s %s", r.Method, path)
}
