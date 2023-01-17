package routing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/devem-tech/echo/internal/types"
)

const (
	partialRouteLen = 2
	fullRouteLen    = 3
)

func Parse(path string) (types.Routes, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var routes map[string]types.Content

	if err = json.Unmarshal(bytes, &routes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
	}

	res := make(types.Routes, len(routes))
	for route, content := range routes {
		res.Add(ParseRoute(route, content))
	}

	return res, nil
}

func ParseRoute(route string, content types.Content) *types.Route {
	if route == "*" {
		return &types.Route{
			Method:  "",
			Code:    0,
			Path:    "*",
			Content: content,
		}
	}

	parts := strings.Split(route, ":")

	if len(parts) == fullRouteLen {
		return &types.Route{
			Method:  parts[0],
			Code:    integer(parts[1]),
			Path:    parts[2],
			Content: content,
		}
	}

	if len(parts) == partialRouteLen {
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
