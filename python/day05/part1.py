from typing import List


def parse(filename: str) -> List[str]:
    with open(filename, "r") as fp:
        strings: List[str] = fp.read().splitlines()

    return strings


def is_nice(s: str) -> bool:
    # rule 1: It contains at least three vowels
    number_of_vowels: int = 0
    for char in s:
        if char in "aeiou":
            number_of_vowels += 1

    good_vowels: bool = number_of_vowels >= 3

    # rule 2: It contains at least one letter that appears twice in a row
    twice_in_a_row: int = 0
    for i in range(len(s) - 1):
        if s[i] == s[i + 1]:
            twice_in_a_row += 1

    good_in_a_row: bool = twice_in_a_row >= 1

    # rule 3: It does not contain the strings ab, cd, pq, or xy
    forbidden_strings: int = 0
    for i in range(len(s) - 1):
        if s[i : i + 2] in ["ab", "cd", "pq", "xy"]:
            forbidden_strings += 1

    does_not_contain_forbidden: bool = forbidden_strings == 0

    return good_vowels and good_in_a_row and does_not_contain_forbidden


def solve(strings: List[str]) -> int:
    total: int = 0

    for string in strings:
        if is_nice(string):
            total += 1

    return total


def solution(filename: str) -> int:
    strings: List[str] = parse(filename)
    return solve(strings)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 258
