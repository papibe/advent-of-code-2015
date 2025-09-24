package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		json_file string
		expected  int
	}{
		{"./example1.txt", 6},
		{"./example2.txt", 6},
		{"./example3.txt", 3},
		{"./example4.txt", 3},
		{"./example5.txt", 0},
		{"./example6.txt", 0},
		{"./example7.txt", 0},
		{"./example8.txt", 0},
		{"./example9.txt", 4},
		{"./example10.txt", 0},
		{"./example11.txt", 6},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%d", tc.json_file, tc.expected), func(t *testing.T) {
			result := solution(tc.json_file)
			if result != tc.expected {
				t.Errorf("got %d; want %d", result, tc.expected)
			}
		})
	}
}
