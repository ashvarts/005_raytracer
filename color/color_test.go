package color_test

import (
	"testing"

	"github.com/ashvarts/raytracer/color"
	"github.com/ashvarts/raytracer/test"
)

func TestColor(t *testing.T) {
	c := color.NewColor(-0.5, 0.4, 1.7)
	if c.Red != -0.5 || c.Green != 0.4 || c.Blue != 1.7 {
		t.Error("Color not made correctly")
	}
}

func TestColorAdd(t *testing.T) {
	c1 := color.NewColor(0.9, 0.6, 0.75)
	c2 := color.NewColor(0.7, 0.1, 0.25)
	want := color.NewColor(1.6, 0.7, 1.0)
	if got := c1.Add(c2); got != want {
		t.Errorf("wanted:%v, got:%v", want, got)
	}
}

func TestColorSub(t *testing.T) {
	c1 := color.NewColor(0.9, 0.6, 0.75)
	c2 := color.NewColor(0.7, 0.1, 0.25)
	want := color.NewColor(0.2, 0.5, 0.5)
	if got := c1.Sub(c2); !test.AproxEquall(got.Blue, want.Blue) || !test.AproxEquall(got.Green, want.Green) || !test.AproxEquall(got.Red, want.Red) {
		t.Errorf("wanted:%v, got:%v", want, got)
	}
}
