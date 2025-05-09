import pytest

from part1 import (
    convert,
    first_rule,
    second_rule,
    solution,
    third_rule,
)


@pytest.mark.parametrize(
    "password,expected",
    [
        ("hijklmmn", True),
        ("abbceffg", False),
        ("abbcegjk", False),
    ],
    # ids=[
    # ],
)
def test_pass_first_rule(password: str, expected: bool) -> None:
    result: int = first_rule(convert(password))
    assert result == expected, f"got {result}, needs {expected}"


@pytest.mark.parametrize(
    "password,expected",
    [
        ("hijklmmn", False),
        ("abbceffg", True),
        ("abbcegjk", True),
    ],
    # ids=[
    # ],
)
def test_pass_second_rule(password: str, expected: bool) -> None:
    result: int = second_rule(convert(password))
    assert result == expected, f"got {result}, needs {expected}"


@pytest.mark.parametrize(
    "password,expected",
    [
        ("hijklmmn", False),
        ("abbceffg", True),
        ("abbcegjk", False),
    ],
    # ids=[
    # ],
)
def test_pass_third_rule(password: str, expected: bool) -> None:
    result: int = third_rule(convert(password))
    assert result == expected, f"got {result}, needs {expected}"


@pytest.mark.parametrize(
    "password,expected",
    [
        # ("abcdefgh", "abcdffaa"),
        ("ghijklmn", "ghjaabcc"),
    ],
    # ids=[
    # ],
)
def test_next_password(password: str, expected: str) -> None:
    result: str = solution(password)
    assert result == expected, f"got {result}, needs {expected}"
