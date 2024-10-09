package color

import (
	"os"

	"github.com/devem-tech/echo/pkg/color"
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

	return color.Black(x)
}

func (c *Color) DarkGray(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.DarkGray(x)
}

func (c *Color) Red(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.Red(x)
}

func (c *Color) LightRed(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.LightRed(x)
}

func (c *Color) Green(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.Green(x)
}

func (c *Color) LightGreen(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.LightGreen(x)
}

func (c *Color) Brown(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.Brown(x)
}

func (c *Color) Yellow(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.Yellow(x)
}

func (c *Color) Blue(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.Blue(x)
}

func (c *Color) LightBlue(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.LightBlue(x)
}

func (c *Color) Purple(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.Purple(x)
}

func (c *Color) LightPurple(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.LightPurple(x)
}

func (c *Color) Cyan(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.Cyan(x)
}

func (c *Color) LightCyan(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.LightCyan(x)
}

func (c *Color) LightGray(x string) string {
	if !c.isEnabled {
		return x
	}

	if c.isTrueColor {
		return "\033[0;90m" + x + "\033[0m"
	}

	return color.LightGray(x)
}

func (c *Color) White(x string) string {
	if !c.isEnabled {
		return x
	}

	return color.White(x)
}
