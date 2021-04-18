package main

import (
	"fmt"
	"testing"
)

func TestCalculateFloor(t *testing.T) {
	// ) causes him to enter the basement at character position 1.
	// ()()) causes him to enter the basement at character position 5.
	cases := []struct {
		symbol   string
		expected int
	}{
		{symbol: ")", expected: 1},
		{symbol: "()())", expected: 5},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("string_%s", c.symbol), func(t *testing.T) {
			_, actual := CalculateFloorAndPosition(c.symbol)
			if actual != c.expected {
				t.Errorf("expected %v, got %v", c.expected, actual)
			}
		})
	}
}
