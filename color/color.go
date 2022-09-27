package color

import (
	"fmt"
	"math"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c Color) Add(co Color) Color {
	return Color{
		c.Red + co.Red,
		c.Green + co.Green,
		c.Blue + co.Blue,
	}
}

func (c Color) Sub(co Color) Color {
	return Color{
		c.Red - co.Red,
		c.Green - co.Green,
		c.Blue - co.Blue,
	}
}

func (c Color) ScalarMulti(n float64) Color {
	return Color{
		c.Red * n,
		c.Green * n,
		c.Blue * n,
	}
}

func (c Color) Multi(co Color) Color {
	return Color{
		c.Red * co.Red,
		c.Green * co.Green,
		c.Blue * co.Blue,
	}
}

func (c Color) String() string {
	// normalize the color values between 0-255
	r := normalized(c.Red)
	g := normalized(c.Green)
	b := normalized(c.Blue)
	return fmt.Sprintf("%d %d %d", r, g, b)
}

func normalized(c float64) int {
	n := math.Round(c * 255)
	if n < 0 {
		return 0
	} else if n > 255 {
		return 255
	}
	return int(n)
}
