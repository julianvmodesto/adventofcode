package main

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		fish []int
		days int
		want int
	}{
		{
			fish: []int{3, 4, 3, 1, 2},
			days: 18,
			want: 26,
		},
		{
			fish: []int{3, 4, 3, 1, 2},
			days: 80,
			want: 5934,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("partOne(%d)", test.days), func(t *testing.T) {
			got := partOne(test.fish, test.days)
			if got != test.want {
				t.Fatalf("partOne(%d) = %d, want %d", test.days, got, test.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		fish []int
		days int
		want int
	}{
		{
			fish: []int{3, 4, 3, 1, 2},
			days: 18,
			want: 26,
		},
		{
			fish: []int{3, 4, 3, 1, 2},
			days: 80,
			want: 5934,
		},
		{
			fish: []int{3, 4, 3, 1, 2},
			days: 256,
			want: 26984457539,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("partTwo(%d)", test.days), func(t *testing.T) {
			got := partTwo(test.fish, test.days)
			if got != test.want {
				t.Fatalf("partTwo(%d) = %d, want %d", test.days, got, test.want)
			}
		})
	}
}
