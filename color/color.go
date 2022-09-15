package color

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
