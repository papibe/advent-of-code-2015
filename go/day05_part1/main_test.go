package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		strs     string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%t", tc.strs, tc.expected), func(t *testing.T) {
			result := is_nice(tc.strs)
			if result != tc.expected {
				t.Errorf("got %t; want %t", result, tc.expected)
			}
		})
	}
}
