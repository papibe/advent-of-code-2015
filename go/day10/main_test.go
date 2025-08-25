package main

import (
	"testing"
)

var expected = map[string]int{
	"d": 72,
	"e": 507,
	"f": 492,
	"g": 114,
	"h": 65412,
	"i": 65079,
	"x": 123,
	"y": 456,
}

func TestSolePart1ExampleData(t *testing.T) {
	wires, instructions := parse("./example.txt")
	wire_values := solve(wires, instructions)

	for wire, value := range expected {
		if value != wire_values[wire] {
			t.Errorf("got %d; want %d", wire_values[wire], value)
		}
	}
}
