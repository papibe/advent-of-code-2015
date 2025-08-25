package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parse(filename string) []string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	return strings.Split(strings.Trim(string(data), "\n"), "\n")
}

func measure(s string) (int, int) {
	re_quote := regexp.MustCompile(`\"`)
	re_slash := regexp.MustCompile(`\\\\`)
	re_hex := regexp.MustCompile(`\\x[0-9a-f][0-9a-f]`)

	length := len(s)

	// remove quotes
	current_in_memory := length - 2
	clean_string := s[1 : len(s)-1]

	var matches []string

	matches = re_quote.FindAllString(clean_string, -1)
	if len(matches) > 0 {
		current_in_memory -= len(matches)
	}

	matches = re_slash.FindAllString(clean_string, -1)
	if len(matches) > 0 {
		current_in_memory -= len(matches)
	}

	matches = re_hex.FindAllString(clean_string, -1)
	if len(matches) > 0 {
		current_in_memory -= len(matches) * 3
	}

	return length, current_in_memory
}

func solve(data []string) int {
	literals := 0
	memory := 0

	for _, str := range data {
		length, current_in_memory := measure(str)

		literals += length
		memory += current_in_memory
	}

	return literals - memory
}

func solution(filename string) int {
	data := parse(filename)
	return solve(data)
}

func main() {
	fmt.Println(solution("./example.txt")) // 12
	fmt.Println(solution("./input.txt"))   // 1371
}
