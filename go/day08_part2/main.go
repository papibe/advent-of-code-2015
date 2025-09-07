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

func encode(s string) (int, int) {
	re_hex := regexp.MustCompile(`\\x[0-9a-f][0-9a-f]`)
	re_quote := regexp.MustCompile(`\\"`)
	re_slash := regexp.MustCompile(`\\`)

	length := len(s)

	// remove quotes
	current_in_memory := length - 2
	clean_string := s[1 : len(s)-1]

	hex_matches := re_hex.FindAllString(clean_string, -1)
	quote_matches := re_quote.FindAllString(clean_string, -1)
	slash_matches := re_slash.FindAllString(clean_string, -1)

	hex_count := len(hex_matches)
	slash_count := len(slash_matches)
	quote_count := len(quote_matches)

	current_in_memory += hex_count + quote_count*2 + (slash_count - hex_count - quote_count)

	return length, current_in_memory + 6
}

func solve(data []string) int {
	literals := 0
	memory := 0

	for _, str := range data {
		length, current_in_memory := encode(str)

		literals += length
		memory += current_in_memory
	}

	return memory - literals
}

func solution(filename string) int {
	data := parse(filename)
	return solve(data)
}

func main() {
	fmt.Println(solution("./example.txt")) // 19
	fmt.Println(solution("./input.txt"))   // 2117
}
