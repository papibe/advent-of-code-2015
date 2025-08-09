from dataclasses import dataclass
from typing import Dict, Set, Tuple


@dataclass
class Position:
    row: int
    col: int


DIR_STEPS: Dict[str, Position] = {
    "^": Position(-1, 0),
    "v": Position(1, 0),
    ">": Position(0, 1),
    "<": Position(0, -1),
}


def parse(filename: str) -> str:
    with open(filename, "r") as fp:
        data: str = fp.read().strip()

    return data


def solve(data: str) -> int:
    houses: Set[Tuple[int, int]] = set()

    current_row = current_col = 0
    houses.add((current_row, current_col))

    for direction in data:
        current_row += DIR_STEPS[direction].row
        current_col += DIR_STEPS[direction].col
        houses.add((current_row, current_col))

    return len(houses)


def solution(filename: str) -> int:
    data: str = parse(filename)
    return solve(data)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 2081
