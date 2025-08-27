package main

import (
	"fmt"
	"strings"
)

const I_VALUE = 'i' - 'a'
const O_VALUE = 'o' - 'a'
const L_VALUE = 'l' - 'a'

func convert(p string) []int {
	output := make([]int, len(p))

	for index := range len(p) {
		output[index] = int(p[index]) - 'a'
	}
	return output
}

func unconvert(p []int) []string {
	output := make([]string, len(p))

	for index := range len(p) {
		output[index] = string(rune('a' + p[index]))
	}
	return output
}

func first_rule(p []int) bool {
	for index := range len(p) - 3 + 1 {
		if p[index] == p[index+1]-1 && p[index+1] == p[index+2]-1 {
			return true
		}
	}
	return false
}

func second_rule(p []int) bool {
	for _, value := range p {
		if value == I_VALUE || value == O_VALUE || value == L_VALUE {
			return false
		}
	}
	return true
}

func thrid_rule(p []int) bool {
	for i := range len(p) - 2 + 1 {
		if p[i] == p[i+1] {
			for j := i + 2; j < len(p)-2+1; j++ {
				if p[j] == p[j+1] {
					return true
				}
			}
		}
	}
	return false
}

func is_valid(p []int) bool {
	return first_rule(p) && second_rule(p) && thrid_rule(p)
}

func next_good(p []int, index int) []int {
	p[index] += 1
	for i := index + 1; i < len(p); i++ {
		p[i] = 0
	}
	return p
}

func skip_to_next(p []int) []int {
	for index := range len(p) {
		if p[index] == I_VALUE || p[index] == O_VALUE || p[index] == L_VALUE {
			return next_good(p, index)
		}
	}
	return p
}

func increment(p []int) []int {
	p = skip_to_next(p)

	index := len(p) - 1
	new_digit := p[index] + 1
	p[index] = new_digit % 26
	carry := new_digit / 26

	for index > 0 && carry != 0 {
		index -= 1

		new_digit = p[index] + 1
		p[index] = new_digit % 26
		carry = new_digit / 26
	}

	return p
}

func solution(str_password string) string {
	password := convert(str_password)

	for {
		password = increment(password)
		if is_valid(password) {
			return strings.Join(unconvert(password), "")
		}
	}
}

func main() {
	fmt.Println(solution("hepxcrrq")) // hepxxyzz
}
