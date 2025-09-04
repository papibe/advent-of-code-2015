import pytest

from part2 import encode


@pytest.mark.parametrize(
    "string,expected_length,expected_in_memory",
    [
        ('""', 2, 6),
        ('"abc"', 5, 9),
        (r'"aaa\"aaa"', 10, 16),
        (r'"\x27"', 6, 11),
        (r'"nywbv\\"', 9, 15),
    ],
)
def test_part1_measure(
    string: str, expected_length: int, expected_in_memory: int
) -> None:
    length, in_memory = encode(string)
    assert length == expected_length, f"got {length}, needs {expected_length}"
    assert (
        in_memory == expected_in_memory
    ), f"got {in_memory}, needs {expected_in_memory}"
