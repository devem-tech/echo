package errors

const keyCause = "cause"

func With(key string, value any) func(*ContextError) {
	return func(x *ContextError) {
		x.Context[key] = value
	}
}

func Cause(cause error) func(*ContextError) {
	return func(x *ContextError) {
		x.Context[keyCause] = cause
	}
}
