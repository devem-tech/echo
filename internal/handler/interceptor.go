package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"

	"github.com/devem-tech/echo/internal/compress/gzip"
	"github.com/devem-tech/echo/internal/compress/zlib"
	"github.com/devem-tech/echo/internal/logger"
	"github.com/devem-tech/echo/internal/types"
)

type Interceptor struct {
	w http.ResponseWriter
	r *http.Request
	l logger.Log
	c Color
	v types.Verbose
}

func (i *Interceptor) Handle(fn http.HandlerFunc) {
	if !i.v.IsVerbose() {
		fn(i.w, i.r)

		return
	}

	r, err := httputil.DumpRequest(i.r, true)
	if err != nil {
		i.l.Info(">>> failed to parse request: %v", err)

		return
	}

	i.l.Info(">>> %s %s\n%s", i.r.Method, i.r.URL.Path, r)

	rec := httptest.NewRecorder()

	fn(rec, i.r)

	i.l.Info("<<< %d %s %s%s\n%s", rec.Code, i.r.Method, i.r.URL.Path, i.headers(rec), i.decompress(rec))

	for k, v := range rec.Header() {
		i.w.Header()[k] = v
	}

	i.w.WriteHeader(rec.Code)

	_, _ = rec.Body.WriteTo(i.w)
}

func (i *Interceptor) headers(rec *httptest.ResponseRecorder) string {
	h := rec.Header()

	if i.v < types.VerbosityVeryVerbose || len(h) == 0 {
		return ""
	}

	var sb strings.Builder

	for k, values := range h {
		for _, v := range values {
			sb.WriteString(i.c.Cyan(k) + ": " + v + "\n")
		}
	}

	return "\n" + sb.String()
}

func (i *Interceptor) decompress(rw *httptest.ResponseRecorder) string {
	body := fmt.Sprintf("%v", rw.Body)

	switch rw.Header().Get("Content-Encoding") {
	case "gzip":
		return gzip.Decompress(body)
	case "deflate":
		return zlib.Decompress(body)
	default:
		return body
	}
}
