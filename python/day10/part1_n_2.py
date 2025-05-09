from typing import List


def new_sequence(s: List[str]) -> List[str]:
    output: List[str] = []

    index: int = 0
    while index < len(s):
        char: str = s[index]
        counter: int = 0
        while index < len(s) and s[index] == char:
            index += 1
            counter += 1

        assert 1 <= counter <= 9
        output.append(str(counter))
        output.append(char)

    return output


def cycle(s: List[str], times: int) -> List[str]:
    for _ in range(times):
        s = new_sequence(s)
    return s


def solution(puzzle_input: str, times: int) -> int:
    sequence: List[str] = list(puzzle_input)
    return len(cycle(sequence, times))


if __name__ == "__main__":
    print("Part 1:", solution("1113222113", 40))  # 252594
    print("Part 2:", solution("1113222113", 50))  # 3579328
