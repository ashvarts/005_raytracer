package main

import (
	"fmt"
	"os"

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
var e = environment{tuple.Vector(0, -0.1, 0), tuple.Vector(-0.01, 0, 0)}

func tick(env environment, proj projectile) projectile {
	pos := proj.position.Add(proj.velocity)
	vel := proj.velocity.Add(env.gravity.Add(env.wind))
	return projectile{pos, vel}
}

func main() {
	c := canvas.NewCanvas(20, 10)
	for {
		p = tick(e, p)
		if p.position.Y < 0.0 {
			break
		}
		fmt.Printf("(%f,%f)\n", p.position.X, p.position.Y)
	}

	f, err := os.Create("test.ppm")
	check(err)
	defer f.Close()

	f.WriteString(c.String())
	f.Sync()

}
