package test

import "math"

const epsilon = 0.00001

func AssertAlmostEquall(a float64, b float64) bool {

	return math.Abs(a-b) < epsilon

}
