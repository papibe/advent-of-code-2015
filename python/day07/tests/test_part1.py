from typing import Dict

from part1 import parse, solve

expected: Dict[str, int] = {
    "d": 72,
    "e": 507,
    "f": 492,
    "g": 114,
    "h": 65412,
    "i": 65079,
    "x": 123,
    "y": 456,
}


def test_part1_example_data() -> None:
    wires, instructions = parse("./example.txt")
    wire_values: Dict[str, int] = solve(wires, instructions)

    for wire, value in expected.items():
        assert wire_values[wire] == value
