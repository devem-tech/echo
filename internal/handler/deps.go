package handler

type Color interface {
	LightRed(v string) string
	LightGreen(v string) string
	Yellow(v string) string
	Cyan(v string) string
}
