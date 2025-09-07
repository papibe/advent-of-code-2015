package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		str                string
		expected_length    int
		expected_in_memory int
	}{
		{`""`, 2, 6},
		{`"abc"`, 5, 9},
		{`"aaa\"aaa"`, 10, 16},
		{`\\x27"`, 6, 11},
		{`"nywbv\\"`, 9, 15},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be %d and %d", tc.str, tc.expected_length, tc.expected_in_memory), func(t *testing.T) {
			length, in_memory := encode(tc.str)
			if length != tc.expected_length || in_memory != tc.expected_in_memory {
				t.Errorf("got (%d, %d); want (%d, %d)", length, in_memory, tc.expected_length, tc.expected_in_memory)
			}
		})
	}
}
