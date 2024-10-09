package model

import "github.com/devem-tech/echo/internal/types"

type Request struct {
	Method   types.Method
	Endpoint types.Endpoint
}

type Response struct {
	StatusCode types.StatusCode
	Body       *types.Body
	ProxyURL   *types.ProxyURL
}

type Route struct {
	Request  Request
	Response Response
}
