package errors

import (
	"fmt"
	"strconv"
	"strings"
)

type ContextError struct {
	Msg     string
	Context map[string]any
}

func New(msg string, opts ...func(*ContextError)) *ContextError {
	res := &ContextError{
		Msg:     msg,
		Context: make(map[string]any),
	}

	for _, opt := range opts {
		opt(res)
	}

	return res
}

func (e *ContextError) Error() string {
	if len(e.Context) == 0 {
		return e.Msg
	}

	return e.Msg + " (" + e.context() + ")"
}

func (e *ContextError) context() string {
	var sb strings.Builder

	for k, v := range e.Context {
		sb.WriteString(k)
		sb.WriteByte('=')

		switch val := v.(type) {
		case string:
			sb.WriteString(strconv.Quote(val))
		case error:
			sb.WriteString(strconv.Quote(val.Error()))
		default:
			sb.WriteString(fmt.Sprintf("%v", val))
		}

		sb.WriteByte(' ')
	}

	return sb.String()[:sb.Len()-1]
}
