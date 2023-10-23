package handler

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"net/http/httptest"
	"strings"

	"github.com/klauspost/compress/zstd"

	"github.com/devem-tech/echo/pkg/io"
)

func (i *Interceptor) response(rw *httptest.ResponseRecorder) {
	if !i.prints.CanPrintResponse() {
		return
	}

	i.log.Info("%s %d %s %s%s%s", i.color.LightPurple("<<<"), rw.Code, i.r.Method, i.r.URL.Path, i.h(rw), i.b(rw))
}

func (i *Interceptor) h(rw *httptest.ResponseRecorder) string {
	if !i.prints.CanPrintResponseHeaders() {
		return ""
	}

	return i.headers(rw.Header())
}

func (i *Interceptor) b(rw *httptest.ResponseRecorder) string {
	if !i.prints.CanPrintResponseBody() {
		return ""
	}

	body := fmt.Sprintf("%v", rw.Body)

	switch rw.Header().Get("Content-Encoding") {
	case "gzip":
		body = io.String(gzip.NewReader(strings.NewReader(body)))
	case "deflate":
		body = io.String(zlib.NewReader(strings.NewReader(body)))
	case "zstd":
		body = io.String(zstd.NewReader(strings.NewReader(body)))
	}

	return "\n" + body
}
