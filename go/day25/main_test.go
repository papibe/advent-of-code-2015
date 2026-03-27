package main

import (
	"fmt"
	"testing"
)

func TestPosition(t *testing.T) {
	testCases := []struct {
		row      int
		col      int
		expected int
	}{
		{1, 1, 1},
		{1, 2, 3},
		{1, 3, 6},
		{1, 4, 10},
		{1, 5, 15},
		{1, 6, 21},
		{2, 1, 2},
		{2, 2, 5},
		{2, 3, 9},
		{2, 4, 14},
		{2, 5, 20},
		{3, 1, 4},
		{3, 2, 8},
		{3, 3, 13},
		{3, 4, 19},
		{4, 1, 7},
		{4, 2, 12},
		{4, 3, 18},
		{5, 1, 11},
		{5, 2, 17},
		{6, 1, 16},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("at (%d, %d) should_be_%d", tc.row, tc.col, tc.expected), func(t *testing.T) {
			result := get_order(tc.row, tc.col)
			if result != tc.expected {
				t.Errorf("got %d; want %d", result, tc.expected)
			}
		})
	}
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		row      int
		col      int
		expected int
	}{
		{1, 1, 20151125},
		{1, 2, 18749137},
		{1, 3, 17289845},
		{1, 4, 30943339},
		{1, 5, 10071777},
		{1, 6, 33511524},
		{2, 1, 31916031},
		{2, 2, 21629792},
		{2, 3, 16929656},
		{2, 4, 7726640},
		{2, 5, 15514188},
		{2, 6, 4041754},
		{3, 1, 16080970},
		{3, 2, 8057251},
		{3, 3, 1601130},
		{3, 4, 7981243},
		{3, 5, 11661866},
		{3, 6, 16474243},
		{4, 1, 24592653},
		{4, 2, 32451966},
		{4, 3, 21345942},
		{4, 4, 9380097},
		{4, 5, 10600672},
		{4, 6, 31527494},
		{5, 1, 77061},
		{5, 2, 17552253},
		{5, 3, 28094349},
		{5, 4, 6899651},
		{5, 5, 9250759},
		{5, 6, 31663883},
		{6, 1, 33071741},
		{6, 2, 6796745},
		{6, 3, 25397450},
		{6, 4, 24659492},
		{6, 5, 1534922},
		{6, 6, 27995004},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("at (%d, %d) should_be_%d", tc.row, tc.col, tc.expected), func(t *testing.T) {
			result := solution(tc.row, tc.col)
			if result != tc.expected {
				t.Errorf("got %d; want %d", result, tc.expected)
			}
		})
	}
}
