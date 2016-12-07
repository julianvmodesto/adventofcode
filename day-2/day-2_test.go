package main

import (
	"strings"
	"testing"
)

func TestPressButtons(t *testing.T) {
	in := `
		ULL
		RRDDD
		LURDL
		UUUUD
	`
	buttons := pressButtons(in)
	if strings.EqualFold(buttons, "1985") {
		t.Errorf("Expected buttons 1985, but it was %s instead.", buttons)
	}
}
