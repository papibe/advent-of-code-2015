package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

const INVALID = -1

func md5_hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func is_good(s string, zeros int) bool {
	for i := range zeros {
		if s[i] != '0' {
			return false
		}
	}
	return true
}

func solve(base string, zeros int) int {
	counter := 1

	for {
		str_counter := strconv.Itoa(counter)
		hash := md5_hash(base + str_counter)
		if is_good(hash, zeros) {
			break
		}
		counter++
	}
	return counter
}

func main() {
	fmt.Println("Part 1:", solve("iwrupvqb", 5)) // 346386
	fmt.Println("Part 2:", solve("iwrupvqb", 6)) // 9958218
}
