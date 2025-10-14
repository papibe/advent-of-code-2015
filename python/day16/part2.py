import re
from typing import Dict, List, Match, Optional

AUNT_SUE_EXACTLY: Dict[str, int] = {
    "children": 3,
    "samoyeds": 2,
    "akitas": 0,
    "vizslas": 0,
    "cars": 2,
    "perfumes": 1,
}

AUNT_SUE_GREATER: Dict[str, int] = {
    "cats": 7,
    "trees": 3,
}

AUNT_SUE_LOWER: Dict[str, int] = {
    "pomeranians": 3,
    "goldfish": 5,
}


def parse(filename: str) -> List[Dict[str, int]]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    re_line: str = r"Sue (\d+): (.*)"
    re_thing: str = r"(\w+): (\d+)"
    aunts = []
    for line in data:
        matches: Optional[Match[str]] = re.match(re_line, line)
        assert matches is not None

        things: str = matches.group(2)

        aunt_things = {}

        for thing in things.split(", "):
            thing_match: Optional[Match[str]] = re.match(re_thing, thing)
            assert thing_match is not None

            name: str = thing_match.group(1)
            amount: str = thing_match.group(2)
            aunt_things[name] = int(amount)

        aunts.append(aunt_things)

    return aunts


def solve(aunts: List[Dict[str, int]]) -> int:
    max_score: int = 0
    max_index: int = -1
    max_scores = set()

    for index, aunt in enumerate(aunts, 1):
        aunt_score: int = 0

        for k, v in aunt.items():
            # exact count
            if k in AUNT_SUE_EXACTLY and AUNT_SUE_EXACTLY[k] == v:
                aunt_score += 1

            # greater count
            if k in AUNT_SUE_GREATER and v > AUNT_SUE_GREATER[k]:
                aunt_score += 1

            # fewer count
            if k in AUNT_SUE_LOWER and v < AUNT_SUE_LOWER[k]:
                aunt_score += 1

        if aunt_score > max_score:
            max_score = aunt_score
            max_scores = set([aunt_score])
            max_index = index

        elif aunt_score == max_score:
            max_scores.add(aunt_score)

    assert len(max_scores) == 1

    return max_index


def solution(filename: str) -> int:
    aunts: List[Dict[str, int]] = parse(filename)
    return solve(aunts)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 323
