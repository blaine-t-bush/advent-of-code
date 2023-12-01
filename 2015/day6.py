import re
import utils
from typing import List


class Instruction:
    def __init__(self, line: str):
        result = re.search(
            r"(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)", line
        )

        self.command = result[1]
        self.startCoord = (int(result[2]), int(result[3]))
        self.endCoord = (int(result[4]), int(result[5]))


def parse_instructions(lines: List[str]) -> List[Instruction]:
    return [Instruction(line) for line in lines]


def run_instructions(instructions: List[Instruction]) -> List[tuple]:
    coords = []
    for instruction in instructions:
        print(instruction)
        for x in range(instruction.startCoord[0], instruction.endCoord[0] + 1):
            for y in range(instruction.startCoord[1], instruction.endCoord[1] + 1):
                coord = (x, y)
                if coord in coords:
                    if (
                        instruction.command == "toggle"
                        or instruction.command == "turn off"
                    ):
                        coords.remove(coord)
                else:
                    if (
                        instruction.command == "toggle"
                        or instruction.command == "turn on"
                    ):
                        coords.append(coord)
    return coords


def solve_part_one():
    lines = utils.get_input_as_lines(6)
    instructions = parse_instructions(lines)
    on = run_instructions(instructions)
    return len(on)


print(f"Part 1: {solve_part_one()} lit lights")
