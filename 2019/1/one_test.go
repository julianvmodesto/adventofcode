package main

import "testing"

func TestGetModuleFuel(t *testing.T) {
	var tests = []struct {
		mass int64
		fuel int64
	}{
		{
			mass: 12,
			fuel: 2,
		},
		{
			mass: 14,
			fuel: 2,
		},
		{
			mass: 1969,
			fuel: 654,
		},
		{
			mass: 100756,
			fuel: 33583,
		},
	}

	for i, test := range tests {
		fuel := getModuleFuel(test.mass)
		if fuel != test.fuel {
			t.Errorf("test %d failed: expected %d, but got %d", i, test.fuel, fuel)
		}
	}
}
