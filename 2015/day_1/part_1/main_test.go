package main

import (
	"fmt"
	"testing"
)

func TestCalculateFloor(t *testing.T) {
	// (()) and ()() both result in floor 0.
	// ((( and (()(()( both result in floor 3.
	// ))((((( also results in floor 3.
	// ()) and ))( both result in floor -1 (the first basement level).
	// ))) and )())()) both result in floor -3.
	cases := []struct {
		symbol   string
		expected int
	}{
		{symbol: "(())", expected: 0},
		{symbol: "()()", expected: 0},
		{symbol: "(()(()(", expected: 3},
		{symbol: "(((", expected: 3},
		{symbol: "))(((((", expected: 3},
		{symbol: "())", expected: -1},
		{symbol: "))(", expected: -1},
		{symbol: ")))", expected: -3},
		{symbol: ")())())", expected: -3},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("string_%s", c.symbol), func(t *testing.T) {
			actual := calculateFloor(c.symbol)
			if actual != c.expected {
				t.Errorf("expected %v, got %v", c.expected, actual)
			}
		})
	}
}
