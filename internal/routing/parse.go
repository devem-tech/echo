package routing

import (
	"encoding/json"
	"fmt"
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
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var routes map[string]types.Content
	if err = json.Unmarshal(bytes, &routes); err != nil {
		return nil, fmt.Errorf("unmarshal file: %w", err)
	}

	res := make(types.Routes, len(routes))
	for route, content := range routes {
		res.Add(ParseRoute(route, content))
	}

	return res, nil
}

func ParseRoute(route string, content types.Content) *types.Route {
	if route == types.Wildcard {
		return &types.Route{
			Method:  "",
			Path:    types.Wildcard,
			Code:    0,
			Content: content,
		}
	}

	parts := strings.Split(route, ":")

	if len(parts) == fullRouteLen {
		return &types.Route{
			Method:  parts[0],
			Path:    parts[2],
			Code:    integer(parts[1]),
			Content: content,
		}
	}

	if len(parts) == partialRouteLen {
		firstPart := integer(parts[0])

		method := parts[0]
		code := http.StatusOK

		// If the first part is the code
		if firstPart > 0 {
			method = http.MethodGet
			code = firstPart
		}

		return &types.Route{
			Method:  method,
			Path:    parts[1],
			Code:    code,
			Content: content,
		}
	}

	return &types.Route{
		Method:  http.MethodGet,
		Path:    route,
		Code:    http.StatusOK,
		Content: content,
	}
}

func integer(v string) int {
	res, err := strconv.Atoi(v)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf(">>> (suppressed) int: cast: %w", err))

		return 0
	}

	return res
}
