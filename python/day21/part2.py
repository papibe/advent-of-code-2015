from typing import Iterator, Tuple

WEAPONS = {
    "dagger": {
        "cost": 8,
        "damage": 4,
        "armor": 0,
    },
    "shortsword": {
        "cost": 10,
        "damage": 5,
        "armor": 0,
    },
    "warhammer": {
        "cost": 25,
        "damage": 6,
        "armor": 0,
    },
    "longsword": {
        "cost": 40,
        "damage": 7,
        "armor": 0,
    },
    "greataxe": {
        "cost": 74,
        "damage": 8,
        "armor": 0,
    },
}

ARMOR = {
    "Leather": {
        "cost": 13,
        "damage": 0,
        "armor": 1,
    },
    "Chainmail": {
        "cost": 31,
        "damage": 0,
        "armor": 2,
    },
    "Splintmail": {
        "cost": 53,
        "damage": 0,
        "armor": 3,
    },
    "Bandedmail": {
        "cost": 75,
        "damage": 0,
        "armor": 4,
    },
    "Platemail": {
        "cost": 102,
        "damage": 0,
        "armor": 5,
    },
    "Nothing": {
        "cost": 0,
        "damage": 0,
        "armor": 0,
    },
}

RINGS = {
    "Damage+1": {
        "cost": 25,
        "damage": 1,
        "armor": 0,
    },
    "Damage+2": {
        "cost": 50,
        "damage": 2,
        "armor": 0,
    },
    "Damage+3": {
        "cost": 100,
        "damage": 3,
        "armor": 0,
    },
    "Defense+1": {
        "cost": 20,
        "damage": 0,
        "armor": 1,
    },
    "Defense+2": {
        "cost": 40,
        "damage": 0,
        "armor": 2,
    },
    "Defense+3": {
        "cost": 80,
        "damage": 0,
        "armor": 3,
    },
    "Nothing": {
        "cost": 0,
        "damage": 0,
        "armor": 0,
    },
}


def get_selections() -> Iterator[Tuple[int, int, int, int]]:
    hp: int = 100

    for k, weapon in WEAPONS.items():
        for k, available_armor in ARMOR.items():
            for k, ring1 in RINGS.items():
                for k, ring2 in RINGS.items():

                    # skip duplicate ring, except when both are 'Nothing'
                    if ring1 == ring2 and ring1["cost"] != 0:
                        continue

                    cost: int = weapon["cost"]
                    cost += available_armor["cost"]
                    cost += ring1["cost"]
                    cost += ring2["cost"]

                    damage: int = weapon["damage"]
                    damage += available_armor["damage"]
                    damage += ring1["damage"]
                    damage += ring2["damage"]

                    armor = weapon["armor"]
                    armor += available_armor["armor"]
                    armor += ring1["armor"]
                    armor += ring2["armor"]

                    yield cost, hp, damage, armor


def solve(original_boss_hp: int, boss_damage: int, boss_armor: int) -> int:
    max_cost: int = float("-inf")  # type: ignore

    for cost, hp, damage, armor in get_selections():
        boss_hp: int = original_boss_hp

        while True:
            attack_damage: int = damage - boss_armor
            if attack_damage <= 0:
                raise ValueError("what")

            boss_hp -= attack_damage
            if boss_hp <= 0:
                break

            attack_damage = boss_damage - armor
            if attack_damage <= 0:
                attack_damage = 1

            hp -= attack_damage
            if hp <= 0:
                max_cost = max(max_cost, cost)
                break

    return max_cost


def solution(boss_hp: int, boss_damage: int, boss_armor: int) -> int:
    return solve(boss_hp, boss_damage, boss_armor)


if __name__ == "__main__":
    print(solution(100, 8, 2))  # 91
