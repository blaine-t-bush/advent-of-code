import utils


def calc_smallest_side_surface_area(l: int, w: int, h: int):
    return min([l * w, w * h, h * l])


def calc_total_surface_area(l: int, w: int, h: int):
    return 2 * (l * w + w * h + h * l)


def calc_required_paper_area(l: int, w: int, h: int):
    return calc_total_surface_area(l, w, h) + calc_smallest_side_surface_area(l, w, h)


def parse_dimensions(line: str):
    return [int(dim) for dim in line.split("x")]


def calc_smallest_perimeter(l: int, w: int, h: int):
    return min([2 * (l + w), 2 * (w + h), 2 * (h + l)])


def calc_volume(l: int, w: int, h: int):
    return l * w * h


def calc_required_ribbon_length(l: int, w: int, h: int):
    return calc_smallest_perimeter(l, w, h) + calc_volume(l, w, h)


def solve_part_one():
    lines = utils.get_input_as_lines(2)
    required_area = 0
    for line in lines:
        dims = parse_dimensions(line)
        required_area += calc_required_paper_area(dims[0], dims[1], dims[2])
    return required_area


def solve_part_two():
    lines = utils.get_input_as_lines(2)
    required_len = 0
    for line in lines:
        dims = parse_dimensions(line)
        required_len += calc_required_ribbon_length(dims[0], dims[1], dims[2])
    return required_len


part_one_solution = solve_part_one()
print(f"Part 1: total required wrapping paper {part_one_solution} sq ft")
part_two_solution = solve_part_two()
print(f"Part 2: total required ribbon {part_two_solution} ft")
