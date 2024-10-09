package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/devem-tech/echo/internal/types"
)

type Handler struct {
	log             log
	color           color
	routes          types.Routes
	responseLatency time.Duration
	prints          types.Print
}

func New(
	log log,
	color color,
	routes types.Routes,
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
		code:           0,
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
	// If found, return mock response
	if route, ok := h.routes.Get(r); ok {
		w.Header().Set("Content-Type", "application/json")

		if route.Code > 0 {
			w.WriteHeader(route.Code)
		}

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")

		if err := encoder.Encode(route.Content); err != nil {
			h.internalServerError(w, r, err)

			return
		}

		return
	}

	// If wildcard is found, use forwarding
	if forwardURL, ok := h.routes.ForwardURL(); ok {
		target, err := url.Parse(forwardURL)
		if err != nil {
			h.internalServerError(w, r, err)

			return
		}

		r.Host = target.Host
		httputil.NewSingleHostReverseProxy(target).ServeHTTP(w, r)

		return
	}

	// If not found, return 404 response
	h.notFound(w, r)
}

func (h *Handler) emulateResponseLatency() {
	if h.responseLatency == 0 {
		return
	}

	time.Sleep(h.responseLatency)
}

func (h *Handler) output(w ResponseWriter, r *http.Request) {
	code := fmt.Sprintf("[%d]", w.code)

	switch c := w.code; {
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
