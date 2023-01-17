package types

import (
	"fmt"
	"net/http"
)

type Content any

type Route struct {
	Method  string
	Code    int
	Path    string
	Content Content
}

type Routes map[string]*Route

func (r Routes) Get(req *http.Request) *Route {
	return r[req.Method+":"+req.URL.Path]
}

func (r Routes) Add(route *Route) {
	key := route.Method + ":" + route.Path
	if route.Path == "*" {
		key = "*"
	}

	r[key] = route
}

func (r Routes) GetForwardURL() string {
	if _, contains := r["*"]; contains {
		return fmt.Sprintf("%s", r["*"].Content)
	}

	return ""
}
