package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func parse(filename string) any {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic("file error")
	}

	var generic_data any

	err = json.Unmarshal(data, &generic_data)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	return generic_data
}

func is_map(data any) (map[string]any, bool) {
	data_map, ok := data.(map[string]any)
	return data_map, ok
}

func is_list(data any) ([]any, bool) {
	data_list, ok := data.([]any)
	return data_list, ok
}

func is_int(data any) (float64, bool) {
	value, ok := data.(float64)
	return value, ok
}

func is_srt(data any) (string, bool) {
	value, ok := data.(string)
	return value, ok
}

func solve(data any) int {
	var dfs_parse func(any) float64

	dfs_parse = func(data any) float64 {
		// it's value
		if value, ok := is_int(data); ok {
			return value
		}
		total_sum := 0.0

		// it's a list
		if data_list, ok := is_list(data); ok {
			for _, item := range data_list {
				total_sum += dfs_parse(item)
			}
		}

		// it's a map
		if data_map, ok := is_map(data); ok {
			local_map_sum := 0.0
			red_found := false
			for _, value := range data_map {
				if value_str, ok := is_srt(value); ok && value_str == "red" {
					red_found = true
					break
				}
				local_map_sum += dfs_parse(value)
			}
			if !red_found {
				total_sum += float64(local_map_sum)
			}
		}

		return total_sum
	}

	return int(dfs_parse(data))
}

func solution(filename string) int {
	data := parse(filename)
	return solve(data)
}

func main() {
	fmt.Println(solution("./input.txt")) // 68466
}
