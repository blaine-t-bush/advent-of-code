def has_valid_side_lengths(a: int, b: int, c: int) -> bool:
    all_sides = [a, b, c]
    longest_side = max(all_sides)
    shorter_sides = all_sides
    shorter_sides.pop(all_sides.index(longest_side))
    return sum(shorter_sides) > longest_side


def get_lengths():
    lengths = []
    with open("./inputs/day3.txt") as f:
        for line in f.readlines():
            length = []
            split = line.split(" ")
            while "" in split:
                split.remove("")
            for string in split:
                length.append(int(string.rstrip("\n")))
            lengths.append(length)
    return lengths


def get_lengths_by_column():
    # Get by row
    lengths_by_row = get_lengths()

    # Transform
    lengths = []
    for index in range(0, len(lengths_by_row), 3):
        lengths.append(
            [
                lengths_by_row[index][0],
                lengths_by_row[index + 1][0],
                lengths_by_row[index + 2][0],
            ]
        )
        lengths.append(
            [
                lengths_by_row[index][1],
                lengths_by_row[index + 1][1],
                lengths_by_row[index + 2][1],
            ]
        )
        lengths.append(
            [
                lengths_by_row[index][2],
                lengths_by_row[index + 1][2],
                lengths_by_row[index + 2][2],
            ]
        )

    return lengths


def solve_part_one():
    lengths = get_lengths()
    valid_count = 0
    for sides in lengths:
        if has_valid_side_lengths(sides[0], sides[1], sides[2]):
            valid_count += 1
    return valid_count


def solve_part_two():
    lengths = get_lengths_by_column()
    valid_count = 0
    for sides in lengths:
        if has_valid_side_lengths(sides[0], sides[1], sides[2]):
            valid_count += 1
    return valid_count


print(f"Part 1: {solve_part_one()} valid triangles")
print(f"Part 2: {solve_part_two()} valid triangles")
