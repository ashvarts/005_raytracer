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

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}
