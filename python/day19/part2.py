import re
from dataclasses import dataclass
from random import shuffle
from typing import List, Match, Optional, Tuple


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


def reduce(original_molecule: str, rules: List[Rule]) -> Tuple[int, str]:

    molecule: str = original_molecule

    counter: int = 0
    while True:

        changes = False
        for rule in rules:

            if rule.pattern == "e":
                continue

            # count how many matches of target first
            substitutions = molecule.count(rule.replacement)

            # then replace them all
            if substitutions > 0:
                molecule = molecule.replace(rule.replacement, rule.pattern, 1)
                changes = True
                counter += 1
                break

        if not changes:
            break

    for rule in rules:
        if rule.pattern != "e":
            continue

        # count how many matches of target first
        substitutions = molecule.count(rule.replacement)

        # then replace them all
        if substitutions > 0:
            molecule = molecule.replace(rule.replacement, rule.pattern, 1)
            changes = True
            counter += 1

    return counter, molecule


def solve(molecule: str, rules: List[Rule]) -> int:
    shuffle(rules)
    counter, new_molecule = reduce(molecule, rules)
    while new_molecule != "e":
        shuffle(rules)
        counter, new_molecule = reduce(molecule, rules)

    return counter


def solution(filename: str) -> int:
    molecule, rules = parse(filename)
    return solve(molecule, rules)


if __name__ == "__main__":
    print(solution("./example3.txt"))  # 3
    print(solution("./example4.txt"))  # 6
    print(solution("./input.txt"))  # 207
