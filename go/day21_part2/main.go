package main

import (
	"fmt"
	"math"
)

var WEAPONS = map[string]map[string]int{
	"dagger": {
		"cost":   8,
		"damage": 4,
		"armor":  0,
	},
	"shortsword": {
		"cost":   10,
		"damage": 5,
		"armor":  0,
	},
	"warhammer": {
		"cost":   25,
		"damage": 6,
		"armor":  0,
	},
	"longsword": {
		"cost":   40,
		"damage": 7,
		"armor":  0,
	},
	"greataxe": {
		"cost":   74,
		"damage": 8,
		"armor":  0,
	},
}

var ARMOR = map[string]map[string]int{
	"Leather": {
		"cost":   13,
		"damage": 0,
		"armor":  1,
	},
	"Chainmail": {
		"cost":   31,
		"damage": 0,
		"armor":  2,
	},
	"Splintmail": {
		"cost":   53,
		"damage": 0,
		"armor":  3,
	},
	"Bandedmail": {
		"cost":   75,
		"damage": 0,
		"armor":  4,
	},
	"Platemail": {
		"cost":   102,
		"damage": 0,
		"armor":  5,
	},
	"Nothing": {
		"cost":   0,
		"damage": 0,
		"armor":  0,
	},
}

var RINGS = map[string]map[string]int{
	"Damage+1": {
		"cost":   25,
		"damage": 1,
		"armor":  0,
	},
	"Damage+2": {
		"cost":   50,
		"damage": 2,
		"armor":  0,
	},
	"Damage+3": {
		"cost":   100,
		"damage": 3,
		"armor":  0,
	},
	"Defense+1": {
		"cost":   20,
		"damage": 0,
		"armor":  1,
	},
	"Defense+2": {
		"cost":   40,
		"damage": 0,
		"armor":  2,
	},
	"Defense+3": {
		"cost":   80,
		"damage": 0,
		"armor":  3,
	},
	"Nothing": {
		"cost":   0,
		"damage": 0,
		"armor":  0,
	},
}

func get_selections() [][4]int {
	hp := 100

	results := [][4]int{}

	for _, weapon := range WEAPONS {
		for _, available_armor := range ARMOR {
			for _, ring1 := range RINGS {
				for _, ring2 := range RINGS {

					// skip duplicate ring, except when both are 'Nothing'
					ring1_equal_to_ring2 := ring1["cost"] == ring2["cost"] &&
						ring1["damage"] == ring2["damage"] && ring1["armor"] == ring2["armor"]

					if ring1_equal_to_ring2 && ring1["cost"] != 0 {
						continue
					}

					cost := weapon["cost"]
					cost += available_armor["cost"]
					cost += ring1["cost"]
					cost += ring2["cost"]

					damage := weapon["damage"]
					damage += available_armor["damage"]
					damage += ring1["damage"]
					damage += ring2["damage"]

					armor := weapon["armor"]
					armor += available_armor["armor"]
					armor += ring1["armor"]
					armor += ring2["armor"]

					results = append(results, [4]int{cost, hp, damage, armor})
				}
			}
		}
	}

	return results
}

func solution(original_boss_hp, boss_damage, boss_armor int) int {
	max_cost := math.MinInt

	for _, item := range get_selections() {
		cost, hp, damage, armor := item[0], item[1], item[2], item[3]
		boss_hp := original_boss_hp

		for {
			attack_damage := damage - boss_armor

			boss_hp -= attack_damage
			if boss_hp <= 0 {
				break
			}
			attack_damage = boss_damage - armor
			if attack_damage <= 0 {
				attack_damage = 1
			}
			hp -= attack_damage
			if hp <= 0 {
				max_cost = max(max_cost, cost)
				break
			}
		}
	}
	return max_cost
}

func main() {
	fmt.Println(solution(100, 8, 2)) // 158
}
