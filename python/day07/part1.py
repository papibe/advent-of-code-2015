import re
from collections import deque
from dataclasses import dataclass
from enum import Enum
from typing import Deque, Dict, List, Match, Optional, Tuple


class Operation(Enum):
    ASSIGNMENT = "ASSIGNMENT"
    AND = "AND"
    OR = "OR"
    LSHIFT = "LSHIFT"
    RSHIFT = "RSHIFT"
    NOT = "NOT"


@dataclass
class Instruction:
    operation: Operation
    param1: str
    param2: str
    wire: str
    depends: List[str]


def parse(filename: str) -> Tuple[Dict[str, Instruction], List[Instruction]]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    wires: Dict[str, Instruction] = {}
    instructions: List[Instruction] = []

    re_assignment: str = r"(\w+) -> (\w+)"
    re_operation: str = r"(\w+) (\w+) (\w+) -> (\w+)"
    re_not: str = r"NOT (\w+) -> (\w+)"

    for line in data:
        assignment_match: Optional[Match[str]] = re.match(re_assignment, line)
        operation_match: Optional[Match[str]] = re.match(re_operation, line)
        not_match: Optional[Match[str]] = re.match(re_not, line)

        assert assignment_match or operation_match or not_match, print(line)

        operation: Operation
        param1: str = ""
        param2: str = ""
        wire: str = ""

        if not_match:
            operation = Operation.NOT
            param1 = not_match.group(1)
            wire = not_match.group(2)

        if operation_match:
            param1 = operation_match.group(1)
            str_operation: str = operation_match.group(2)
            param2 = operation_match.group(3)
            wire = operation_match.group(4)
            match str_operation:
                case "AND":
                    operation = Operation.AND
                case "OR":
                    operation = Operation.OR
                case "LSHIFT":
                    operation = Operation.LSHIFT
                case "RSHIFT":
                    operation = Operation.RSHIFT
                case _:
                    raise ValueError(f"{str_operation} not known")

        if assignment_match:
            operation = Operation.ASSIGNMENT
            param1 = assignment_match.group(1)
            wire = assignment_match.group(2)

        instruction = Instruction(operation, param1, param2, wire, [])
        wires[wire] = instruction

        instructions.append(instruction)

    return wires, instructions


def build_dependencies(
    wires: Dict[str, Instruction], instructions: List[Instruction]
) -> Dict[str, List[str]]:
    wire_dependencies: Dict[str, List[str]] = {}
    for wire in wires:
        wire_dependencies[wire] = []

    # update dependencies
    for instruction in instructions:
        match instruction.operation:
            case Operation.NOT:
                if instruction.param1 in wires:
                    instruction.depends.append(instruction.param1)
                    if instruction.wire not in wire_dependencies:
                        wire_dependencies[instruction.wire] = []

                    wire_dependencies[instruction.wire].append(instruction.param1)

            case Operation.ASSIGNMENT:
                if instruction.param1 in wires:
                    instruction.depends.append(instruction.param1)
                    if instruction.wire not in wire_dependencies:
                        wire_dependencies[instruction.wire] = []

                    wire_dependencies[instruction.wire].append(instruction.param1)

            case Operation.AND | Operation.OR | Operation.LSHIFT | Operation.RSHIFT:
                if instruction.param1 in wires:
                    instruction.depends.append(instruction.param1)
                    if instruction.wire not in wire_dependencies:
                        wire_dependencies[instruction.wire] = []

                    wire_dependencies[instruction.wire].append(instruction.param1)

                if instruction.param2 in wires:
                    instruction.depends.append(instruction.param2)
                    if instruction.wire not in wire_dependencies:
                        wire_dependencies[instruction.wire] = []

                    wire_dependencies[instruction.wire].append(instruction.param2)

            case _:
                raise ValueError("Unknown instruction")

    return wire_dependencies


def run_instruction(
    wire: str, wires: Dict[str, Instruction], wires_values: Dict[str, int]
) -> None:
    instruction = wires[wire]
    match instruction.operation:
        case Operation.NOT:
            if instruction.param1 in wires:
                value = wires_values[instruction.param1]
            else:
                value = int(instruction.param1)

            wires_values[wire] = value ^ 0xFFFF
            assert 0 <= wires_values[wire] < 65535

        case Operation.ASSIGNMENT:
            if instruction.param1 in wires:
                value = wires_values[instruction.param1]
            else:
                value = int(instruction.param1)

            wires_values[wire] = value

        case Operation.AND:
            if instruction.param1 in wires:
                value1 = wires_values[instruction.param1]
            else:
                value1 = int(instruction.param1)

            if instruction.param2 in wires:
                value2 = wires_values[instruction.param2]
            else:
                value2 = int(instruction.param2)

            wires_values[wire] = value1 & value2

        case Operation.OR:
            if instruction.param1 in wires:
                value1 = wires_values[instruction.param1]
            else:
                value1 = int(instruction.param1)

            if instruction.param2 in wires:
                value2 = wires_values[instruction.param2]
            else:
                value2 = int(instruction.param2)

            wires_values[wire] = value1 | value2

        case Operation.LSHIFT:
            if instruction.param1 in wires:
                value1 = wires_values[instruction.param1]
            else:
                value1 = int(instruction.param1)

            if instruction.param2 in wires:
                value2 = wires_values[instruction.param2]
            else:
                value2 = int(instruction.param2)

            wires_values[wire] = value1 << value2

            assert 0 <= wires_values[wire] < 65535

        case Operation.RSHIFT:
            if instruction.param1 in wires:
                value1 = wires_values[instruction.param1]
            else:
                value1 = int(instruction.param1)

            if instruction.param2 in wires:
                value2 = wires_values[instruction.param2]
            else:
                value2 = int(instruction.param2)

            wires_values[wire] = value1 >> value2

            assert 0 <= wires_values[wire] < 65535

        case _:
            raise ValueError("WTF")


def kahn_top_sort(wire_dependencies: Dict[str, List[str]]) -> List[str]:
    in_degree: Dict[str, int] = {}
    wire: str
    for wire in wire_dependencies:
        in_degree[wire] = 0

    for _wire, dependencies in wire_dependencies.items():
        for neighbor_wire in dependencies:
            in_degree[neighbor_wire] += 1

    queue: Deque[str] = deque([wire for wire, deps in in_degree.items() if deps == 0])

    ordered_output: List[str] = []

    while queue:
        wire = queue.popleft()
        ordered_output.append(wire)

        for neighbor_wire in wire_dependencies[wire]:
            in_degree[neighbor_wire] -= 1
            if in_degree[neighbor_wire] == 0:
                queue.append(neighbor_wire)

    return ordered_output


def solve(
    wires: Dict[str, Instruction],
    instructions: List[Instruction],
) -> Dict[str, int]:
    wire_dependencies: Dict[str, List[str]] = build_dependencies(wires, instructions)
    sorted_wires: List[str] = kahn_top_sort(wire_dependencies)

    wires_values: Dict[str, int] = {}

    for index in range(len(sorted_wires) - 1, -1, -1):
        run_instruction(sorted_wires[index], wires, wires_values)

    return wires_values


def solution(filename: str) -> int:
    wires, instructions = parse(filename)
    wire_values: Dict[str, int] = solve(wires, instructions)

    return wire_values["a"]


if __name__ == "__main__":
    print(solution("./input.txt"))  # 16076
