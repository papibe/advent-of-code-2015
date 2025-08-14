import hashlib


def md5_hash(string: str) -> str:
    # Create an MD5 hash object
    hash_object = hashlib.md5(string.encode())

    # Get the hexadecimal representation of the hash
    hex_dig = hash_object.hexdigest()

    return hex_dig


def is_good(s: str, zeros: int) -> bool:
    for i in range(zeros):
        if s[i] != "0":
            return False
    return True


def solve(base: str, zeros: int) -> int:
    counter: int = 1

    hash_: str = md5_hash(base + str(counter))
    while not is_good(hash_, zeros):
        counter += 1
        hash_ = md5_hash(base + str(counter))

    return counter


if __name__ == "__main__":
    print("Part 1:", solve("iwrupvqb", 5))  # 346386
    print("Part 2:", solve("iwrupvqb", 6))  # 9958218
