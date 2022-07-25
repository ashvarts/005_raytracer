package tuple

type Tuple struct {
	x, y, z, w float64
}

func (t Tuple) isPoint() bool {
	if int(t.w) == 1 {
		return true
	} else {
		return false
	}
}

func (t Tuple) isVector() bool {
	if int(t.w) == 0 {
		return true
	} else {
		return false
	}
}

func point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}
