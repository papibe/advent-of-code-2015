package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operation string

const (
	ASSIGNMENT Operation = "ASSIGNMENT"
	AND        Operation = "AND"
	OR         Operation = "OR"
	LSHIFT     Operation = "LSHIFT"
	RSHIFT     Operation = "RSHIFT"
	NOT        Operation = "NOT"
)

type Instruction struct {
	operation Operation
	param1    string
	param2    string
	wire      string
	depends   []string
}

var STR_TO_OPERATION = map[string]Operation{
	"ASSIGNMENT": ASSIGNMENT,
	"AND":        AND,
	"OR":         OR,
	"LSHIFT":     LSHIFT,
	"RSHIFT":     RSHIFT,
	"NOT":        NOT,
}

func parse(filename string) (map[string]Instruction, []Instruction) {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	wires := make(map[string]Instruction)
	instructions := []Instruction{}

	re_assignment := regexp.MustCompile(`(\w+) -> (\w+)`)
	re_operation := regexp.MustCompile(`(\w+) (\w+) (\w+) -> (\w+)`)
	re_not := regexp.MustCompile(`NOT (\w+) -> (\w+)`)

	for _, line := range lines {
		assignment_match := re_assignment.FindStringSubmatch(line)
		operation_match := re_operation.FindStringSubmatch(line)
		not_match := re_not.FindStringSubmatch(line)

		var operation Operation
		var param1 string
		var param2 string
		var wire string

		if len(not_match) > 0 {
			operation = NOT
			param1 = not_match[1]
			wire = not_match[2]
		} else if len(operation_match) > 0 {
			param1 = operation_match[1]
			operation = STR_TO_OPERATION[operation_match[2]]
			param2 = operation_match[3]
			wire = operation_match[4]
		} else if len(assignment_match) > 0 {
			operation = ASSIGNMENT
			param1 = assignment_match[1]
			wire = assignment_match[2]
		}

		instruction := Instruction{operation, param1, param2, wire, []string{}}
		instructions = append(instructions, instruction)

		wires[wire] = instruction
	}
	return wires, instructions
}

func run_instructions(
	wire string, wires map[string]Instruction, wires_values *map[string]int,
) {

	instruction := wires[wire]
	var value int

	switch instruction.operation {
	case NOT:
		_, in_wires := wires[instruction.param1]
		if in_wires {
			value = (*wires_values)[instruction.param1]
		} else {
			value, _ = strconv.Atoi(instruction.param1)
		}
		(*wires_values)[wire] = value ^ 0xFFFF

	case ASSIGNMENT:
		_, in_wires := wires[instruction.param1]
		if in_wires {
			value = (*wires_values)[instruction.param1]
		} else {
			value, _ = strconv.Atoi(instruction.param1)
		}
		(*wires_values)[wire] = value

	case AND:
		var value1, value2 int
		_, in_wires := wires[instruction.param1]
		if in_wires {
			value1 = (*wires_values)[instruction.param1]
		} else {
			value1, _ = strconv.Atoi(instruction.param1)
		}

		_, in_wires = wires[instruction.param2]
		if in_wires {
			value2 = (*wires_values)[instruction.param2]
		} else {
			value2, _ = strconv.Atoi(instruction.param2)
		}
		(*wires_values)[wire] = value1 & value2

	case OR:
		var value1, value2 int
		_, in_wires := wires[instruction.param1]
		if in_wires {
			value1 = (*wires_values)[instruction.param1]
		} else {
			value1, _ = strconv.Atoi(instruction.param1)
		}

		_, in_wires = wires[instruction.param2]
		if in_wires {
			value2 = (*wires_values)[instruction.param2]
		} else {
			value2, _ = strconv.Atoi(instruction.param2)
		}
		(*wires_values)[wire] = value1 | value2

	case LSHIFT:
		var value1, value2 int
		_, in_wires := wires[instruction.param1]
		if in_wires {
			value1 = (*wires_values)[instruction.param1]
		} else {
			value1, _ = strconv.Atoi(instruction.param1)
		}

		_, in_wires = wires[instruction.param2]
		if in_wires {
			value2 = (*wires_values)[instruction.param2]
		} else {
			value2, _ = strconv.Atoi(instruction.param2)
		}
		(*wires_values)[wire] = value1 << value2

	case RSHIFT:
		var value1, value2 int
		_, in_wires := wires[instruction.param1]
		if in_wires {
			value1 = (*wires_values)[instruction.param1]
		} else {
			value1, _ = strconv.Atoi(instruction.param1)
		}

		_, in_wires = wires[instruction.param2]
		if in_wires {
			value2 = (*wires_values)[instruction.param2]
		} else {
			value2, _ = strconv.Atoi(instruction.param2)
		}
		(*wires_values)[wire] = value1 >> value2

	default:
		panic("Instruction unknown")
	}
}

