char_off = "."
char_on = "#"


class Screen:
    def __init__(self, width: int, height: int):
        self.width = width
        self.height = height

        self.pixels = []
        for y in range(self.height):
            self.pixels.append(char_off * self.width)

    def draw_rect(self, width: int, height: int):
        for y in range(height):
            self.pixels[y] = char_on * width + self.pixels[y][width:]

    def print(self):
        for string in self.pixels:
            print(string)

    def rotate_row(self, row_index: int, shift: int):
        shift = shift % self.width
        if shift != 0:
            self.pixels[row_index] = (
                self.pixels[row_index][-shift:]
                + self.pixels[row_index][: len(self.pixels[row_index]) - shift]
            )

    def rotate_col(self, col_index: int, shift: int):
        shift = shift % self.height
        if shift != 0:
            current_col = ""
            for y in range(self.height):
                current_col += self.pixels[y][col_index]
            new_col = current_col[-shift:] + current_col[: len(current_col) - shift]
            for y in range(self.height):
                self.pixels[y] = (
                    self.pixels[y][:col_index]
                    + new_col[y]
                    + self.pixels[y][col_index + 1 :]
                )

    def count_lit(self):
        count = 0
        for row in self.pixels:
            for char in row:
                if char == char_on:
                    count += 1

        return count


def solve_part_one():
    # Read inputs.
    with open("./inputs/day8.txt") as f:
        inputs = f.readlines()

    # Run commands one by one.
    screen = Screen(50, 6)
    for input in inputs:
        input = input.rstrip("\n")
        if input[0:4] == "rect":
            split = input[5:].split("x")
            print(f"rect {split[0]} by {split[1]}")
            screen.draw_rect(int(split[0]), int(split[1]))
        elif input[0:10] == "rotate row":
            split = input[13:].split(" by ")
            print(f"rotate row={split[0]} by {split[1]}")
            screen.rotate_row(int(split[0]), int(split[1]))
        elif input[0:13] == "rotate column":
            split = input[16:].split(" by ")
            print(f"rotate col={split[0]} by {split[1]}")
            screen.rotate_col(int(split[0]), int(split[1]))

    return screen


print("Part 1:")
screen = solve_part_one()
screen.print()
print(f"{screen.count_lit()} lit pixels")
