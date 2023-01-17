package routing

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/devem-tech/echo/internal/types"
)

const (
	methodAndRoute        = 2
	methodAndCodeAndRoute = 3
)

func Parse(path string) (types.Routes, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	var routes map[string]any

	if err = json.Unmarshal(bytes, &routes); err != nil {
		return nil, err //nolint:wrapcheck
	}

	res := make(types.Routes, len(routes))
	for route, content := range routes {
		res.Add(parseRoute(route, content))
	}

	return res, nil
}

func parseRoute(route string, content any) *types.Route {
	if route == "*" {
		return &types.Route{
			Path:    "*",
			Content: content,
		}
	}

	parts := strings.Split(route, ":")

	if len(parts) == methodAndCodeAndRoute {
		return &types.Route{
			Method:  parts[0],
			Code:    integer(parts[1]),
			Path:    parts[2],
			Content: content,
		}
	}

	if len(parts) == methodAndRoute {
		i := integer(parts[0])

		method := parts[0]
		code := http.StatusOK

		if i > 0 {
			method = http.MethodGet
			code = i
		}

		return &types.Route{
			Method:  method,
			Code:    code,
			Path:    parts[1],
			Content: content,
		}
	}

	return &types.Route{
		Method:  http.MethodGet,
		Code:    http.StatusOK,
		Path:    route,
		Content: content,
	}
}

func integer(v string) int {
	i, _ := strconv.Atoi(v)

	return i
}
