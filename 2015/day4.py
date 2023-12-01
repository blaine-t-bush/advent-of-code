import hashlib
import utils


def create_md5_hash(string: str):
    return hashlib.md5(string.encode("utf-8")).hexdigest()


def has_5_leading_zeros(string: str):
    if len(string) < 5 or string[0:5] != "00000":
        return False
    return True


def has_6_leading_zeros(string: str):
    if len(string) < 6 or string[0:6] != "000000":
        return False
    return True


def solve_part_one():
    input = utils.get_input_as_lines(4)[0]
    number = 1
    while True:
        hash = create_md5_hash(input + str(number))
        if has_5_leading_zeros(hash):
            return number
        number += 1


def solve_part_two():
    input = utils.get_input_as_lines(4)[0]
    number = 1
    while True:
        hash = create_md5_hash(input + str(number))
        if has_6_leading_zeros(hash):
            return number
        number += 1


print(f"Part 1: lowest number {solve_part_one()}")
print(f"Part 2: lowest number {solve_part_two()}")
