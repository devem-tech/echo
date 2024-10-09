package handler

import (
	"bytes"
	"io"
	"net/http"
)

func (i *Interceptor) request() {
	if !i.prints.CanPrintRequest() {
		return
	}

	i.log.Info("%s %s %s%s%s", i.color.LightBlue(">>>"), i.r.Method, i.r.URL.Path, i.H(), i.B())
}

func (i *Interceptor) H() string {
	if !i.prints.CanPrintRequestHeaders() {
		return ""
	}

	headers := i.r.Header
	headers["Host"] = []string{
		i.r.Host,
	}

	return i.headers(headers)
}

func (i *Interceptor) B() string {
	if !i.prints.CanPrintRequestBody() {
		return ""
	}

	if i.r.Body == http.NoBody {
		return ""
	}

	body, err := io.ReadAll(i.r.Body)
	if err != nil {
		i.log.Info(">>> failed to read request: %v", err)

		return ""
	}

	i.r.Body = io.NopCloser(bytes.NewBuffer(body))

	return "\n" + string(body)
}
