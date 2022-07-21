package main

import (
	"testing"
)

func Test_tuple_isPoint(t *testing.T) {
	type fields struct {
		x float64
		y float64
		z float64
		w float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"isPoint", fields{4.3, -4.2, 3.1, 1.0}, true},
		{"isPoint", fields{4.3, -4.2, 3.1, 0.0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tuple{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
				w: tt.fields.w,
			}
			if got := tr.isPoint(); got != tt.want {
				t.Errorf("tuple.isPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tuple_isVector(t *testing.T) {
	type fields struct {
		x float64
		y float64
		z float64
		w float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"isVector", fields{4.3, -4.2, 3.1, 1.0}, false},
		{"isVector", fields{4.3, -4.2, 3.1, 0.0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tuple{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
				w: tt.fields.w,
			}
			if got := tr.isVector(); got != tt.want {
				t.Errorf("tuple.isVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
