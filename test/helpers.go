package test

import (
	"math"

	"github.com/ashvarts/raytracer/artcolor"
)

const epsilon = 0.00001

func AproxEquall(a, b interface{}) bool {
	switch a := a.(type) {
	case float64:
		return aproxEqual(a, b.(float64))
	case artcolor.Color:
		return aproxEqual(a.Red, b.(artcolor.Color).Red) && aproxEqual(a.Green, b.(artcolor.Color).Green) && aproxEqual(a.Blue, b.(artcolor.Color).Blue)
	default:
		return false
	}
}

func aproxEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
