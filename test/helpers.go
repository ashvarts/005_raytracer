package test

import (
	"math"

	"github.com/ashvarts/raytracer/color"
)

const epsilon = 0.00001

func AproxEquall(a, b interface{}) bool {
	switch a := a.(type) {
	case float64:
		return aproxEqual(a, b.(float64))
	case color.Color:
		return aproxEqual(a.Red, b.(color.Color).Red) && aproxEqual(a.Green, b.(color.Color).Green) && aproxEqual(a.Blue, b.(color.Color).Blue)
	default:
		return false
	}
}

func aproxEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
