def is_length_4_unique_palindrome(string: str):
    if len(string) != 4:
        raise ValueError(f"Expected string of length 4, received {string}")
    return string[0] == string[3] and string[1] == string[2] and string[0] != string[1]


def is_length_3_unique_palindrome(string: str):
    if len(string) != 3:
        raise ValueError(f"Expected string of length 4, received {string}")
    return string[0] == string[2] and string[0] != string[1]


def has_abba(string: str):
    if len(string) < 4:
        return False
    elif len(string) == 4:
        return is_length_4_unique_palindrome(string)
    else:
        for start_index in range(0, len(string) - 3):
            substring = string[start_index : start_index + 4]
            if is_length_4_unique_palindrome(substring):
                return True

    return False


def has_aba(string: str):
    if len(string) == 3:
        return is_length_3_unique_palindrome(string), [string]
    elif len(string) > 3:
        abas = []
        for start_index in range(0, len(string) - 2):
            substring = string[start_index : start_index + 3]
            if is_length_3_unique_palindrome(substring):
                abas.append(substring)
        if len(abas) > 0:
            return True, abas

    return False, None


def supports_tls(input):
    hypernet_sequences, supernet_sequences = parse_input(input)
    for hypernet_sequence in hypernet_sequences:
        if has_abba(hypernet_sequence):
            return False
    for supernet_sequence in supernet_sequences:
        if has_abba(supernet_sequence):
            return True

    return False


def supports_ssl(input):
    hypernet_sequences, supernet_sequences = parse_input(input)
    for supernet_sequence in supernet_sequences:
        has, abas = has_aba(supernet_sequence)
        if has:
            for aba in abas:
                bab = aba[1] + aba[0] + aba[1]
                for hypernet_sequence in hypernet_sequences:
                    if bab in hypernet_sequence:
                        return True

    return False


def parse_input(input):
    # Find hypernet sequences (strings between square brackets).
    index_start = -1
    bracket_pairs = []
    while True:
        index_start = input.find("[", index_start + 1)
        if index_start == -1:
            break
        index_end = input.find("]", index_start + 1)

        bracket_pairs.append((index_start, index_end))

    hypernet_sequences = []
    for bracket_pair in bracket_pairs:
        hypernet_sequences.append(
            input[bracket_pair[0] + 1 : bracket_pair[1]].rstrip("\n")
        )

    # Find non-hypernet sequences
    supernet_sequences = []
    for index, bracket_pair in enumerate(bracket_pairs):
        if index == 0:
            supernet_sequences.append(input[0 : bracket_pair[0]].rstrip("\n"))
        else:
            supernet_sequences.append(
                input[bracket_pairs[index - 1][1] + 1 : bracket_pairs[index][0]].rstrip(
                    "\n"
                )
            )

        if index == len(bracket_pairs) - 1:
            supernet_sequences.append(input[bracket_pair[1] + 1 :].rstrip("\n"))

    return hypernet_sequences, supernet_sequences


def read_inputs():
    with open("./inputs/day7.txt") as f:
        inputs = f.readlines()

    return inputs


def solve_part_one():
    inputs = read_inputs()
    count = 0
    for input in inputs:
        if supports_tls(input):
            count += 1

    return count


def solve_part_two():
    inputs = read_inputs()
    count = 0
    for input in inputs:
        if supports_ssl(input):
            count += 1

    return count


# An IP supports TLS if it has an ABBA.
# The IP also must not have an ABBA within any hypernet sequences.

print(f"Part 1: {solve_part_one()} support(s) TLS")
print(f"Part 2: {solve_part_two()} support(s) TLS")
