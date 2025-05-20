import re
from typing import Dict, List, Match, Optional

AUNT_SUE_THINGS: Dict[str, int] = {
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1,
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
            if k in AUNT_SUE_THINGS and AUNT_SUE_THINGS[k] == v:
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
    print(solution("./input.txt"))  # 213
