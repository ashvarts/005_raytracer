package canvas_test

import (
	"bytes"
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

func TestWritingPixelToCanvas(t *testing.T) {
	c := canvas.NewCanvas(10, 20)
	red := color.NewColor(1, 0, 0)

	c.WritePixel(2, 3, red)
	gott := c.PixelAt(2, 3)
	_ = gott
	if got := c.PixelAt(2, 3); got != red {
		t.Errorf("Expected pixel to be %v, got:%v", red, got)
	}
}

func TestCanvasToPPM(t *testing.T) {
	buf := bytes.Buffer{}
	c := canvas.NewCanvas(5, 3)
	err := c.WritePPMHeader(&buf)
	if err != nil {
		t.Fatal(err)
	}
	got := buf.String()
	want := `P3
5 3
255`
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestCanvasBodyToPPM(t *testing.T) {

	t.Run("1,1", func(t *testing.T) {
		buf := bytes.Buffer{}
		c := canvas.NewCanvas(1, 1)
		err := c.WritePPMBody(&buf)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `0 0 0`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("5,3", func(t *testing.T) {
		buf := bytes.Buffer{}
		c := canvas.NewCanvas(5, 3)
		err := c.WritePPMBody(&buf)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("5,3 with pixels", func(t *testing.T) {
		buf := bytes.Buffer{}
		c := canvas.NewCanvas(5, 3)
		c.WritePixel(0, 0, color.NewColor(1.5, 0, 0))
		c.WritePixel(2, 1, color.NewColor(0, 0.5, 0))
		c.WritePixel(4, 2, color.NewColor(-0.5, 0, 1))
		err := c.WritePPMBody(&buf)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("no line over 70 chars long", func(t *testing.T) {
		buf := bytes.Buffer{}
		c := canvas.NewCanvas(10, 2)

		// set every pixel to almost white
		for pixel := range c.Pixels {
			c.WritePixel(pixel.X, pixel.Y, color.NewColor(1, 0.8, 0.6))
		}
		err := c.WritePPMBody(&buf)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
