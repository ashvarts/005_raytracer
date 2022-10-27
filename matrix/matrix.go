package matrix

type Dim struct {
	R, C int
}

type Matrix [][]float64

func NewMatrix(r, c int) Matrix {
	return make(Matrix, r*c)
}
