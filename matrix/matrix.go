package matrix

type Dim struct {
	R, C int
}

type Matrix map[Dim]float64

func NewMatrix() Matrix {
	return make(Matrix)
}
