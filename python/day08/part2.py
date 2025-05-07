import re
from typing import List, Tuple


def parse(filename: str) -> List[str]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    return data


def encode(string: str) -> Tuple[int, int]:

    re_quote: str = r'\\"'
    re_slash: str = r'\\(?!")(?!x[0-9a-f][0-9a-f])'
    re_hex: str = r"\\x[0-9a-f][0-9a-f]"

    length: int = len(string)

    # remove quotes
    current_in_memory: int = length - 2
    clean_string: str = string[1 : len(string) - 1]

    matches: List[str]

    matches = re.findall(re_quote, clean_string)
    if matches:
        current_in_memory += len(matches) * 2

    matches = re.findall(re_slash, clean_string)
    if matches:
        current_in_memory += len(matches)

    matches = re.findall(re_hex, clean_string)
    if matches:
        current_in_memory += len(matches)

    return length, current_in_memory + 6


def solve(data: List[str]) -> int:
    literals: int = 0
    memory: int = 0

    for string in data:
        length, current_in_memory = encode(string)

        literals += length
        memory += current_in_memory

    return memory - literals


def solution(filename: str) -> int:
    data: List[str] = parse(filename)
    return solve(data)


if __name__ == "__main__":
    print(solution("./example.txt"))  # 19
    print(solution("./input.txt"))  # 2117
