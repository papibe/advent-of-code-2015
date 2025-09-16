from typing import List

I_VALUE: int = ord("i") - ord("a")
O_VALUE: int = ord("o") - ord("a")
L_VALUE: int = ord("l") - ord("a")


def convert(p: str) -> List[int]:
    return [ord(c) - ord("a") for c in p]


def un_convert(p: List[int]) -> List[str]:
    return [chr(ord("a") + n) for n in p]


def first_rule(p: List[int]) -> bool:
    """
    Passwords must include one increasing straight of at least three letters, like abc,
    bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count
    """
    # looking for a increasing sequence
    for index in range(len(p) - 3 + 1):
        if p[index] == p[index + 1] - 1 and p[index + 1] == p[index + 2] - 1:
            return True

    return False


def second_rule(p: List[int]) -> bool:
    """
    Passwords may not contain the letters i, o, or l, as these letters can be mistaken
    for other characters and are therefore confusing
    """
    for value in p:
        if value in [I_VALUE, O_VALUE, L_VALUE]:
            return False

    return True


def third_rule(p: List[int]) -> bool:
    """
    Passwords must contain at least two different, non-overlapping pairs of letters,
    like aa, bb, or zz
    """
    for i in range(len(p) - 2 + 1):
        if p[i] == p[i + 1]:
            for j in range(i + 2, len(p) - 2 + 1):
                if p[j] == p[j + 1]:
                    return True

    return False


def is_valid(p: List[int]) -> bool:
    return first_rule(p) and second_rule(p) and third_rule(p)


def increment2(p: List[int]) -> List[int]:
    carry: int = 1

    for index in range(len(p) - 1, -1, -1):
        new_digit: int = p[index] + carry
        if new_digit in [I_VALUE, O_VALUE, L_VALUE]:
            new_digit += 1

        p[index] = new_digit % 26
        carry = new_digit // 26

        if carry == 0:
            break

    return p


def next_good(p: List[int], index: int) -> List[int]:
    p[index] += 1
    for i in range(index + 1, len(p)):
        p[i] = 0
    return p


def skip_to_next(p: List[int]) -> List[int]:
    for index in range(len(p)):
        if p[index] in [I_VALUE, O_VALUE, L_VALUE]:
            return next_good(p, index)

    return p


def increment(p: List[int]) -> List[int]:
    p = skip_to_next(p)

    index: int = len(p) - 1
    new_digit: int = p[index] + 1
    p[index] = new_digit % 26
    carry: int = new_digit // 26

    while index > 0 and carry != 0:
        index -= 1

        new_digit = p[index] + 1
        p[index] = new_digit % 26
        carry = new_digit // 26

    return p


def solution(str_password: str) -> str:
    password: List[int] = convert(str_password)

    while True:
        password = increment(password)
        if is_valid(password):
            return "".join(un_convert(password))

    return ""


if __name__ == "__main__":
    print(solution("hepxcrrq"))  # hepxxyzz
