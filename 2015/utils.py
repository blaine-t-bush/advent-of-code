def get_input_as_lines(day: int):
    lines = []
    with open(f"./inputs/day{day}.txt") as f:
        for line in f.readlines():
            lines.append(line)
    return lines
