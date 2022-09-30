package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ashvarts/raytracer/artcolor"
	"github.com/ashvarts/raytracer/canvas"
	"github.com/ashvarts/raytracer/tuple"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type projectile struct {
	position tuple.Tuple
	velocity tuple.Tuple
}

type environment struct {
	gravity tuple.Tuple
	wind    tuple.Tuple
}

var p = projectile{tuple.Point(0, 1, 0), tuple.Vector(1, 1, 0)}
var e = environment{tuple.Vector(0, 0, 0), tuple.Vector(0, 0, 0)}

func tick(env environment, proj projectile) projectile {
	pos := proj.position.Add(proj.velocity)
	vel := proj.velocity.Add(env.gravity.Add(env.wind))
	return projectile{pos, vel}
}

var white = artcolor.NewColor(0.8, 0.8, 0.8)

func main() {
	c := canvas.NewCanvas(20, 20)
	// for i := 0; i < 500; i++ {
	// 	p = tick(e, p)
	// 	// co := coordinates(p)
	// 	// c.WritePixel(co.X, c.Height-co.Y, white)
	// 	c.WritePixel(c.Height-i, i, white)
	// 	if p.position.Y < 0.0 {
	// 		break
	// 	}

	// }

	for y := 0; y < c.Height; y++ {
		if y%2 == 0 {
			for x := 0; x < c.Width; x++ {
				c.WritePixel(y, x, white)
			}
		}

	}

	f, err := os.Create("test.ppm")
	check(err)
	defer f.Close()

	fmt.Println(c.String())
	f.WriteString(c.String())
	f.Sync()

}

func coordinates(p projectile) canvas.Coordinates {
	x := int(math.Round(p.position.X))
	y := int(math.Round(p.position.Y))
	return canvas.Coordinates{X: x, Y: y}
}
