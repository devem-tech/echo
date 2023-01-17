package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/devem-tech/echo/internal/logger"
	"github.com/devem-tech/echo/internal/types"
)

type Handler struct {
	l  logger.Log
	c  Color
	r  types.Routes
	rl int
	v  types.Verbose
}

func New(log logger.Log, color Color, routes types.Routes, responseLatency int, verbose types.Verbose) *Handler {
	return &Handler{l: log, c: color, r: routes, rl: responseLatency, v: verbose}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw := &ResponseWriter{ResponseWriter: w}
	i := &Interceptor{rw, r, h.l, h.c, h.v}

	i.Handle(h.handle)
	h.latency()
	h.log(rw, r)
}

func (h *Handler) handle(w http.ResponseWriter, r *http.Request) {
	// Get mock response
	res := h.r.Get(r)

	// If found, return mock response
	if res != nil {
		w.Header().Set("Content-Type", "application/json")

		if res.Code > 0 {
			w.WriteHeader(res.Code)
		}

		if err := json.NewEncoder(w).Encode(res.Content); err != nil {
			h.failedToBuildResponse(w, r, err)

			return
		}

		return
	}

	// If wildcard is found, use forwarding
	if forwardURL := h.r.GetForwardURL(); forwardURL != "" {
		u, err := url.Parse(forwardURL)
		if err != nil {
			h.failedToBuildResponse(w, r, err)

			return
		}

		r.Host = u.Host
		httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)

		return
	}

	// If not found, return 404 response
	h.routeNotFound(w, r)
}

func (h *Handler) latency() {
	if h.rl <= 0 {
		return
	}

	time.Sleep(time.Duration(h.rl) * time.Millisecond)
}

func (h *Handler) log(w *ResponseWriter, r *http.Request) {
	code := fmt.Sprintf("[%d]", w.code)

	switch c := w.code; {
	case c >= http.StatusOK && c < http.StatusMultipleChoices:
		code = h.c.LightGreen(code)
	case c >= http.StatusMultipleChoices && c < http.StatusInternalServerError:
		code = h.c.Yellow(code)
	default:
		code = h.c.LightRed(code)
	}

	h.l.Info("%s %s %s", code, r.Method, r.URL.Path)
}

func (h *Handler) routeNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("route=%q not found", r.URL.Path), http.StatusNotFound)
}

func (h *Handler) failedToBuildResponse(w http.ResponseWriter, _ *http.Request, err error) {
	http.Error(w, fmt.Sprintf("failed to build response: %v", err), http.StatusInternalServerError)
}
