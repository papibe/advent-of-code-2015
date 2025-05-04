from dataclasses import dataclass
from typing import List


@dataclass
class Dimension:
    L: int
    w: int
    h: int


def parse(filename: str) -> List[Dimension]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    dimensions: List[Dimension] = []
    for line in data:
        parts: List[str] = line.split("x")
        L: int = int(parts[0])
        w: int = int(parts[1])
        h: int = int(parts[2])

        dimensions.append(Dimension(L, h, w))

    return dimensions


def solve(dimensions: List[Dimension]) -> int:
    wrapping_paper: int = 0

    for dim in dimensions:
        area: int = (2 * dim.L * dim.w) + (2 * dim.w * dim.h) + (2 * dim.h * dim.L)
        small_side: int = min(dim.L * dim.w, dim.w * dim.h, dim.h * dim.L)
        wrapping_paper += area + small_side

    return wrapping_paper


def solution(filename: str) -> int:
    dimensions: List[Dimension] = parse(filename)
    return solve(dimensions)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 1598415
