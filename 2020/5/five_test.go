package main

import (
	"strings"
	"testing"
)

func TestGetSeatID(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{

		{
			name:     "BFFFBBFRRR: row 70, column 7, seat ID 567",
			line:     "BFFFBBFRRR",
			expected: 567,
		},
		{
			name:     "FFFBBBFRRR: row 14, column 7, seat ID 119",
			line:     "FFFBBBFRRR",
			expected: 119,
		},
		{
			name:     "BBFFBBFRLL: row 102, column 4, seat ID 820",
			line:     "BBFFBBFRLL",
			expected: 820,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			letters := strings.Split(test.line, "")
			if got := GetSeatID(letters); got != test.expected {
				t.Fatalf("getSeatID(%q) = %d, want %d\n", test.line, got, test.expected)
			}
		})
	}
}
