import re
from dataclasses import dataclass
from enum import Enum
from typing import Dict, List, Match, Optional, Set, Tuple


class Operation(Enum):
    ON = "turn on"
    OFF = "turn off"
    TOGGLE = "toggle"


@dataclass
class Instruction:
    operation: Operation
    row1: int
    col1: int
    row2: int
    col2: int


STR_TO_OPERATION: Dict[str, Operation] = {
    "turn on": Operation.ON,
    "turn off": Operation.OFF,
    "toggle": Operation.TOGGLE,
}


def parse(filename: str) -> List[Instruction]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    re_line: str = r"(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)"

    instructions: List[Instruction] = []

    for line in data:
        matches: Optional[Match[str]] = re.match(re_line, line)
        assert matches is not None
        operation: Operation = STR_TO_OPERATION[matches.group(1)]
        row1: int = int(matches.group(2))
        col1: int = int(matches.group(3))
        row2: int = int(matches.group(4))
        col2: int = int(matches.group(5))

        instructions.append(Instruction(operation, row1, col1, row2, col2))

    return instructions


def solve(instructions: List[Instruction]) -> int:
    lights: Set[Tuple[int, int]] = set()

    for instruction in instructions:

        match instruction.operation:

            case Operation.ON:
                for row in range(instruction.row1, instruction.row2 + 1):
                    for col in range(instruction.col1, instruction.col2 + 1):
                        lights.add((row, col))

            case Operation.OFF:
                for row in range(instruction.row1, instruction.row2 + 1):
                    for col in range(instruction.col1, instruction.col2 + 1):
                        lights.discard((row, col))

            case Operation.TOGGLE:
                for row in range(instruction.row1, instruction.row2 + 1):
                    for col in range(instruction.col1, instruction.col2 + 1):
                        if (row, col) in lights:
                            lights.remove((row, col))
                        else:
                            lights.add((row, col))

    return len(lights)


def solution(filename: str) -> int:
    instructions: List[Instruction] = parse(filename)
    return solve(instructions)


if __name__ == "__main__":
    print(solution("./input.txt"))  # 377891
