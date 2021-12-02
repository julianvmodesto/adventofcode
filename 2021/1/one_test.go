package main

import "testing"

func TestSliding(t *testing.T) {
	in := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	want := 5
	if got := largerMeasurementsSliding(in); got != want {
		t.Fatalf("largerMeasurementsSliding() = %d, want %d", got, want)
	}
}
