package matrix

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	tests := []struct {
		name string
		want Matrix
	}{{
		name: "Create New",
		want: make(Matrix),
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
