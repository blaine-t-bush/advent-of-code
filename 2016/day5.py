import hashlib


def starts_with_characters(string: str, char: str, count: int) -> bool:
    return len(string) >= count and string[0:count] == char * count


def solve_part_one():
    input = "ugkcyxxp"
    index = 0
    password = ""
    while True:
        hash = hashlib.md5(f"{input}{index}".encode()).hexdigest()
        if starts_with_characters(hash, "0", 5):
            password += hash[5]
            print(password + "_" * (8 - len(password)))
        if len(password) == 8:
            break
        index += 1

    return password


def solve_part_two():
    input = "ugkcyxxp"
    index = 0
    password = ["_", "_", "_", "_", "_", "_", "_", "_"]
    while True:
        hash = hashlib.md5(f"{input}{index}".encode()).hexdigest()
        if (
            starts_with_characters(hash, "0", 5)
            and hash[5].isdigit()
            and int(hash[5]) >= 0
            and int(hash[5]) <= 7
            and password[int(hash[5])] == "_"
        ):
            password[int(hash[5])] = hash[6]
            print("".join(password))
        if password.count("_") == 0:
            break
        index += 1

    return "".join(password)


print("Part 1:")
solve_part_one()
print("Part 2:")
solve_part_two()
