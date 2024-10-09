package handler

import "net/http"

type ResponseWriter struct {
	http.ResponseWriter
	code int
}

func (r *ResponseWriter) WriteHeader(code int) {
	r.ResponseWriter.WriteHeader(code)
	r.code = code
}
