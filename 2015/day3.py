import utils


def get_visited_coords(instructions):
    coords = [(0, 0)]
    for instruction in instructions:
        match instruction:
            case ">":
                coords.append((coords[-1][0] + 1, coords[-1][1]))
            case "<":
                coords.append((coords[-1][0] - 1, coords[-1][1]))
            case "^":
                coords.append((coords[-1][0], coords[-1][1] + 1))
            case "v":
                coords.append((coords[-1][0], coords[-1][1] - 1))
            case _:
                raise ValueError(
                    f'Expected instruction ">", "<", "^", or "v", received {instruction}'
                )
    return coords


def get_count_by_coord(coords):
    counts = {}
    for coord in coords:
        if coord in counts:
            counts[coord] += 1
        else:
            counts[coord] = 1
    return counts


def solve_part_one():
    instructions = utils.get_input_as_lines(3)[0]
    coords = get_visited_coords(instructions)
    counts = get_count_by_coord(coords)
    return len(counts)


def solve_part_two():
    instructions = utils.get_input_as_lines(3)[0]
    instructions_1 = ""
    instructions_2 = ""
    for index, instruction in enumerate(instructions):
        if index % 2 == 0:
            instructions_1 += instruction
        else:
            instructions_2 += instruction

    coords = get_visited_coords(instructions_1) + get_visited_coords(instructions_2)
    counts = get_count_by_coord(coords)
    return len(counts)


print(f"Part 1: {solve_part_one()} houses")
print(f"Part 2: {solve_part_two()} houses")
