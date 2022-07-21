package main

import "testing"

func TestTupleIsPoint(t *testing.T) {
	tup := tuple{4.3, -4.2, 3.1, 1.0}
	if tup.isPoint() != true {
		t.Error("tuple is point should be true")
	}

	if tup.isVector() != false {
		t.Error("tuple is vector should be false")
	}
}
