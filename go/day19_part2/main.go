package main

import (
	"fmt"
	"math/rand/v2"
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

func reduce(original_molecule string, rules []Rule) (int, string) {
	molecule := original_molecule

	counter := 0
	for {
		changes := false
		for _, rule := range rules {
			if rule.pattern == "e" {
				continue
			}
			// count how many matches of target first
			substitutions := strings.Count(molecule, rule.replacement)

			// then replace them all
			if substitutions > 0 {
				molecule = strings.Replace(molecule, rule.replacement, rule.pattern, 1)
				changes = true
				counter++
				break
			}
		}
		if !changes {
			break
		}
	}

	for _, rule := range rules {
		if rule.pattern != "e" {
			continue
		}
		// count how many matches of target first
		substitutions := strings.Count(molecule, rule.replacement)

		// then replace them all
		if substitutions > 0 {
			molecule = strings.Replace(molecule, rule.replacement, rule.pattern, 1)
			counter++
		}
	}
	return counter, molecule
}

func shuffle(l *[]Rule) {
	n := len(*l)
	for i := n - 1; i > 0; i-- {
		j := rand.IntN(n)
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
	}
}

func solve(molecule string, rules []Rule) int {
	shuffle(&rules)
	counter, new_molecule := reduce(molecule, rules)
	for new_molecule != "e" {
		shuffle(&rules)
		counter, new_molecule = reduce(molecule, rules)
	}
	return counter
}

func solution(filename string) int {
	molecule, rules := parse(filename)
	return solve(molecule, rules)
}

func main() {
	fmt.Println(solution("./example3.txt")) // 3
	fmt.Println(solution("./example4.txt")) // 6
	fmt.Println(solution("./input.txt"))    // 207
}
