package color

import "testing"

func TestColor(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)
	if c.red != -0.5 || c.green != 0.4 || c.blue != 1.7 {
		t.Error("Color not made correctly")
	}
}
