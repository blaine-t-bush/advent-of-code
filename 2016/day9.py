def read_input():
    with open("./inputs/day9.txt") as f:
        input = f.readlines()
    return input[0]


def find_next_marker(string: str, start: int):
    index_start = string.find("(", start)
    if index_start == -1:
        return None, None
    index_end = string.find(")", index_start + 1)
    if index_end == -1:
        return None, None
    return index_start, index_end


def parse_marker(string: str, start: int, end: int):
    marker_split = string[start + 1 : end].split("x")
    return int(marker_split[0]), int(marker_split[1])


def decompress_next(string: str, start_index: int):
    next_start, next_end = find_next_marker(string, start_index)
    if next_start is None or next_end is None:
        return string, 0

    characters, repeats = parse_marker(string, next_start, next_end)
    if characters is None or repeats is None:
        return string, 0

    decompressed = (
        string[:next_start]
        + string[next_end + 1 : next_end + 1 + characters] * repeats
        + string[next_end + 1 + characters :]
    )
    return decompressed, next_start + characters * repeats


def decompress(string: str):
    decompressed, skip = decompress_next(string, 0)
    while decompressed != string:
        string = decompressed
        decompressed, skip = decompress_next(decompressed, skip)
    return decompressed


def decompress_recursive(string: str):
    decompressed, skip = decompress_next(string, 0)
    while decompressed != string:
        string = decompressed
        decompressed, skip = decompress_next(decompressed, 0)
    return decompressed


def get_recursively_decompressed_length(string: str) -> int:
    original = string
    length = 0
    # get next marker parameters
    # add length up to and including next marker
    # remove characters up to and including next marker
    return 0


def solve_part_one():
    input = read_input()
    decompressed = decompress(input)
    return len(decompressed)


def solve_part_two():
    input = read_input()
    decompressed = decompress_recursive(input)
    return len(decompressed)


print(f"Part 1: decompressed length {solve_part_one()}")
print(f"Part 2: decompressed length {solve_part_two()}")
