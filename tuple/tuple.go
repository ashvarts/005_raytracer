package tuple

import "math"

type Tuple struct {
	X, Y, Z, W float64
}

func (t Tuple) IsPoint() bool {
	if int(t.W) == 1 {
		return true
	} else {
		return false
	}
}

func (t Tuple) IsVector() bool {
	if int(t.W) == 0 {
		return true
	} else {
		return false
	}
}

func (t Tuple) Add(tt Tuple) Tuple {
	return Tuple{
		t.X + tt.X,
		t.Y + tt.Y,
		t.Z + tt.Z,
		t.W + tt.W,
	}
}

func (t Tuple) Sub(tt Tuple) Tuple {
	return Tuple{
		t.X - tt.X,
		t.Y - tt.Y,
		t.Z - tt.Z,
		t.W - tt.W,
	}
}

func (t Tuple) Negate() Tuple {
	return Tuple{
		-1 * t.X,
		-1 * t.Y,
		-1 * t.Z,
		-1 * t.W,
	}
}

// Multiply scales a vector by a scalar
func (t Tuple) Multiply(s float64) Tuple {

	return Tuple{
		t.X * s,
		t.Y * s,
		t.Z * s,
		t.W * s,
	}
}

// Magnitude return the magnitude of a vector
func (t Tuple) Magnitude() float64 {
	sumOfSquares := math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(t.W, 2)
	return math.Sqrt(sumOfSquares)
}

// Normalize returns the normalized for of a vector
func (t Tuple) Normalize() Tuple {
	magnitude := t.Magnitude()
	return Tuple{
		t.X / magnitude,
		t.Y / magnitude,
		t.Z / magnitude,
		t.W / magnitude,
	}
}

// Dot takes a vector and returns the dot product.
func (t Tuple) Dot(v Tuple) float64 {
	return t.X*v.X + t.Y*v.Y + t.Z*v.Z + t.W + v.W
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}
