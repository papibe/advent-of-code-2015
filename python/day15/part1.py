import re
from dataclasses import dataclass
from typing import List, Match, Optional


@dataclass
class Ingredient:
    name: str
    capacity: int
    durability: int
    flavor: int
    texture: int
    calories: int


def parse(filename: str) -> List[Ingredient]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    re_line: str = (
        r"(\w+): capacity ([0-9-]+), durability ([0-9-]+), flavor ([0-9-]+), texture ([0-9-]+), calories ([0-9-]+)"
    )

    ingredients: List[Ingredient] = []

    for line in data:
        matches: Optional[Match[str]] = re.match(re_line, line)
        assert matches is not None

        name: str = matches.group(1)
        capacity: int = int(matches.group(2))
        durability: int = int(matches.group(3))
        flavor: int = int(matches.group(4))
        texture: int = int(matches.group(5))
        calories: int = int(matches.group(6))

        ingredients.append(
            Ingredient(name, capacity, durability, flavor, texture, calories)
        )

    return ingredients


def partition(n: int, k: int) -> List[List[int]]:

    solutions: List[List[int]] = []

    def _partitions(n: int, k: int, numbers: List[int]) -> None:
        if k == 0:
            if n == 0:
                solutions.append(numbers)
            return
        if n == 0:
            return
        for i in range(1, n - k + 1 + 1):
            _partitions(n - i, k - 1, numbers + [i])

    _partitions(n, k, [])
    return solutions


def solve(ingredients: List[Ingredient]) -> int:
    highest_scoring: int = 0
    n: int = len(ingredients)

    for part_list in partition(100, n):
        capacity_score: int = 0
        durability_score: int = 0
        flavor_score: int = 0
        texture_score: int = 0

        for index, ingredient in enumerate(ingredients):
            capacity_score += ingredient.capacity * part_list[index]
            durability_score += ingredient.durability * part_list[index]
            flavor_score += ingredient.flavor * part_list[index]
            texture_score += ingredient.texture * part_list[index]

        score: int
        if (
            capacity_score < 0
            or durability_score < 0
            or flavor_score < 0
            or texture_score < 0
        ):
            score = 0
        else:
            score = capacity_score * durability_score * flavor_score * texture_score

        highest_scoring = max(highest_scoring, score)

    return highest_scoring


def solution(filename: str) -> int:
    ingredients: List[Ingredient] = parse(filename)
    return solve(ingredients)


if __name__ == "__main__":
    print(solution("./example.txt"))  # 62842880
    print(solution("./input.txt"))  # 18965440
