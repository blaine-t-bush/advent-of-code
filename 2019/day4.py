def get_range():
    with open('./inputs/day4.txt') as f:
        for line in f.readlines():
            splitted = line.split('-')
            range_min = int(splitted[0])
            range_max = int(splitted[1])

    return range_min, range_max


def obeys_rule_one(password):
    return len(password) == 6 and password.isdigit()


def obeys_rule_two(password, range_min, range_max):
    return int(password) >= range_min and int(password) <= range_max


def obeys_rule_three(password):
    for i in range(1, len(password)):
        if password[i] == password[i-1]:
            return True
    return False


def contains_doubles(password):
    if password[0] == password[1] and password[1] != password[2]:
        return True
    for i in range(1, 4):
        if password[i-1] != password[i] and password[i] == password[i+1] and password[i+1] != password[i+2]:
            return True
    if password[3] != password[4] and password[4] == password[5]:
        return True
    return False


def contains_no_triples(password):
    for i in range(0, 4):
        if password[i] == password[i+1] and password[i+1] == password[i+2]:
            return False
    return True


def obeys_rule_four(password):
    for i in range(1, len(password)):
        if int(password[i]) < int(password[i-1]):
            return False
    return True


def obeys_rules_part_one(password, range_min, range_max):
    return (obeys_rule_one(password)
            and obeys_rule_two(password, range_min, range_max)
            and obeys_rule_three(password)
            and obeys_rule_four(password))


def obeys_rules_part_two(password, range_min, range_max):
    return (obeys_rule_one(password)
            and obeys_rule_two(password, range_min, range_max)
            and contains_doubles(password)
            and obeys_rule_four(password))


def solve_part_one():
    range_min, range_max = get_range()
    count = 0
    for num in range(range_min, range_max+1):
        if obeys_rules_part_one(str(num), range_min, range_max):
            count += 1

    return count


def solve_part_two():
    range_min, range_max = get_range()
    count = 0
    for num in range(range_min, range_max+1):
        if obeys_rules_part_two(str(num), range_min, range_max):
            count += 1

    return count


print(f'Part 1: {solve_part_one()}')
print(f'Part 2: {solve_part_two()}')
