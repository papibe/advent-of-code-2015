from dataclasses import dataclass
from enum import Enum
from typing import Dict, List, Tuple


class InstructionType(Enum):
    HLF = "hlf"
    TPL = "tpl"
    INC = "inc"
    JUMP = "jmp"
    JIE = "jie"
    JIO = "jio"


@dataclass
class Instruction:
    operation: InstructionType
    param1: str
    param2: str


def parse(filename: str) -> Tuple[List[Instruction], Dict[str, int]]:
    with open(filename, "r") as fp:
        data: List[str] = fp.read().splitlines()

    instructions: List[Instruction] = []
    registers: Dict[str, int] = {}

    for line in data:
        split_line: List[str] = line.split(" ")
        instr: str = split_line[0]
        params: str = "".join(split_line[1:])

        instruction_type: InstructionType
        param1: str = ""
        param2: str = ""

        match instr:
            case InstructionType.HLF.value:
                instruction_type = InstructionType.HLF
                param1 = split_line[1]
                registers[param1] = 0

            case InstructionType.TPL.value:
                instruction_type = InstructionType.TPL
                param1 = split_line[1]
                registers[param1] = 0

            case InstructionType.INC.value:
                instruction_type = InstructionType.INC
                param1 = split_line[1]
                registers[param1] = 0

            case InstructionType.JUMP.value:
                instruction_type = InstructionType.JUMP
                param1 = split_line[1]

            case InstructionType.JIE.value:
                instruction_type = InstructionType.JIE
                split_params = params.split(",")
                param1 = split_params[0]
                param2 = split_params[1]
                registers[param1] = 0

            case InstructionType.JIO.value:
                instruction_type = InstructionType.JIO
                split_params = params.split(",")
                param1 = split_params[0]
                param2 = split_params[1]
                registers[param1] = 0

            case _:
                raise ValueError("Unknown instruction")

        instructions.append(Instruction(instruction_type, param1, param2))

    return instructions, registers


def solve(
    instructions: List[Instruction], registers: Dict[str, int], output_register: str
) -> int:

    pointer: int = 0

    while pointer < len(instructions):
        instruction: Instruction = instructions[pointer]
        instr: InstructionType = instruction.operation

        match instr:
            case InstructionType.HLF:
                registers[instruction.param1] //= 2

            case InstructionType.TPL:
                registers[instruction.param1] *= 3

            case InstructionType.INC:
                registers[instruction.param1] += 1

            case InstructionType.JUMP:
                pointer += int(instruction.param1)
                continue

            case InstructionType.JIE:
                if registers[instruction.param1] % 2 == 0:
                    pointer += int(instruction.param2)
                    continue

            case InstructionType.JIO:
                if registers[instruction.param1] == 1:
                    pointer += int(instruction.param2)
                    continue

        pointer += 1

    return registers[output_register]


def solution(filename: str, output_register: str) -> int:
    instructions, registers = parse(filename)
    return solve(instructions, registers, output_register)


if __name__ == "__main__":
    print(solution("./example.txt", "a"))  # 2
    print(solution("./input.txt", "b"))  # 255
