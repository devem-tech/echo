package logger

type Contract interface {
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Fatal(err error)
}
