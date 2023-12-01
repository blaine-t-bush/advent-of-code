import utils


def solve_part_one():
    input = utils.get_input_as_lines(1)[0]
    floor = 0
    for instruction in input:
        match instruction:
            case "(":
                floor += 1
            case ")":
                floor -= 1
            case _:
                raise ValueError(
                    f'Expected instruction "(" or ")", received {instruction}'
                )
    return floor


def solve_part_two():
    input = utils.get_input_as_lines(1)[0]
    floor = 0
    for index, instruction in enumerate(input):
        match instruction:
            case "(":
                floor += 1
            case ")":
                floor -= 1
            case _:
                raise ValueError(
                    f'Expected instruction "(" or ")", received {instruction}'
                )

        if floor == -1:
            return index + 1

    return None


part_one_solution = solve_part_one()
print(f"Part 1: final floor {part_one_solution}")
part_two_solution = solve_part_two()
print(f"Part 2: first basement achieved at index {part_two_solution}")
