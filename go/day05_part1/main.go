package main

import (
	"fmt"
	"os"
	"strings"
)

func parse(filename string) []string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	return strings.Split(strings.Trim(string(data), "\n"), "\n")
}

func is_nice(s string) bool {
	// rule 1: It contains at least three vowels
	number_of_vowels := 0
	for _, char := range s {
		if strings.Contains("aeiou", string(char)) {
			number_of_vowels++
		}
	}
	good_vowels := (number_of_vowels >= 3)

	// rule 2: It contains at least one letter that appears twice in a row
	twice_in_a_row := 0
	for i := range len(s) - 1 {
		if s[i] == s[i+1] {
			twice_in_a_row++
		}
	}
	good_in_a_row := (twice_in_a_row >= 1)

	// rule 3: It does not contain the strings ab, cd, pq, or xy
	forbidden_strings := 0
	for i := range len(s) - 1 {
		pair := s[i : i+2]
		for _, substrings := range []string{"ab", "cd", "pq", "xy"} {
			if pair == substrings {
				forbidden_strings++
				break
			}
		}
	}
	does_not_contain_forbidden := (forbidden_strings == 0)

	return good_vowels && good_in_a_row && does_not_contain_forbidden
}

func solve(strs []string) int {
	total := 0

	for _, str := range strs {
		if is_nice(str) {
			total++
		}
	}
	return total
}

func solution(filename string) int {
	strs := parse(filename)
	return solve(strs)
}

func main() {
	fmt.Println(solution("./input.txt")) // 258
}
