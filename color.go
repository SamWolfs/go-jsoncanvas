package jsoncanvas

import (
	"fmt"
	"regexp"
)

type Color string

const (
	ColorRed    Color = "1"
	ColorOrange Color = "2"
	ColorYellow Color = "3"
	ColorGreen  Color = "4"
	ColorCyan   Color = "5"
	ColorPurple Color = "6"
)

func (c *Color) Validate() error {
	// check if color is a hex color or a preset color
	if c == nil || *c == "" {
		return nil
	}

	if (*c)[0] == '#' {
		hexCode := string(*c)[1:]
		return validateHexColor(hexCode)
	} else if *c == ColorRed || *c == ColorOrange || *c == ColorYellow || *c == ColorGreen || *c == ColorCyan || *c == ColorPurple {
		// preset color
		return nil
	} else {
		return fmt.Errorf("invalid color: %s", *c)
	}
}

func validateHexColor(code string) error {
	r, _ := regexp.Compile("^#?([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$")

	if r.MatchString(code) {
		return nil
	} else {
		return fmt.Errorf("invalid hex color: %s", code)
	}
}
