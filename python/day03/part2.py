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

    santa_row = santa_col = 0
    robo_row = robo_col = 0

    houses.add((santa_row, santa_col))

    for idx_turn, direction in enumerate(data, 1):
        if idx_turn % 2 == 0:
            santa_row += DIR_STEPS[direction].row
            santa_col += DIR_STEPS[direction].col
            houses.add((santa_row, santa_col))

        else:
            robo_row += DIR_STEPS[direction].row
            robo_col += DIR_STEPS[direction].col
            houses.add((robo_row, robo_col))

    return len(houses)


def solution(filename: str) -> int:
    data: str = parse(filename)
    return solve(data)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 2341
