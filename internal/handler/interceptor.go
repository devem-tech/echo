package handler

import (
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"

	"github.com/devem-tech/echo/internal/types"
)

type Interceptor struct {
	rw     http.ResponseWriter
	r      *http.Request
	log    log
	color  color
	prints types.Print
}

func (i *Interceptor) handle(fn http.HandlerFunc) {
	if i.prints.IsEmpty() {
		fn(i.rw, i.r)

		return
	}

	rw := httptest.NewRecorder()

	i.request()
	fn(rw, i.r)
	i.response(rw)

	for k, v := range rw.Header() {
		i.rw.Header()[k] = v
	}

	i.rw.WriteHeader(rw.Code)

	_, _ = rw.Body.WriteTo(i.rw)
}

func (i *Interceptor) headers(headers http.Header) string {
	if len(headers) == 0 {
		return ""
	}

	var res strings.Builder

	keys := make([]string, 0, len(headers))
	for key := range headers {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		for _, value := range headers[key] {
			res.WriteString("\n" + i.color.Cyan(key) + ": " + value)
		}
	}

	return res.String()
}
