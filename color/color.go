package color

type Color struct {
	red   float64
	green float64
	blue  float64
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}
