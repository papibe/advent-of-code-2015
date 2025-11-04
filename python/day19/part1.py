import re
from dataclasses import dataclass
from typing import List, Match, Optional, Set, Tuple


@dataclass
class Rule:
    pattern: str
    replacement: str


def parse(filename: str) -> Tuple[str, List[Rule]]:
    with open(filename, "r") as fp:
        data: str = fp.read()

    blocks: List[str] = data.split("\n\n")

    # parse rules
    re_rule: str = r"(\w+) => (\w+)"
    rules: List[Rule] = []

    for line in blocks[0].splitlines():
        matches: Optional[Match[str]] = re.match(re_rule, line)
        assert matches is not None

        pattern: str = matches.group(1)
        replacement: str = matches.group(2)

        rules.append(Rule(pattern, replacement))

    # get molecule
    molecule: str = blocks[1].strip()

    return molecule, rules


def solve(molecule: str, rules: List[Rule]) -> int:
    new_molecules: Set[str] = set()

    for rule in rules:
        for matches in re.finditer(rule.pattern, molecule):
            start_index: int = matches.start()
            end_index: int = matches.end()

            new_molecules.add(
                molecule[:start_index] + rule.replacement + molecule[end_index:]
            )

    return len(new_molecules)


def solution(filename: str) -> int:
    molecule, rules = parse(filename)
    return solve(molecule, rules)


if __name__ == "__main__":
    print(solution("./example1.txt"))  # 4
    print(solution("./example2.txt"))  # 7
    print(solution("./input.txt"))  # 576
