package color

type Color struct {
	red   float64
	green float64
	blue  float64
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c Color) Add(co Color) Color {
	return Color{
		c.red + co.red,
		c.green + co.green,
		c.blue + co.blue,
	}
}
