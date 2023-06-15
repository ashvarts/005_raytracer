package matrix

type Matrix [][]float64

// A matrix is a n-by-n shape and contains a number in each spot.
// Two matrix are equal when:
// 1. They have the same dimensions.
// 2. The corresponding numbers in each matrix are aproximately equal.
func Equal(a, b Matrix) bool {
	return false
}
