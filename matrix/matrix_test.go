package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatrix(t *testing.T) {
	assert := assert.New(t)
	m := Matrix{
		[]float64{1, 2, 3, 4},
		[]float64{5.5, 6.5, 7.5, 8.5},
		[]float64{9, 10, 11, 12},
		[]float64{13.5, 14.5, 15.5, 16.5},
	}
	assert.Equal(m[0][0], 1.0)
	assert.Equal(m[0][3], 4.0)
	assert.Equal(m[1][0], 5.5)
	assert.Equal(m[1][2], 7.5)
	assert.Equal(m[2][2], 11.0)
	assert.Equal(m[3][0], 13.5)
	assert.Equal(m[3][2], 15.5)
}
