package canvas

import (
	"bytes"
	"fmt"
	"io"

	"github.com/ashvarts/raytracer/color"
)

type Coordinates struct {
	X, Y int
}

type Pixel map[Coordinates]color.Color

type Canvas struct {
	Width, Height int
	Pixels        Pixel
}

func NewCanvas(w, h int) Canvas {
	canvas := Canvas{
		Width:  w,
		Height: h,
		Pixels: make(Pixel, w*h),
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			canvas.Pixels[Coordinates{x, y}] = color.NewColor(0, 0, 0)
		}
	}
	return canvas
}

func (c Canvas) WritePixel(x int, y int, col color.Color) {
	cordonates := Coordinates{x, y}
	c.Pixels[cordonates] = col
}

func (c Canvas) PixelAt(x, y int) color.Color {
	cordonates := Coordinates{x, y}
	return c.Pixels[cordonates]
}

func (c Canvas) WritePPMHeader(w io.Writer) error {
	_, err := fmt.Fprintf(w, "P3\n%d %d\n255", c.Width, c.Height)
	if err != nil {
		return err
	}
	return nil
}

func (c Canvas) WritePPMBody(buf *bytes.Buffer) error {
	buf.WriteString(c.PixelAt(0, 0).String())
	return nil
}
