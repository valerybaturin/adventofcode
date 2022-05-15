package main

import (
	"testing"
)

func Test_countDepthIncrease(t *testing.T) {
	// 199 (N/A - no previous measurement)
	// 200 (increased)
	// 208 (increased)
	// 210 (increased)
	// 200 (decreased)
	// 207 (increased)
	// 240 (increased)
	// 269 (increased)
	// 260 (decreased)
	// 263 (increased)
	cases := []struct {
		depths   []int
		expected int
	}{
		// In this example, there are 7 measurements that are larger than the previous measurement.
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
	}

	for _, c := range cases {
		t.Run("depth test", func(t *testing.T) {
			actual := countDepthIncrease(c.depths)
			if actual != c.expected {
				t.Errorf("expected %v, got %v", c.expected, actual)
			}
		})
	}
}
