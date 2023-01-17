package routing //nolint:testpackage

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/devem-tech/echo/internal/types"
)

func Test_parseRoute(t *testing.T) { //nolint:funlen
	t.Parallel()

	tests := []struct {
		name string
		in   string
		want *types.Route
	}{
		{
			name: "full",
			in:   "POST:201:/v1/orders",
			want: &types.Route{
				Method: http.MethodPost,
				Code:   http.StatusCreated,
				Path:   "/v1/orders",
			},
		},
		{
			name: "default",
			in:   "/health",
			want: &types.Route{
				Method: http.MethodGet,
				Code:   200,
				Path:   "/health",
			},
		},
		{
			name: "no_method",
			in:   "401:/products",
			want: &types.Route{
				Method: http.MethodGet,
				Code:   http.StatusUnauthorized,
				Path:   "/products",
			},
		},
		{
			name: "no_code",
			in:   "GET:/v1/orders/10",
			want: &types.Route{
				Method: http.MethodGet,
				Code:   http.StatusOK,
				Path:   "/v1/orders/10",
			},
		},
		{
			name: "wildcard",
			in:   "*",
			want: &types.Route{
				Path: "*",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := parseRoute(tt.in, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}
