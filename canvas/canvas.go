package canvas

import (
	"bytes"
	"fmt"

	"github.com/ashvarts/raytracer/artcolor"
)

type Coordinates struct {
	X, Y int
}

type Pixel map[Coordinates]artcolor.Color

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
			canvas.Pixels[Coordinates{x, y}] = artcolor.NewColor(0, 0, 0)
		}
	}
	return canvas
}

func (c Canvas) WritePixel(x int, y int, col artcolor.Color) {
	cordonates := Coordinates{x, y}
	c.Pixels[cordonates] = col
}

func (c Canvas) PixelAt(x, y int) artcolor.Color {
	cordonates := Coordinates{x, y}
	return c.Pixels[cordonates]
}

func (c Canvas) writePPMHeader(buf *bytes.Buffer) {
	buf.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height))
}

func (c Canvas) writePPMBody(buf *bytes.Buffer) {
	for row := 0; row < c.Height; row++ {
		for col := 0; col < c.Width; col++ {
			buf.WriteString(c.PixelAt(col, row).String())
			// if last pixel, exit
			if row == c.Height-1 && col == c.Width-1 {
				break
				// if not first, but a multiple of 4, or last in col then start new line
			} else if col != 0 && col%4 == 0 || col == c.Width-1 {
				buf.WriteString("\n")
			} else {
				buf.WriteString(" ")
			}
		}
	}
}

func (c Canvas) writePPMFooter(buf *bytes.Buffer) {
	buf.WriteString("\n")
}

func (c Canvas) String() string {
	buf := &bytes.Buffer{}
	c.writePPMHeader(buf)
	c.writePPMBody(buf)
	c.writePPMFooter(buf)
	return buf.String()
}
