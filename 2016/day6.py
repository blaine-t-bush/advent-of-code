def get_most_common_character(strings, index: int):
    occurrences = {}
    for string in strings:
        char = string[index]
        if char in occurrences:
            occurrences[char] = occurrences[char] + 1
        else:
            occurrences[char] = 1

    return max(occurrences, key=occurrences.get)


def get_least_common_character(strings, index: int):
    occurrences = {}
    for string in strings:
        char = string[index]
        if char in occurrences:
            occurrences[char] = occurrences[char] + 1
        else:
            occurrences[char] = 1

    return min(occurrences, key=occurrences.get)


def solve_part_one():
    with open("./inputs/day6.txt") as f:
        strings = [s.rstrip("\n") for s in f.readlines()]

    code = ""
    for i in range(len(strings[0])):
        code += get_most_common_character(strings, i)

    return code


def solve_part_two():
    with open("./inputs/day6.txt") as f:
        strings = [s.rstrip("\n") for s in f.readlines()]

    code = ""
    for i in range(len(strings[0])):
        code += get_least_common_character(strings, i)

    return code


print(f"Part 1: message '{solve_part_one()}'")
print(f"Part 2: message '{solve_part_two()}'")
