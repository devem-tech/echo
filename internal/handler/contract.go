package handler

type log interface {
	Info(format string, v ...any)
}

type color interface {
	LightRed(x string) string
	LightGreen(x string) string
	Yellow(x string) string
	LightBlue(x string) string
	LightPurple(x string) string
	Cyan(x string) string
}
