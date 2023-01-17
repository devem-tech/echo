package types

import (
	"fmt"
	"net/http"
)

type Route struct {
	Method  string
	Code    int
	Path    string
	Content any
}

type Routes map[string]*Route

func (r Routes) Get(req *http.Request) *Route {
	return r[req.Method+":"+req.URL.Path]
}

func (r Routes) Add(route *Route) {
	k := route.Method + ":" + route.Path
	if route.Path == "*" {
		k = "*"
	}

	r[k] = route
}

func (r Routes) GetForwardURL() string {
	if _, found := r["*"]; found {
		return fmt.Sprintf("%s", r["*"].Content)
	}

	return ""
}
