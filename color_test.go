package jsoncanvas

import (
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidColor(t *testing.T) {
	validColors := [8]Color{
		ColorRed, ColorOrange, ColorYellow,
		ColorGreen, ColorCyan, ColorPurple,
		"#FfFfFa", "#000",
	}

	for _, color := range validColors {
		err := (&color).Validate()
		require.NoError(t, err)
	}
}

func TestInvalidColor(t *testing.T) {
	invalidColors := [2]Color{"green", "7"}

	for _, color := range invalidColors {
		err := (&color).Validate()
		require.Error(t, err)
	}
}

func TestInvalidHexColor(t *testing.T) {
	invalidColors := [6]Color{"#XYZ", "#XXYYZZ", "#A", "#AB", "#AABB", "#AABBC"}

	for _, color := range invalidColors {
		err := (&color).Validate()
		require.Error(t, err)
	}
}
