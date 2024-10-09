package color

import (
	"os"
)

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

type Color struct {
	isEnabled   bool
	isTrueColor bool
}

func New(isEnabled bool) *Color {
	return &Color{
		isEnabled:   isEnabled,
		isTrueColor: os.Getenv("COLORTERM") == "truecolor",
	}
}

func (c *Color) Black(x string) string {
	if !c.isEnabled {
		return x
	}

	return black + x + reset
}

func (c *Color) DarkGray(x string) string {
	if !c.isEnabled {
		return x
	}

	return darkGray + x + reset
}

func (c *Color) Red(x string) string {
	if !c.isEnabled {
		return x
	}

	return red + x + reset
}

func (c *Color) LightRed(x string) string {
	if !c.isEnabled {
		return x
	}

	return lightRed + x + reset
}

func (c *Color) Green(x string) string {
	if !c.isEnabled {
		return x
	}

	return green + x + reset
}

func (c *Color) LightGreen(x string) string {
	if !c.isEnabled {
		return x
	}

	return lightGreen + x + reset
}

func (c *Color) Brown(x string) string {
	if !c.isEnabled {
		return x
	}

	return brown + x + reset
}

func (c *Color) Yellow(x string) string {
	if !c.isEnabled {
		return x
	}

	return yellow + x + reset
}

func (c *Color) Blue(x string) string {
	if !c.isEnabled {
		return x
	}

	return blue + x + reset
}

func (c *Color) LightBlue(x string) string {
	if !c.isEnabled {
		return x
	}

	return lightBlue + x + reset
}

func (c *Color) Purple(x string) string {
	if !c.isEnabled {
		return x
	}

	return purple + x + reset
}

func (c *Color) LightPurple(x string) string {
	if !c.isEnabled {
		return x
	}

	return lightPurple + x + reset
}

func (c *Color) Cyan(x string) string {
	if !c.isEnabled {
		return x
	}

	return cyan + x + reset
}

func (c *Color) LightCyan(x string) string {
	if !c.isEnabled {
		return x
	}

	return lightCyan + x + reset
}

func (c *Color) LightGray(x string) string {
	if !c.isEnabled {
		return x
	}

	if c.isTrueColor {
		return "\033[0;90m" + x + "\033[0m"
	}

	return lightGray + x + reset
}

func (c *Color) White(x string) string {
	if !c.isEnabled {
		return x
	}

	return white + x + reset
}
