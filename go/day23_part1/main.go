package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InstructionType string

const (
	HLF  InstructionType = "hlf"
	TPL  InstructionType = "tpl"
	INC  InstructionType = "inc"
	JUMP InstructionType = "jmp"
	JIE  InstructionType = "jie"
	JIO  InstructionType = "jio"
)

type Instruction struct {
	operation InstructionType
	param1    string
	param2    string
}

var STR_TO_INSTRUCTIONTYPE = map[string]InstructionType{
	"hlf": HLF,
	"tpl": TPL,
	"inc": INC,
	"jmp": JUMP,
	"jie": JIE,
	"jio": JIO,
}

func parse(filename string) (map[string]int, []Instruction) {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	instructions := []Instruction{}
	registers := make(map[string]int)

	for _, line := range lines {
		split_line := strings.Split(line, " ")
		instr := STR_TO_INSTRUCTIONTYPE[split_line[0]]
		params := strings.Join(split_line[1:], "")

		var param1 string
		var param2 string

		switch instr {
		case HLF:
			param1 = split_line[1]
			registers[param1] = 0
		case TPL:
			param1 = split_line[1]
			registers[param1] = 0
		case INC:
			param1 = split_line[1]
			registers[param1] = 0
		case JUMP:
			param1 = split_line[1]
		case JIE:
			split_params := strings.Split(params, ",")
			param1 = split_params[0]
			param2 = split_params[1]
			registers[param1] = 0
		case JIO:
			split_params := strings.Split(params, ",")
			param1 = split_params[0]
			param2 = split_params[1]
			registers[param1] = 0
		}

		instruction := Instruction{instr, param1, param2}
		instructions = append(instructions, instruction)

	}
	return registers, instructions
}

func solve(
	instructions []Instruction, registers map[string]int, output_register string,
) int {
	pointer := 0

	for pointer < len(instructions) {
		instruction := instructions[pointer]
		instruction_type := instruction.operation

		switch instruction_type {
		case HLF:
			registers[instruction.param1] /= 2
		case TPL:
			registers[instruction.param1] *= 3
		case INC:
			registers[instruction.param1] += 1
		case JUMP:
			value, _ := strconv.Atoi(instruction.param1)
			pointer += value
			continue
		case JIE:
			if registers[instruction.param1]%2 == 0 {
				value, _ := strconv.Atoi(instruction.param2)
				pointer += value
				continue
			}
		case JIO:
			if registers[instruction.param1] == 1 {
				value, _ := strconv.Atoi(instruction.param2)
				pointer += value
				continue
			}
		}

		pointer++
	}
	return registers[output_register]
}

func solution(filename string, output_register string) int {
	registers, instructions := parse(filename)
	return solve(instructions, registers, output_register)
}

func main() {
	fmt.Println(solution("./example.txt", "a")) // 2
	fmt.Println(solution("./input.txt", "b"))   // 255
}
