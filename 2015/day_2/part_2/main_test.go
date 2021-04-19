package main

import "testing"

func TestBox_Ribbon(t *testing.T) {
	cases := []struct {
		box    Box
		ribbon int
	}{
		{box: Box{l: 2, w: 3, h: 4}, ribbon: 10},
		{box: Box{l: 1, w: 1, h: 10}, ribbon: 4},
	}

	for _, c := range cases {
		actual := c.box.Ribbon()
		if actual != c.ribbon {
			t.Fatalf("expected 0, got: %v", actual)
		}
	}
}

func TestBox_Bow(t *testing.T) {
	cases := []struct {
		box Box
		bow int
	}{
		{box: Box{l: 2, w: 3, h: 4}, bow: 24},
		{box: Box{l: 1, w: 1, h: 10}, bow: 10},
	}

	for _, c := range cases {
		actual := c.box.Bow()
		if actual != c.bow {
			t.Fatalf("expected 0, got: %v", actual)
		}
	}
}
