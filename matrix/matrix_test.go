package matrix

import (
	"fmt"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m := Matrix{
		[]float64{24, 23},
		[]float64{4, 5},
	}
	fmt.Println(m)
}
