from typing import Dict, List


def parse(filename: str) -> List[str]:
    with open(filename, "r") as fp:
        strings: List[str] = fp.read().splitlines()

    return strings


def is_nice(s: str) -> bool:
    # check rule 1
    rule1 = False
    pairs: Dict[str, List[int]] = {}

    for i in range(len(s) - 1):
        pair: str = s[i : i + 2]
        if pair not in pairs:
            pairs[pair] = []
        pairs[pair].append(i)

        n: int = len(pairs[pair])
        if n > 1:
            for j in range(n):
                for k in range(j + 1, n):
                    if pairs[pair][k] > pairs[pair][j] + 1:
                        rule1 = True
                        break
                else:
                    continue
                break
        if rule1:
            break

    # check rule 2
    rule2 = False
    for i in range(len(s) - 2):
        if s[i] == s[i + 2]:
            rule2 = True
            break

    return rule1 and rule2


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
