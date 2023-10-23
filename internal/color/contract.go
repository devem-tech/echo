package color

type Contract interface { //nolint:interfacebloat
	Black(x string) string
	DarkGray(x string) string
	Red(x string) string
	LightRed(x string) string
	Green(x string) string
	LightGreen(x string) string
	Brown(x string) string
	Yellow(x string) string
	Blue(x string) string
	LightBlue(x string) string
	Purple(x string) string
	LightPurple(x string) string
	Cyan(x string) string
	LightCyan(x string) string
	LightGray(x string) string
	White(x string) string
}
