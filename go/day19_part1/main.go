package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Rule struct {
	pattern     string
	replacement string
}

func parse(filename string) (string, []Rule) {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}
	blocks := strings.Split(strings.Trim(string(data), "\n"), "\n\n")

	// get molecule
	molecule := strings.Trim(blocks[1], "\n")

	// parse rules
	re_rule := regexp.MustCompile(`(\w+) => (\w+)`)
	rules := []Rule{}

	lines := strings.Split(blocks[0], "\n")
	for _, line := range lines {
		matches := re_rule.FindStringSubmatch(line)

		pattern := matches[1]
		replacement := matches[2]

		rules = append(rules, Rule{pattern, replacement})
	}
	return molecule, rules
}

func solve(molecule string, rules []Rule) int {
	new_molecules := NewSet[string]()

	for _, rule := range rules {
		re_pattern := regexp.MustCompile(rule.pattern)
		for _, indexes := range re_pattern.FindAllStringIndex(molecule, -1) {
			start_index := indexes[0]
			end_index := indexes[1]

			new_molecules.add(molecule[:start_index] + rule.replacement + molecule[end_index:])
		}
	}
	return new_molecules.len()
}

func solution(filename string) int {
	molecule, rules := parse(filename)
	return solve(molecule, rules)
}

func main() {
	fmt.Println(solution("./example1.txt")) // 4
	fmt.Println(solution("./example2.txt")) // 7
	fmt.Println(solution("./input.txt"))    // 576
}
