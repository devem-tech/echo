package routing_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/devem-tech/echo/internal/routing"
	"github.com/devem-tech/echo/internal/types"
)

func Test_parseRoute(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		in   string
		want *types.Route
	}{
		"full": {
			in: "POST:201:/v1/orders",
			want: &types.Route{
				Method:  http.MethodPost,
				Code:    http.StatusCreated,
				Path:    "/v1/orders",
				Content: nil,
			},
		},
		"default": {
			in: "/health",
			want: &types.Route{
				Method:  http.MethodGet,
				Code:    200,
				Path:    "/health",
				Content: nil,
			},
		},
		"no_method": {
			in: "401:/products",
			want: &types.Route{
				Method:  http.MethodGet,
				Code:    http.StatusUnauthorized,
				Path:    "/products",
				Content: nil,
			},
		},
		"no_code": {
			in: "GET:/v1/orders/10",
			want: &types.Route{
				Method:  http.MethodGet,
				Code:    http.StatusOK,
				Path:    "/v1/orders/10",
				Content: nil,
			},
		},
		"wildcard": {
			in: "*",
			want: &types.Route{
				Method:  "",
				Code:    0,
				Path:    "*",
				Content: nil,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := routing.ParseRoute(tt.in, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
