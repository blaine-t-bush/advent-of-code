import utils


def contains_vowels_at_least(string: str, count: int):
    vowel_count = len("".join(c for c in string if c in "aeiou"))
    if vowel_count >= count:
        return True
    return False


def contains_double(string: str):
    for i in range(len(string) - 1):
        if string[i] == string[i + 1]:
            return True
    return False


def contains_special_sequences(string: str):
    special_sequences = ["ab", "cd", "pq", "xy"]
    for special_sequence in special_sequences:
        if special_sequence in string:
            return True
    return False


def contains_sandwich(string: str):
    for i in range(len(string) - 2):
        if string[i] == string[i + 2]:
            return True
    return False


def contains_double_pair(string: str):
    for i in range(len(string) - 4):
        for j in range(i + 2, len(string) - 1):
            if string[i : i + 2] == string[j : j + 2]:
                return True

    return False


def solve_part_one():
    inputs = utils.get_input_as_lines(5)
    count = 0
    for input in inputs:
        if (
            contains_vowels_at_least(input, 3)
            and contains_double(input)
            and not (contains_special_sequences(input))
        ):
            count += 1
    return count


def solve_part_two():
    inputs = utils.get_input_as_lines(5)
    count = 0
    for input in inputs:
        if contains_sandwich(input) and contains_double_pair(input):
            count += 1
    return count


print(f"Part 1: {solve_part_one()} nice strings")
print(f"Part 2: {solve_part_two()} nice strings")
