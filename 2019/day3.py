def get_moves():
    moves = []
    with open('./inputs/day3.txt') as f:
        for line in f.readlines():
            moves.append(line.rstrip('\n').split(','))
    return moves


def get_points(moves):
    points = [(0, 0)]
    for move in moves:
        dir = move[0]
        match dir:
            case 'U':
                point = (points[-1][0], points[-1][1] + int(move[1:]))
            case 'D':
                point = (points[-1][0], points[-1][1] - int(move[1:]))
            case 'L':
                point = (points[-1][0] - int(move[1:]), points[-1][1])
            case 'R':
                point = (points[-1][0] + int(move[1:]), points[-1][1])
            case _:
                raise ValueError(
                    f'Invalid direction {dir} (expected U, R, D, or L)')
        points.append(point)

    return points


def get_lines(points):
    lines = []
    for i in range(len(points)-1):
        lines.append([points[i], points[i+1]])

    return lines


def get_intersections(lines_a, lines_b):
    # Loop through lines in lines_a.
    # For each one, determine if it is vertical or horizontal.
    # If it is vertical, loop through all horizontal lines in lines_b, and vice versa.
    intersections = []
    for line_a in lines_a:
        for line_b in lines_b:
            if is_lines_intersecting(line_a, line_b):
                if is_line_vertical(line_a):
                    intersection = (line_a[0][0], line_b[0][1])
                else:
                    intersection = (line_b[0][0], line_a[0][1])
                if intersection != (0, 0):
                    intersections.append(intersection)

    return intersections


def get_distances(points):
    return [get_distance(point) for point in points]


def get_distance(point):
    return abs(point[0]) + abs(point[1])


def get_distance_to_point(point, lines):
    distance = 0
    point_found = False
    for line in lines:
        point_a = line[0]
        point_b = line[1]
        current_point = point_a
        # Up
        if is_line_vertical(line) and point_b[1] > point_a[1]:
            for i in range(abs(point_b[1] - point_a[1])):
                distance += 1
                current_point = (current_point[0], current_point[1] + 1)
                if current_point == point:
                    point_found = True
                    break
            pass
        # Down
        elif is_line_vertical(line) and point_b[1] < point_a[1]:
            for i in range(abs(point_b[1] - point_a[1])):
                distance += 1
                current_point = (current_point[0], current_point[1] - 1)
                if current_point == point:
                    point_found = True
                    break
            pass
        # Left
        elif is_line_horizontal(line) and point_b[0] < point_a[0]:
            for i in range(abs(point_b[0] - point_a[0])):
                distance += 1
                current_point = (current_point[0] - 1, current_point[1])
                if current_point == point:
                    point_found = True
                    break
            pass
        # Right
        elif is_line_horizontal(line) and point_b[0] > point_a[0]:
            for i in range(abs(point_b[0] - point_a[0])):
                distance += 1
                current_point = (current_point[0] + 1, current_point[1])
                if current_point == point:
                    point_found = True
                    break
            pass

        if point_found:
            break

    return distance


def is_lines_intersecting(line_a, line_b):
    # Can intersect in one of two ways:
    # 1. Line A is horizontal and line B is vertical,
    #    and the Y value of line A is between the Y values of line B,
    #    and the X value of line B is between the X values of line A;
    # 2. Line A is vertical and line B is horizontal,
    #    and the Y value of line B is between the Y values of line A,
    #    and the X value of line A is between the X values of line B;
    if (is_line_horizontal(line_a) and is_line_vertical(line_b)
        and is_between(line_a[0][1], line_b[0][1], line_b[1][1])
            and is_between(line_b[0][0], line_a[0][0], line_a[1][0])):
        return True
    elif (is_line_vertical(line_a) and is_line_horizontal(line_b)
          and is_between(line_b[0][1], line_a[0][1], line_a[1][1])
          and is_between(line_a[0][0], line_b[0][0], line_b[1][0])):
        return True

    return False


def is_line_vertical(line):
    # Line is vertical if the x-coordinates of both points are the same
    if line[0][0] == line[1][0]:
        return True
    return False


def is_line_horizontal(line):
    if line[0][1] == line[1][1]:
        return True
    return False


def is_between(val, bound_1, bound_2):
    if (val <= bound_1 and val >= bound_2) or (val >= bound_1 and val <= bound_2):
        return True
    return False


def solve_part_one():
    moves = get_moves()
    points_a = get_points(moves[0])
    points_b = get_points(moves[1])
    lines_a = get_lines(points_a)
    lines_b = get_lines(points_b)
    intersections = get_intersections(lines_a, lines_b)
    distances = get_distances(intersections)
    return min(distances)


def solve_part_two():
    # Get intersections.
    moves = get_moves()
    points_a = get_points(moves[0])
    points_b = get_points(moves[1])
    lines_a = get_lines(points_a)
    lines_b = get_lines(points_b)
    intersections = get_intersections(lines_a, lines_b)

    # For each intersection, loop through lines to count how many steps it would
    # take to reach it for each
    dists = []
    for intersection in intersections:
        dist_a = get_distance_to_point(intersection, lines_a)
        dist_b = get_distance_to_point(intersection, lines_b)
        dists.append(dist_a+dist_b)

    return min(dists)


print(solve_part_one())
print(solve_part_two())
