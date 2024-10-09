package model

import (
	"fmt"
	"net/http"

	"github.com/devem-tech/echo/internal/types"
	"github.com/devem-tech/echo/pkg/errors"
)

type Key string

type RouteMap map[Key]Route

func (m RouteMap) Add(route Route) error {
	key := Key(fmt.Sprintf("%s:%s", route.Request.Method, route.Request.Endpoint))
	if _, ok := m[key]; ok {
		return errors.New("duplicate", errors.With("key", key))
	}

	m[key] = route

	return nil
}

func (m RouteMap) Get(req *http.Request) (Route, bool) {
	key := Key(fmt.Sprintf("%s:%s", req.Method, req.URL.Path))

	route, ok := m[key]
	if ok {
		return route, true
	}

	key = Key(fmt.Sprintf("%s:%s", types.MethodWildcard, types.EndpointWildcard))

	route, ok = m[key]

	return route, ok
}
