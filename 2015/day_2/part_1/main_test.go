package main

import (
	"fmt"
	"testing"
)

func TestBox_Area(t *testing.T) {
	box := NewBox(2, 3, 4)
	actual := box.Area()
	if actual != 52 {
		t.Fatalf("expected 0, got: %v", actual)
	}
}

func TestBox_SmallArea(t *testing.T) {
	box := NewBox(2, 3, 4)
	actual := box.SmallArea()
	if actual != 6 {
		t.Fatalf("expected 0, got: %v", actual)
	}
}

func TestAllArea(t *testing.T) {
	// A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper
	// plus 6 square feet of slack, for a total of 58 square feet.
	// A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper
	// plus 1 square foot of slack, for a total of 43 square feet.
	cases := []struct {
		box  Box
		area int
	}{
		{box: Box{l: 2, w: 3, h: 4}, area: 58},
		{box: Box{l: 1, w: 1, h: 10}, area: 43},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("box %v", c.box), func(t *testing.T) {
			actual := c.box.Area() + c.box.SmallArea()
			if actual != c.area {
				t.Fatalf("expected %v, got: %v", c.area, actual)
			}
		})
	}
}

func TestParseDimensions(t *testing.T) {
	cases := []struct {
		dimension string
		box       Box
	}{
		{dimension: "2x3x4", box: Box{l: 2, w: 3, h: 4}},
		{dimension: "1x1x10", box: Box{l: 1, w: 1, h: 10}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("parsing %s", c.dimension), func(t *testing.T) {
			actual := parseDimensions(c.dimension)
			if *actual != c.box {
				t.Fatalf("expected %v, got: %v", c.box, actual)
			}
		})
	}
}
