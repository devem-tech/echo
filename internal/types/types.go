package types

import (
	"encoding/json"
	"net/url"
)

const Wildcard = "*"

type Method string

const MethodWildcard = Wildcard

type Endpoint string

const EndpointWildcard = Wildcard

type StatusCode int

const StatusCodeUndefined = -1

type (
	Body     = json.RawMessage
	ProxyURL = url.URL
)
