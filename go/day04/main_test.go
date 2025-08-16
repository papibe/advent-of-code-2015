package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		input_str string
		zeros     int
		expected  int
	}{
		{"abcdef", 5, 609043},
		{"pqrstuv", 5, 1048970},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%d", tc.input_str, tc.expected), func(t *testing.T) {
			result := solve(tc.input_str, tc.zeros)
			if result != tc.expected {
				t.Errorf("got %d; want %d", result, tc.expected)
			}
		})
	}
}
