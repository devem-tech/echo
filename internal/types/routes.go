package types

import (
	"fmt"
	"net/http"
)

const Wildcard = "*"

type Route struct {
	Method  string
	Path    string
	Code    int
	Content Content
}

type Content any // json || url

type Routes map[string]*Route

func (r Routes) Add(route *Route) {
	key := route.Method + ":" + route.Path
	if route.Path == Wildcard {
		key = Wildcard
	}

	r[key] = route
}

func (r Routes) Get(req *http.Request) (*Route, bool) {
	route, ok := r[req.Method+":"+req.URL.Path]

	return route, ok
}

func (r Routes) ForwardURL() (string, bool) {
	if route, ok := r[Wildcard]; ok {
		return fmt.Sprintf("%s", route.Content), true
	}

	return "", false
}
