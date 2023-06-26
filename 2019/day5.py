from intcode import *


def solve_part_one():
    codes = get_program('./inputs/day5.txt')
    codes, output = run_program(codes)
    print(codes)
    print(output)


solve_part_one()
