import re


def rotate_alphabet(string: str, steps: int) -> str:
    alphabet = "abcdefghijklmnopqrstuvwxyz"
    steps = steps % len(alphabet)
    new = ""
    for char in string:
        if char == "-":
            new += " "
        else:
            new += alphabet[(alphabet.index(char) + steps) % len(alphabet)]
    return new


def sort_by_occurrence_then_alphabetically(string: str) -> str:
    # Get characters by number of occurrences
    occurrences = {}
    for char in string:
        if char in occurrences:
            occurrences[char] = occurrences[char] + 1
        else:
            occurrences[char] = 1

    ordered = ""
    # Start with the highest count and work down
    for test_count in range(max(occurrences.values()), 0, -1):
        letters_with_count = ""
        for letter, count in occurrences.items():
            if count == test_count:
                letters_with_count += letter
        ordered += "".join(sorted(letters_with_count))

    return ordered


def parse_room(room: str):
    res = re.search("(.*)-(\d+)\[(\w+)\]", room)
    return res[1], res[2], res[3]


def is_valid_room(name: str, checksum: str) -> bool:
    return (
        checksum == sort_by_occurrence_then_alphabetically(name.replace("-", ""))[0:5]
    )


def get_valid_rooms():
    with open("./inputs/day4.txt") as f:
        rooms = f.readlines()

    valid_rooms = []
    for room in rooms:
        name, id, checksum = parse_room(room)
        if is_valid_room(name, checksum):
            valid_rooms.append([name, int(id), checksum])

    return valid_rooms


def solve_part_one():
    valid_rooms = get_valid_rooms()
    return sum([room[1] for room in valid_rooms])


def solve_part_two():
    valid_rooms = get_valid_rooms()
    decryped_rooms = [
        [rotate_alphabet(room[0], room[1]), room[1], room[2]] for room in valid_rooms
    ]
    return [room[1] for room in decryped_rooms if "north" in room[0]]


print(f"Part 1: sum {solve_part_one()}")
print(f"Part 2: decrypted {solve_part_two()}")
