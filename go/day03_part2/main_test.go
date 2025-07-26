package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		directions string
		expected   int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%d", tc.directions, tc.expected), func(t *testing.T) {
			result := solve(tc.directions)
			if result != tc.expected {
				t.Errorf("got %d; want %d", result, tc.expected)
			}
		})
	}
}
