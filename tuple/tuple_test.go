package tuple

import (
	"reflect"
	"testing"
)

func Test_Tuple_isPoint(t *testing.T) {
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
			tr := Tuple{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
				w: tt.fields.w,
			}
			if got := tr.IsPoint(); got != tt.want {
				t.Errorf("Tuple.isPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Tuple_isVector(t *testing.T) {
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
			tr := Tuple{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
				w: tt.fields.w,
			}
			if got := tr.IsVector(); got != tt.want {
				t.Errorf("Tuple.isVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_point(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want Tuple
	}{
		{"new point", args{4, -4, 3}, Tuple{4, -4, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Point(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("point() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vector(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want Tuple
	}{
		{"new vector", args{4, -4, 3}, Tuple{4, -4, 3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vector(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTupleAdd(t *testing.T) {
	t1 := Point(3, -2, 5)
	t2 := Vector(-2, 3, 1)

	sum := t1.Add(t2)
	want := Tuple{1, 1, 6, 1}
	if sum != want {
		t.Errorf("expected: %v, got: %v", sum, want)
	}
}

func TestTupleSub(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	sum := p1.Sub(p2)
	want := Vector(-2, -4, -6)
	if sum != want {
		t.Errorf("expected: %v, got: %v", sum, want)
	}
}

func TestTuple_Sub(t *testing.T) {
	tests := []struct {
		name string
		t1   Tuple
		t2   Tuple
		want Tuple
	}{
		{"Subtract two point", Point(3, 2, 1), Point(5, 6, 7), Vector(-2, -4, -6)},
		{"Subtract a vector from point", Point(3, 2, 1), Vector(5, 6, 7), Point(-2, -4, -6)},
		{"Subtract two vectors", Vector(3, 2, 1), Vector(5, 6, 7), Vector(-2, -4, -6)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t1.Sub(tt.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
