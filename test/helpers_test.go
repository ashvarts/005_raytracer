package test

import "testing"

func TestAssertAlmostEquall(t *testing.T) {
	type args struct {
		a float64
		b float64
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aproxEquall(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("aproxEquall() = %v, want %v", got, tt.want)
			}
		})
	}
}