func build_dependencies(
	wires map[string]Instruction, instructions []Instruction,
) map[string][]string {

	wire_dependencies := make(map[string][]string)
	for wire := range wires {
		wire_dependencies[wire] = []string{}
	}

	// update dependencies
	for _, instruction := range instructions {
		switch instruction.operation {
		case NOT:
			_, in_wires := wires[instruction.param1]
			if in_wires {
				instruction.depends = append(instruction.depends, instruction.param1)
				_, in_wire_dependencies := wire_dependencies[instruction.wire]
				if !in_wire_dependencies {
					wire_dependencies[instruction.wire] = []string{}
				}
				wire_dependencies[instruction.wire] = append(wire_dependencies[instruction.wire], instruction.param1)
			}

		case ASSIGNMENT:
			_, in_wires := wires[instruction.param1]
			if in_wires {
				instruction.depends = append(instruction.depends, instruction.param1)
				_, in_wire_dependencies := wire_dependencies[instruction.wire]
				if !in_wire_dependencies {
					wire_dependencies[instruction.wire] = []string{}
				}
				wire_dependencies[instruction.wire] = append(wire_dependencies[instruction.wire], instruction.param1)
			}

		case AND, OR, LSHIFT, RSHIFT:
			_, in_wires := wires[instruction.param1]
			if in_wires {
				instruction.depends = append(instruction.depends, instruction.param1)
				_, in_wire_dependencies := wire_dependencies[instruction.wire]
				if !in_wire_dependencies {
					wire_dependencies[instruction.wire] = []string{}
				}
				wire_dependencies[instruction.wire] = append(wire_dependencies[instruction.wire], instruction.param1)
			}

			_, in_wires = wires[instruction.param2]
			if in_wires {
				instruction.depends = append(instruction.depends, instruction.param2)
				_, in_wire_dependencies := wire_dependencies[instruction.wire]
				if !in_wire_dependencies {
					wire_dependencies[instruction.wire] = []string{}
				}
				wire_dependencies[instruction.wire] = append(wire_dependencies[instruction.wire], instruction.param2)
			}

		default:
			panic("Unknown instruction")

		}
	}

	return wire_dependencies
}

func kahn_top_sort(wire_dependencies map[string][]string) []string {
	in_degree := make(map[string]int)

	for wire := range wire_dependencies {
		in_degree[wire] = 0
	}

	for _, dependencies := range wire_dependencies {
		for _, neighbor_wire := range dependencies {
			in_degree[neighbor_wire]++
		}
	}

	queue := NewQueue[string]()
	for wire, deps := range in_degree {
		if deps == 0 {
			queue.append(wire)
		}
	}

	ordered_output := []string{}

	for !queue.is_empty() {
		wire := queue.popleft()
		ordered_output = append(ordered_output, wire)

		for _, neighbor_wire := range wire_dependencies[wire] {
			in_degree[neighbor_wire]--
			if in_degree[neighbor_wire] == 0 {
				queue.append(neighbor_wire)
			}
		}
	}
	return ordered_output
}

func solve(wires map[string]Instruction, instructions []Instruction) map[string]int {
	wire_dependencies := build_dependencies(wires, instructions)
	sorted_wires := kahn_top_sort(wire_dependencies)

	wires_values := make(map[string]int)

	for index := len(sorted_wires) - 1; index >= 0; index-- {
		run_instructions(sorted_wires[index], wires, &wires_values)
	}

	return wires_values
}

func solution(filename string) int {
	wires, instructions := parse(filename)
	wire_values := solve(wires, instructions)

	return wire_values["a"]
}

func main() {
	fmt.Println(solution("./input.txt")) // 16076
}
