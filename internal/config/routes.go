package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/devem-tech/echo/internal/model"
	"github.com/devem-tech/echo/internal/types"
	"github.com/devem-tech/echo/pkg/errors"
)

//nolint:revive,stylecheck
const (
	len_Wildcard                   = 1
	len_Method_Endpoint            = 2
	len_Method_Endpoint_StatusCode = 3
)

func ReadRouteFile(filepath string) (model.RouteMap, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, errors.New(
			"file: read",
			errors.Cause(err),
			errors.With("filepath", filepath),
		)
	}

	var routes map[string]json.RawMessage
	if err = json.Unmarshal(bytes, &routes); err != nil {
		return nil, errors.New(
			"file: unmarshal",
			errors.Cause(err),
			errors.With("filepath", filepath),
		)
	}

	res := make(model.RouteMap, len(routes))

	for req, resp := range routes {
		route, err := ParseRoute(req, resp)
		if err != nil {
			return nil, fmt.Errorf("route: parse: %w", err)
		}

		if err = res.Add(*route); err != nil {
			return nil, fmt.Errorf("route: add: %w", err)
		}
	}

	return res, nil
}

//nolint:funlen,cyclop
func ParseRoute(req string, resp json.RawMessage) (*model.Route, error) {
	// Default route
	res := &model.Route{
		Request: model.Request{
			Method:   types.MethodWildcard,
			Endpoint: types.EndpointWildcard,
		},
		Response: model.Response{
			StatusCode: types.StatusCodeUndefined,
			Body:       nil,
			ProxyURL:   nil,
		},
	}

	// Make sure the route is not empty
	if req == "" {
		return nil, errors.New("must not be empty", errors.With("route", req))
	}

	// Make sure the route length is correct
	parts := strings.Split(req, ":")
	if len(parts) > len_Method_Endpoint_StatusCode {
		return nil, errors.New("must have 3 parts", errors.With("route", req))
	}

	var ok bool

	switch len(parts) {
	// Format <METHOD>:<ENDPOINT>:<STATUS_CODE>
	case len_Method_Endpoint_StatusCode:
		if res.Request.Method, ok = validateMethod(parts[0]); !ok {
			return nil, errors.New(
				"invalid method",
				errors.With("route", req),
				errors.With("method", parts[0]),
			)
		}

		if res.Request.Endpoint, ok = validateEndpoint(parts[1]); !ok {
			return nil, errors.New(
				"invalid endpoint",
				errors.With("route", req),
				errors.With("endpoint", parts[1]),
			)
		}

		if res.Response.StatusCode, ok = validateCode(parts[2]); !ok {
			return nil, errors.New(
				"invalid status code",
				errors.With("route", req),
				errors.With("status_code", parts[2]),
			)
		}

	// Format <METHOD>:<ENDPOINT>
	case len_Method_Endpoint:
		if res.Request.Method, ok = validateMethod(parts[0]); !ok {
			return nil, errors.New(
				"invalid method",
				errors.With("route", req),
				errors.With("method", parts[0]),
			)
		}

		if res.Request.Endpoint, ok = validateEndpoint(parts[1]); !ok {
			return nil, errors.New(
				"invalid endpoint",
				errors.With("route", req),
				errors.With("endpoint", parts[1]),
			)
		}

	// Wildcard
	case len_Wildcard:
		if parts[0] != types.Wildcard {
			return nil, errors.New("must be wildcard", errors.With("route", req))
		}
	}

	// Assume it's a proxy url
	uri, err := url.ParseRequestURI(strings.Trim(string(resp), "\""))
	if err == nil && uri.Scheme != "" && uri.Host != "" {
		res.Response.ProxyURL = uri

		if res.Response.StatusCode != types.StatusCodeUndefined {
			return nil, errors.New(
				"status code must be omitted",
				errors.With("route", req),
				errors.With("proxy_url", strings.Trim(string(resp), "\"")),
			)
		}

		return res, nil
	}

	// Assume it's s json response
	var x types.Body
	if err = json.Unmarshal(resp, &x); err == nil {
		res.Response.Body = &x

		if res.Response.StatusCode == types.StatusCodeUndefined {
			res.Response.StatusCode = http.StatusOK
		}

		return res, nil
	}

	// Looks like the http response is invalid
	return nil, errors.New(
		"invalid response format",
		errors.Cause(err),
		errors.With("response", string(resp)),
	)
}
