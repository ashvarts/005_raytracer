package main

type tuple struct {
	x, y, z, w float64
}

func (t tuple) isPoint() bool {
	if int(t.w) == 1 {
		return true
	} else {
		return false
	}
}

func (t tuple) isVector() bool {
	if int(t.w) == 0 {
		return true
	} else {
		return false
	}
}
