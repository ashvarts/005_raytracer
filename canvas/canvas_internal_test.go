package canvas

import (
	"bytes"
	"testing"

	"github.com/ashvarts/raytracer/artcolor"
	"github.com/stretchr/testify/assert"
)

func TestCanvasToPPM(t *testing.T) {
	buf := &bytes.Buffer{}
	c := NewCanvas(5, 3)
	c.writePPMHeader(buf)
	got := buf.String()
	want := `P3
5 3
255
`
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func TestCanvasBodyToPPM(t *testing.T) {

	t.Run("1,1", func(t *testing.T) {
		buf := bytes.Buffer{}
		c := NewCanvas(1, 1)
		c.writePPMBody(&buf)
		got := buf.String()
		want := `0 0 0`
		assert.Equal(t, want, got)
	})
	t.Run("5,3", func(t *testing.T) {
		buf := bytes.Buffer{}
		c := NewCanvas(5, 3)
		c.writePPMBody(&buf)
		got := buf.String()
		want := `0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0`
		assert.Equal(t, want, got)
	})
	t.Run("5,3 with pixels", func(t *testing.T) {
		buf := bytes.Buffer{}
		c := NewCanvas(5, 3)
		c.WritePixel(0, 0, artcolor.NewColor(1.5, 0, 0))
		c.WritePixel(2, 1, artcolor.NewColor(0, 0.5, 0))
		c.WritePixel(4, 2, artcolor.NewColor(-0.5, 0, 1))
		c.writePPMBody(&buf)
		got := buf.String()
		want := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255`
		assert.Equal(t, want, got)
	})

	t.Run("no line over 70 chars long", func(t *testing.T) {
		buf := bytes.Buffer{}
		c := NewCanvas(10, 2)

		// set every pixel to almost white
		for pixel := range c.Pixels {
			c.WritePixel(pixel.X, pixel.Y, artcolor.NewColor(1, 0.8, 0.6))
		}
		c.writePPMBody(&buf)
		got := buf.String()
		want := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153
255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153
255 204 153`
		assert.Equal(t, want, got)

	})
}
