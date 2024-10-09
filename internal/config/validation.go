package config

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/devem-tech/echo/internal/types"
)

//nolint:gochecknoglobals
var (
	allowedMethods = map[string]struct{}{
		http.MethodGet:     {},
		http.MethodPost:    {},
		http.MethodPut:     {},
		http.MethodDelete:  {},
		http.MethodPatch:   {},
		http.MethodOptions: {},
		http.MethodHead:    {},
	}

	allowedCodes = map[int]struct{}{
		http.StatusOK:                  {},
		http.StatusCreated:             {},
		http.StatusAccepted:            {},
		http.StatusNoContent:           {},
		http.StatusMovedPermanently:    {},
		http.StatusFound:               {},
		http.StatusNotModified:         {},
		http.StatusBadRequest:          {},
		http.StatusUnauthorized:        {},
		http.StatusForbidden:           {},
		http.StatusNotFound:            {},
		http.StatusMethodNotAllowed:    {},
		http.StatusInternalServerError: {},
		http.StatusNotImplemented:      {},
		http.StatusBadGateway:          {},
		http.StatusServiceUnavailable:  {},
	}
)

func validateMethod(x string) (types.Method, bool) {
	x = strings.ToUpper(x)

	_, ok := allowedMethods[x]
	if !ok {
		return "", false
	}

	return types.Method(x), true
}

func validateEndpoint(x string) (types.Endpoint, bool) {
	uri, err := url.ParseRequestURI(x)
	if err != nil || uri.Path == "" || uri.Path == types.Wildcard {
		return "", false
	}

	return types.Endpoint(uri.Path), true
}

func validateCode(in string) (types.StatusCode, bool) {
	x, err := strconv.Atoi(in)
	if err != nil {
		return 0, false
	}

	_, ok := allowedCodes[x]
	if !ok {
		return 0, false
	}

	return types.StatusCode(x), true
}
