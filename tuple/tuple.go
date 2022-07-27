package tuple

type Tuple struct {
	x, y, z, w float64
}

func (t Tuple) IsPoint() bool {
	if int(t.w) == 1 {
		return true
	} else {
		return false
	}
}

func (t Tuple) IsVector() bool {
	if int(t.w) == 0 {
		return true
	} else {
		return false
	}
}

func (t Tuple) Add(tt Tuple) Tuple {
	return Tuple{
		t.x + tt.x,
		t.y + tt.y,
		t.z + tt.z,
		t.w + tt.w,
	}
}

func (t Tuple) Sub(tt Tuple) Tuple {
	return Tuple{
		t.x - tt.x,
		t.y - tt.y,
		t.z - tt.z,
		t.w + tt.w,
	}
}

func (t Tuple) Negate() Tuple {
	return Tuple{
		-1 * t.x,
		-1 * t.y,
		-1 * t.z,
		-1 * t.w,
	}
}

// Multiply scales a vector by a scalar
func (t Tuple) Multiply(s float64) Tuple {

	return Tuple{
		t.x * s,
		t.y * s,
		t.z * s,
		t.w * s,
	}
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}
