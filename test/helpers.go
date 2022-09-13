package test

import (
	"math"
)

const epsilon = 0.00001

func AproxEquall(a, b interface{}) bool {
	switch a := a.(type) {
	case float64:
		return math.Abs(a-b.(float64)) < epsilon
	default:
		return false
	}
}
