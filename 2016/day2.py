class Keypad:
    def __init__(self, start: tuple, layout):
        self.coords = layout
        self.coord = start

    def x(self) -> int:
        return self.coord[0]

    def y(self) -> int:
        return self.coord[1]

    def set_coord(self, coord: tuple):
        self.coord = coord

    def move(self, dir: str):
        match dir:
            case "U":
                newCoord = (self.x(), self.y() - 1)
                if newCoord in self.coords:
                    self.set_coord(newCoord)
            case "R":
                newCoord = (self.x() + 1, self.y())
                if newCoord in self.coords:
                    self.set_coord(newCoord)
            case "D":
                newCoord = (self.x(), self.y() + 1)
                if newCoord in self.coords:
                    self.set_coord(newCoord)
            case "L":
                newCoord = (self.x() - 1, self.y())
                if newCoord in self.coords:
                    self.set_coord(newCoord)
            case _:
                raise ValueError(f"Expected dir U, R, D, or L, got {dir}")


def get_code(keypad: Keypad, inputs: list):
    # Simulate movements
    keypresses = []
    for line in inputs:
        for direction in line:
            keypad.move(direction)
        keypresses.append(keypad.coord)

    return "".join([keypad.coords[c] for c in keypresses])


def solve_part_one():
    # Initialize keypad
    # 0 1 2
    # =====
    # 1 2 3
    # 4 5 6
    # 7 8 9
    keypad = Keypad(
        (1, 1),
        {
            (0, 0): "1",
            (1, 0): "2",
            (2, 0): "3",
            (0, 1): "4",
            (1, 1): "5",
            (2, 1): "6",
            (0, 2): "7",
            (1, 2): "8",
            (2, 2): "9",
        },
    )

    # Get inputs
    inputs = []
    with open("./inputs/day2.txt") as f:
        for line in f.readlines():
            inputs.append(line.rstrip("\n"))

    return get_code(keypad, inputs)


def solve_part_two():
    # 0 1 2 3 4
    # =========
    #     1
    #   2 3 4
    # 5 6 7 8 9
    #   A B C
    #     D
    keypad = Keypad(
        (0, 2),
        {
            (2, 0): "1",
            (1, 1): "2",
            (2, 1): "3",
            (3, 1): "4",
            (0, 2): "5",
            (1, 2): "6",
            (2, 2): "7",
            (3, 2): "8",
            (4, 2): "9",
            (1, 3): "A",
            (2, 3): "B",
            (3, 3): "C",
            (2, 4): "D",
        },
    )

    # Get inputs
    inputs = []
    with open("./inputs/day2.txt") as f:
        for line in f.readlines():
            inputs.append(line.rstrip("\n"))

    return get_code(keypad, inputs)


print(f"Part 1: keypresses {solve_part_one()}")
print(f"Part 2: keypresses {solve_part_two()}")
