package main

import "testing"

func TestOperate(t *testing.T) {
	var tests = []struct {
		program   []int
		operation [4]int
		expected  []int
	}{
		{
			program:   []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			operation: [4]int{1, 9, 10, 3},
			expected:  []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			program:   []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
			operation: [4]int{2, 3, 11, 0},
			expected:  []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			program:   []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
			operation: [4]int{99, 30, 40, 50},
			expected:  []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
	}

	for i, test := range tests {
		var intcode = NewIntcode(test.program)
		intcode.Operate(test.operation[0], test.operation[1], test.operation[2], test.operation[3])
		if !testSliceEqual(intcode.program, test.expected) {
			t.Fatalf("test %d failed: expected programs to be equal, but got:\n%v\n%v\n", i, intcode.program, test.program)
		}
	}
}

func TestProcess(t *testing.T) {
	var tests = []struct {
		program  []int
		expected []int
	}{
		{
			program:  []int{},
			expected: []int{},
		},
		{
			program:  []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			expected: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
		{
			program:  []int{1, 0, 0, 0, 99},
			expected: []int{2, 0, 0, 0, 99},
		},
		{
			program:  []int{2, 3, 0, 3, 99},
			expected: []int{2, 3, 0, 6, 99},
		},
		{
			program:  []int{2, 4, 4, 5, 99, 0},
			expected: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			program:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for i, test := range tests {
		var intcode = NewIntcode(test.program)
		intcode.Process()
		if !testSliceEqual(intcode.program, test.expected) {
			t.Fatalf("test %d failed: expected programs to be equal, but got:\n%v\n%v\n", i, intcode.program, test.program)
		}
	}
}

func testSliceEqual(a, b []int) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
