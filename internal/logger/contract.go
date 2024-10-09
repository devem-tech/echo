package logger

type color interface {
	Red(x string) string
	LightGray(x string) string
}
