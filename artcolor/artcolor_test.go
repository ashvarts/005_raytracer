package artcolor_test

import (
	"testing"

	"github.com/ashvarts/raytracer/artcolor"
	"github.com/ashvarts/raytracer/test"
)

func TestColor(t *testing.T) {
	c := artcolor.NewColor(-0.5, 0.4, 1.7)
	if c.Red != -0.5 || c.Green != 0.4 || c.Blue != 1.7 {
		t.Error("Color not made correctly")
	}
}

func TestColorAdd(t *testing.T) {
	c1 := artcolor.NewColor(0.9, 0.6, 0.75)
	c2 := artcolor.NewColor(0.7, 0.1, 0.25)
	want := artcolor.NewColor(1.6, 0.7, 1.0)
	if got := c1.Add(c2); got != want {
		t.Errorf("wanted:%v, got:%v", want, got)
	}
}

func TestColorSub(t *testing.T) {
	c1 := artcolor.NewColor(0.9, 0.6, 0.75)
	c2 := artcolor.NewColor(0.7, 0.1, 0.25)
	want := artcolor.NewColor(0.2, 0.5, 0.5)
	if got := c1.Sub(c2); !test.AproxEquall(got.Blue, want.Blue) || !test.AproxEquall(got.Green, want.Green) || !test.AproxEquall(got.Red, want.Red) {
		t.Errorf("wanted:%v, got:%v", want, got)
	}
}

func TestColorScalarMulti(t *testing.T) {
	c1 := artcolor.NewColor(0.2, 0.3, 0.4)
	scalar := 2.0
	want := artcolor.NewColor(0.4, 0.6, 0.8)
	if got := c1.ScalarMulti(scalar); got != want {
		t.Errorf("wanted:%v, got:%v", want, got)
	}
}

func TestColorMulti(t *testing.T) {
	c1 := artcolor.NewColor(1.0, 0.2, 0.4)
	c2 := artcolor.NewColor(0.9, 1, 0.1)
	want := artcolor.NewColor(0.9, 0.2, 0.04)
	if got := c1.Multi(c2); !test.AproxEquall(got, want) {
		t.Errorf("wanted:%v, got:%v", want, got)
	}
}
