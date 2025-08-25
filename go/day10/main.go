package main

import (
	"fmt"
	"strconv"
)

func new_sequence(s string) string {
	output := []rune{}

	index := 0
	for index < len(s) {
		char := s[index]
		counter := 0
		for index < len(s) && s[index] == char {
			index++
			counter++
		}
		str_counter := strconv.Itoa(counter)
		output = append(output, rune(str_counter[0]))
		output = append(output, rune(char))
	}
	return string(output)
}

func cycle(s string, times int) string {
	for range times {
		s = new_sequence(s)
	}
	return s
}

func solution(puzzle_input string, times int) int {

	return len(cycle(puzzle_input, times))
}

func main() {
	fmt.Println("Part 1:", solution("1113222113", 40)) // 252594
	fmt.Println("Part 2:", solution("1113222113", 50)) // 3579328
}
