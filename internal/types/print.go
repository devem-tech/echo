package types

import "strings"

type Print string

func (p Print) IsEmpty() bool {
	return p == ""
}

func (p Print) CanPrintRequest() bool {
	return p.CanPrintRequestHeaders() || p.CanPrintRequestBody()
}

func (p Print) CanPrintRequestHeaders() bool {
	return strings.Contains(string(p), "H")
}

func (p Print) CanPrintRequestBody() bool {
	return strings.Contains(string(p), "B")
}

func (p Print) CanPrintResponse() bool {
	return p.CanPrintResponseHeaders() || p.CanPrintResponseBody()
}

func (p Print) CanPrintResponseHeaders() bool {
	return strings.Contains(string(p), "h")
}

func (p Print) CanPrintResponseBody() bool {
	return strings.Contains(string(p), "b")
}
