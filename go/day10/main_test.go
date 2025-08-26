package main

import (
	"fmt"
	"testing"
)

func TestCycle(t *testing.T) {
	testCases := []struct {
		puzzle_input string
		times        int
		expected     string
	}{
		{"1", 1, "11"},
		{"11", 1, "21"},
		{"21", 1, "1211"},
		{"1211", 1, "111221"},
		{"111221", 1, "312211"},
		{"1", 5, "312211"},
		{"1", 4, "111221"},
		{"1", 3, "1211"},
		{"1", 2, "21"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%v", tc.puzzle_input, tc.expected), func(t *testing.T) {
			result := cycle(tc.puzzle_input, tc.times)
			if result != tc.expected {
				t.Errorf("got %v; want %v", result, tc.expected)
			}
		})
	}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		puzzle_input string
		times        int
		expected     int
	}{
		{"1", 1, 2},
		{"11", 1, 2},
		{"21", 1, 4},
		{"1211", 1, 6},
		{"111221", 1, 6},
		{"1", 5, 6},
		{"1", 4, 6},
		{"1", 3, 4},
		{"1", 2, 2},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%d", tc.puzzle_input, tc.expected), func(t *testing.T) {
			result := solution(tc.puzzle_input, tc.times)
			if result != tc.expected {
				t.Errorf("got %d; want %d", result, tc.expected)
			}
		})
	}
}
