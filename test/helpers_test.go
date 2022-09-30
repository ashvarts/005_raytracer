package test

import (
	"testing"

	"github.com/ashvarts/raytracer/artcolor"
)

func TestAssertAlmostEquall(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"diff less than epsilon", args{1.000000001, 1.0000000009}, true},
		{"diff less than epsilon", args{-1.000000001, -1.0000000009}, true},
		{"diff more than epsilon", args{1.000000001, 2.0000000009}, false},
		{"diff more than epsilon", args{1.000000001, -2.0000000009}, false},
		{"diff more than epsilon", args{1.100000001, 1.2000000009}, false},
		{"diff artcolor.Color", args{artcolor.NewColor(1, 1, 1), artcolor.NewColor(1, 1, 1)}, true},
		{"diff artcolor.Color", args{artcolor.NewColor(1, 0, 1), artcolor.NewColor(1, 1, 1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AproxEquall(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("aproxEquall() = %v, want %v", got, tt.want)
			}
		})
	}
}
