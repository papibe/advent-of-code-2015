package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		dimension []Dimension
		expected  int
	}{
		{[]Dimension{{2, 3, 4}}, 34},
		{[]Dimension{{1, 1, 10}}, 14},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%d", tc.dimension, tc.expected), func(t *testing.T) {
			result := solve(tc.dimension)
			if result != tc.expected {
				t.Errorf("got %d; want %d", result, tc.expected)
			}
		})
	}
}
