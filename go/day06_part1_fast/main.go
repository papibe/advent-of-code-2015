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
	ON     Operation = "turn on"
	OFF    Operation = "turn off"
	TOGGLE Operation = "toggle"
)

var STR_TO_OPERATION = map[string]Operation{
	"turn on":  ON,
	"turn off": OFF,
	"toggle":   TOGGLE,
}

type Instruction struct {
	operation Operation
	row1      int
	col1      int
	row2      int
	col2      int
}

func parse(filename string) []Instruction {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	re_line := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	instructions := []Instruction{}

	for _, line := range lines {
		matches := re_line.FindStringSubmatch(line)
		operation := STR_TO_OPERATION[matches[1]]
		row1, _ := strconv.Atoi(matches[2])
		col1, _ := strconv.Atoi(matches[3])
		row2, _ := strconv.Atoi(matches[4])
		col2, _ := strconv.Atoi(matches[5])

		instructions = append(instructions, Instruction{operation, row1, col1, row2, col2})
	}
	return instructions
}

func solve(instructions []Instruction) int {
	lights := make([][]int, 1000)
	for row := range 1000 {
		lights[row] = make([]int, 1000)
	}

	for _, instruction := range instructions {
		switch instruction.operation {
		case ON:
			for row := instruction.row1; row <= instruction.row2; row++ {
				for col := instruction.col1; col <= instruction.col2; col++ {
					lights[row][col] = 1
				}
			}

		case OFF:
			for row := instruction.row1; row <= instruction.row2; row++ {
				for col := instruction.col1; col <= instruction.col2; col++ {
					lights[row][col] = 0
				}
			}

		case TOGGLE:
			for row := instruction.row1; row <= instruction.row2; row++ {
				for col := instruction.col1; col <= instruction.col2; col++ {
					if lights[row][col] == 0 {
						lights[row][col] = 1
					} else {
						lights[row][col] = 0
					}
				}
			}

		default:
			panic("Instruction unknown")
		}
	}

	// count lights on
	lights_on := 0
	for row := range 1000 {
		for col := range 1000 {
			lights_on += lights[row][col]
		}
	}

	return lights_on
}

func solution(filename string) int {
	instructions := parse(filename)
	return solve(instructions)
}

func main() {
	fmt.Println(solution("./input.txt")) // 377891
}
