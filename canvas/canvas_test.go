package canvas_test

import (
	"testing"

	"github.com/ashvarts/raytracer/canvas"
	"github.com/ashvarts/raytracer/color"
)

func TestNewCanvas(t *testing.T) {
	c := canvas.NewCanvas(10, 20)
	if c.Width != 10 || c.Height != 20 {
		t.Errorf("expected width=10, got=%d; height=20,got=%d", c.Width, c.Height)
	}

	if c.Pixels == nil {
		t.Fatal("pixels should not be nil")
	}

	for _, col := range c.Pixels {
		if col != color.NewColor(0, 0, 0) {
			t.Errorf("pixel should be color(0,0,0)")
		}
	}
}
