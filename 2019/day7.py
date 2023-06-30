from intcode import get_program, run_program
from itertools import permutations


def run_amps(codes, sequence):
    # Pass first element of input sequence and 0 to amp A, and run program.
    codes_A, output_A = run_program(codes, [sequence[0], 0])

    # Pass second element of input sequence and amp A output to amp B, and run program. Repeat.
    codes_B, output_B = run_program(codes, [sequence[1], output_A])
    codes_C, output_C = run_program(codes, [sequence[2], output_B])
    codes_D, output_D = run_program(codes, [sequence[3], output_C])
    codes_E, output_E = run_program(codes, [sequence[4], output_D])

    return output_E


def run_amps_feedback_mode(codes, sequence):
    # Run program once.
    codes_A, output_A = run_program(codes, [sequence[0], 0], feedback_mode=True)
    codes_B, output_B = run_program(codes, [sequence[1], output_A], feedback_mode=True)
    codes_C, output_C = run_program(codes, [sequence[2], output_B], feedback_mode=True)
    codes_D, output_D = run_program(codes, [sequence[3], output_C], feedback_mode=True)
    codes_E, output_E = run_program(codes, [sequence[4], output_D], feedback_mode=True)

    return [codes_A, codes_B, codes_C, codes_D, codes_E], [
        output_A,
        output_B,
        output_C,
        output_D,
        output_E,
    ]


def solve_part_one():
    # Read program.
    codes = get_program("./inputs/day7.txt")

    # Determine all possible input sequences.
    sequences = list(permutations([0, 1, 2, 3, 4]))

    # Run amps for all possible sequences.
    max_output = 0
    for sequence in sequences:
        output = run_amps(codes, list(sequence))
        if output > max_output:
            max_output = output

    return max_output


def solve_part_two():
    # Read program.
    codes = get_program("./inputs/day7.txt")

    # Determine all possible input sequences.
    sequences = list(permutations([5, 6, 7, 8, 9]))

    # Run amps for all possible sequences.
    max_output = 0
    for sequence in sequences:
        codes_list, outputs_list = run_amps_feedback_mode(codes, list(sequence))
        if output > max_output:
            max_output = output

    return max_output


print(solve_part_one())
