package logger

type Color interface {
	Red(v string) string
	LightGray(v string) string
}
