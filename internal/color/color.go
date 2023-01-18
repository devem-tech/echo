package color

import "github.com/devemio/mockio/pkg/color"

type Color struct {
	enabled bool
}

func New(enabled bool) *Color {
	return &Color{
		enabled: enabled,
	}
}

func (c *Color) Black(value string) string {
	if !c.enabled {
		return value
	}

	return color.Black(value)
}

func (c *Color) DarkGray(value string) string {
	if !c.enabled {
		return value
	}

	return color.DarkGray(value)
}

func (c *Color) Red(value string) string {
	if !c.enabled {
		return value
	}

	return color.Red(value)
}

func (c *Color) LightRed(value string) string {
	if !c.enabled {
		return value
	}

	return color.LightRed(value)
}

func (c *Color) Green(value string) string {
	if !c.enabled {
		return value
	}

	return color.Green(value)
}

func (c *Color) LightGreen(value string) string {
	if !c.enabled {
		return value
	}

	return color.LightGreen(value)
}

func (c *Color) Brown(value string) string {
	if !c.enabled {
		return value
	}

	return color.Brown(value)
}

func (c *Color) Yellow(value string) string {
	if !c.enabled {
		return value
	}

	return color.Yellow(value)
}

func (c *Color) Blue(value string) string {
	if !c.enabled {
		return value
	}

	return color.Blue(value)
}

func (c *Color) LightBlue(value string) string {
	if !c.enabled {
		return value
	}

	return color.LightBlue(value)
}

func (c *Color) Purple(value string) string {
	if !c.enabled {
		return value
	}

	return color.Purple(value)
}

func (c *Color) LightPurple(value string) string {
	if !c.enabled {
		return value
	}

	return color.LightPurple(value)
}

func (c *Color) Cyan(value string) string {
	if !c.enabled {
		return value
	}

	return color.Cyan(value)
}

func (c *Color) LightCyan(value string) string {
	if !c.enabled {
		return value
	}

	return color.LightCyan(value)
}

func (c *Color) LightGray(value string) string {
	if !c.enabled {
		return value
	}

	return color.LightGray(value)
}

func (c *Color) White(value string) string {
	if !c.enabled {
		return value
	}

	return color.White(value)
}
