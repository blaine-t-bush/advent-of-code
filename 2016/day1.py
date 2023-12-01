def get_input():
    with open("./inputs/day1.txt") as f:
        for line in f.readlines():
            list = line.split(", ")
    return list


def rotate(current_facing, clockwise=True):
    directions = ["N", "E", "S", "W"]
    current_index = directions.index(current_facing)
    if clockwise:
        next_index = current_index + 1
    else:
        next_index = current_index - 1
    return directions[next_index % 4]


def get_distance(start: tuple, end: tuple) -> int:
    return abs(end[0] - start[0]) + abs(end[1] - start[1])


def get_next_facing(current_facing: str, step):
    # Get new facing
    if step[0] == "R":
        clockwise = True
    else:
        clockwise = False

    return rotate(current_facing, clockwise)


def get_next_coord(facing: str, current_coord: tuple, distance: int):
    # Get new coord
    match facing:
        case "N":
            coord = (current_coord[0], current_coord[1] + distance)
        case "E":
            coord = (current_coord[0] + distance, current_coord[1])
        case "S":
            coord = (current_coord[0], current_coord[1] - distance)
        case "W":
            coord = (current_coord[0] - distance, current_coord[1])

    return coord


def solve_part_one():
    current_coord = (0, 0)
    current_facing = "N"
    steps = get_input()
    for step in steps:
        distance = int(step[1:])
        current_facing = get_next_facing(current_facing, step)
        current_coord = get_next_coord(current_facing, current_coord, distance)

    return get_distance((0, 0), current_coord)


def solve_part_two():
    current_coord = (0, 0)
    current_facing = "N"
    visited = [current_coord]
    steps = get_input()
    for step in steps:
        distance = int(step[1:])
        current_facing = get_next_facing(current_facing, step)
        for i in range(distance):
            current_coord = get_next_coord(current_facing, current_coord, 1)
            if current_coord in visited:
                return get_distance((0, 0), current_coord)
            else:
                visited.append(current_coord)

    return None


print(f"Part 1: distance {solve_part_one()}")
print(f"Part 2: distance {solve_part_two()}")
