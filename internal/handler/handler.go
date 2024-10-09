package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/devem-tech/echo/internal/model"
	"github.com/devem-tech/echo/internal/types"
	"github.com/devem-tech/echo/pkg/errors"
)

type Handler struct {
	log             log
	color           color
	routes          model.RouteMap
	responseLatency time.Duration
	prints          types.Print
}

func New(
	log log,
	color color,
	routes model.RouteMap,
	responseLatency time.Duration,
	prints types.Print,
) *Handler {
	return &Handler{
		log:             log,
		color:           color,
		routes:          routes,
		responseLatency: responseLatency,
		prints:          prints,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rw := ResponseWriter{
		ResponseWriter: w,
		statusCode:     0,
	}

	interceptor := &Interceptor{
		rw:     &rw,
		r:      r,
		log:    h.log,
		color:  h.color,
		prints: h.prints,
	}

	interceptor.handle(h.handle)

	h.emulateResponseLatency()

	h.output(rw, r)
}

func (h *Handler) handle(w http.ResponseWriter, r *http.Request) {
	// If not found, return 404
	route, ok := h.routes.Get(r)
	if !ok {
		h.notFound(w, r)

		return
	}

	// Return 200
	if route.Response.Body != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(int(route.Response.StatusCode))

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")

		if err := encoder.Encode(*route.Response.Body); err != nil {
			h.internalServerError(w, r, err)

			return
		}

		return
	}

	// Reverse proxy
	if route.Response.ProxyURL != nil {
		r.Host = route.Response.ProxyURL.Host
		httputil.NewSingleHostReverseProxy(route.Response.ProxyURL).ServeHTTP(w, r)

		return
	}

	h.internalServerError(w, r, errors.New("no response found"))
}

func (h *Handler) emulateResponseLatency() {
	if h.responseLatency == 0 {
		return
	}

	time.Sleep(h.responseLatency)
}

func (h *Handler) output(w ResponseWriter, r *http.Request) {
	code := fmt.Sprintf("[%d]", w.statusCode)

	switch c := w.statusCode; {
	case c >= http.StatusOK && c < http.StatusMultipleChoices:
		code = h.color.LightGreen(code)
	case c >= http.StatusMultipleChoices && c < http.StatusInternalServerError:
		code = h.color.Yellow(code)
	default:
		code = h.color.LightRed(code)
	}

	h.log.Info("%s %s %s", code, r.Method, r.URL.Path)
}

func (h *Handler) notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Route=%q not found", r.URL.Path), http.StatusNotFound)
}

func (h *Handler) internalServerError(w http.ResponseWriter, _ *http.Request, err error) {
	http.Error(w, fmt.Sprintf("Failed: %v", err), http.StatusInternalServerError)
}
