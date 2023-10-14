package color

type Contract interface { //nolint:interfacebloat
	Black(string) string
	DarkGray(string) string
	Red(string) string
	LightRed(string) string
	Green(string) string
	LightGreen(string) string
	Brown(string) string
	Yellow(string) string
	Blue(string) string
	LightBlue(string) string
	Purple(string) string
	LightPurple(string) string
	Cyan(string) string
	LightCyan(string) string
	LightGray(string) string
	White(string) string
}
