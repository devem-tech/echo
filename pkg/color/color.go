package color

const (
	reset       = "\033[0m"
	black       = "\033[0;30m"
	darkGray    = "\033[1;30m"
	red         = "\033[0;31m"
	lightRed    = "\033[1;31m"
	green       = "\033[0;32m"
	lightGreen  = "\033[1;32m"
	brown       = "\033[0;33m"
	yellow      = "\033[1;33m"
	blue        = "\033[0;34m"
	lightBlue   = "\033[1;34m"
	purple      = "\033[0;35m"
	lightPurple = "\033[1;35m"
	cyan        = "\033[0;36m"
	lightCyan   = "\033[1;36m"
	lightGray   = "\033[0;37m"
	white       = "\033[1;37m"
)

func Black(x string) string {
	return black + x + reset
}

func DarkGray(x string) string {
	return darkGray + x + reset
}

func Red(x string) string {
	return red + x + reset
}

func LightRed(x string) string {
	return lightRed + x + reset
}

func Green(x string) string {
	return green + x + reset
}

func LightGreen(x string) string {
	return lightGreen + x + reset
}

func Brown(x string) string {
	return brown + x + reset
}

func Yellow(x string) string {
	return yellow + x + reset
}

func Blue(x string) string {
	return blue + x + reset
}

func LightBlue(x string) string {
	return lightBlue + x + reset
}

func Purple(x string) string {
	return purple + x + reset
}

func LightPurple(x string) string {
	return lightPurple + x + reset
}

func Cyan(x string) string {
	return cyan + x + reset
}

func LightCyan(x string) string {
	return lightCyan + x + reset
}

func LightGray(x string) string {
	return lightGray + x + reset
}

func White(x string) string {
	return white + x + reset
}
