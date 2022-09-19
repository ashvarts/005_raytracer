package canvas

import "github.com/ashvarts/raytracer/color"

type Coordinates struct {
	X, Y int
}

type Pixel map[Coordinates]color.Color

type Canvas struct {
	Width, Height int
	Pixels        []Pixel
}

func NewCanvas(w, h int) Canvas {
	canvas := Canvas{
		Width:  w,
		Height: h,
		Pixels: make([]Pixel, 0),
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			p := Pixel{
				Coordinates{x, y}: color.NewColor(0, 0, 0),
			}
			canvas.Pixels = append(canvas.Pixels, p)
		}
	}
	return canvas
}
