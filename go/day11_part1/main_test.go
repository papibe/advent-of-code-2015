package main

import (
	"fmt"
	"testing"
)

func TestFirstRule(t *testing.T) {
	testCases := []struct {
		password string
		expected bool
	}{
		{"hijklmmn", true},
		{"abbceffg", false},
		{"abbcegjk", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%t", tc.password, tc.expected), func(t *testing.T) {
			result := first_rule(convert(tc.password))
			if result != tc.expected {
				t.Errorf("got %t; want %t", result, tc.expected)
			}
		})
	}
}

func TestSecondRule(t *testing.T) {
	testCases := []struct {
		password string
		expected bool
	}{
		{"hijklmmn", false},
		{"abbceffg", true},
		{"abbcegjk", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%t", tc.password, tc.expected), func(t *testing.T) {
			result := second_rule(convert(tc.password))
			if result != tc.expected {
				t.Errorf("got %t; want %t", result, tc.expected)
			}
		})
	}
}

func TestThirdRule(t *testing.T) {
	testCases := []struct {
		password string
		expected bool
	}{
		{"hijklmmn", false},
		{"abbceffg", true},
		{"abbcegjk", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v_should_be_%t", tc.password, tc.expected), func(t *testing.T) {
			result := thrid_rule(convert(tc.password))
			if result != tc.expected {
				t.Errorf("got %t; want %t", result, tc.expected)
			}
		})
	}
}

func TestNextPassword(t *testing.T) {
	testCases := []struct {
		password string
		expected string
	}{
		{"ghijklmn", "ghjaabcc"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_should_be_%s", tc.password, tc.expected), func(t *testing.T) {
			result := solution(tc.password)
			if result != tc.expected {
				t.Errorf("got %s; want %s", result, tc.expected)
			}
		})
	}
}
