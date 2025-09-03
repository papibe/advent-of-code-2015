import pytest

from part1 import measure


@pytest.mark.parametrize(
    "string,expected_length,expected_in_memory",
    [
        ('""', 2, 0),
        ('"abc"', 5, 3),
        (r'"aaa\"aaa"', 10, 7),
        (r'"\x27"', 6, 1),
        ('r"nywbv\\"', 9, 6),
    ],
)
def test_part1_measure(
    string: str, expected_length: int, expected_in_memory: int
) -> None:
    length, in_memory = measure(string)
    assert length == expected_length, f"got {length}, needs {expected_length}"
    assert (
        in_memory == expected_in_memory
    ), f"got {in_memory}, needs {expected_in_memory}"
