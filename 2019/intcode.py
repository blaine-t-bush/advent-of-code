valid_position_mode_opcodes = [1, 2, 3, 4, 99]


def get_program(filename):
    codes = []
    with open(filename) as f:
        for line in f.readlines():
            for value in line.split(","):
                codes.append(int(value))

    return codes


def get_opcode_mode(opcode):
    if opcode in valid_position_mode_opcodes:
        return 0  # position mode
    else:
        return 1  # immediate mode


def is_opcode_position_mode(opcode):
    return get_opcode_mode(opcode) == 0


def is_opcode_immediate_mode(opcode):
    return get_opcode_mode(opcode) == 1


def parse_instruction_opcode(instruction):
    if len(str(instruction)) == 1:
        return int(str(instruction))
    return int(str(instruction)[-2:])


def parse_instruction_parameter_mode(parameter_index, instruction):
    # parameter index should be 0, 1, or 2
    match parameter_index:
        case 0:
            if len(str(instruction)) < 3:
                return 0
            return int(str(instruction)[-3])
        case 1:
            if len(str(instruction)) < 4:
                return 0
            return int(str(instruction)[-4])
        case 2:
            if len(str(instruction)) < 5:
                return 0
            return int(str(instruction)[-5])
        case _:
            raise ValueError(f"Invalid parameter index {parameter_index}")


def run_program(codes, inputs, feedback_mode=False):
    current_index = 0
    running = True
    input_count = 0
    output = 0
    while running:
        instruction = codes[current_index]
        current_opcode = parse_instruction_opcode(instruction)
        parameter_1_mode = parse_instruction_parameter_mode(0, instruction)
        parameter_2_mode = parse_instruction_parameter_mode(1, instruction)
        match current_opcode:
            case 1:
                if parameter_1_mode == 1:
                    val_1 = codes[current_index + 1]
                else:
                    val_1 = codes[codes[current_index + 1]]

                if parameter_2_mode == 1:
                    val_2 = codes[current_index + 2]
                else:
                    val_2 = codes[codes[current_index + 2]]

                # "Parameters that an instruction writes to will never be in immediate mode"
                codes[codes[current_index + 3]] = val_1 + val_2
                current_index += 4
            case 2:
                if parameter_1_mode == 1:
                    val_1 = codes[current_index + 1]
                else:
                    val_1 = codes[codes[current_index + 1]]

                if parameter_2_mode == 1:
                    val_2 = codes[current_index + 2]
                else:
                    val_2 = codes[codes[current_index + 2]]

                # "Parameters that an instruction writes to will never be in immediate mode"
                codes[codes[current_index + 3]] = val_1 * val_2
                current_index += 4
            case 3:
                # "Parameters that an instruction writes to will never be in immediate mode"
                codes[codes[current_index + 1]] = inputs[input_count]
                input_count += 1
                current_index += 2
            case 4:
                if parameter_1_mode == 1:
                    output = codes[current_index + 1]
                else:
                    output = codes[codes[current_index + 1]]
                current_index += 2
                if feedback_mode:
                    running = False
                    break
            case 5:
                if parameter_1_mode == 1:
                    val_1 = codes[current_index + 1]
                else:
                    val_1 = codes[codes[current_index + 1]]

                if parameter_2_mode == 1:
                    val_2 = codes[current_index + 2]
                else:
                    val_2 = codes[codes[current_index + 2]]

                if val_1 != 0:
                    current_index = val_2
                else:
                    current_index += 3
            case 6:
                if parameter_1_mode == 1:
                    val_1 = codes[current_index + 1]
                else:
                    val_1 = codes[codes[current_index + 1]]

                if parameter_2_mode == 1:
                    val_2 = codes[current_index + 2]
                else:
                    val_2 = codes[codes[current_index + 2]]

                if val_1 == 0:
                    current_index = val_2
                else:
                    current_index += 3
            case 7:
                if parameter_1_mode == 1:
                    val_1 = codes[current_index + 1]
                else:
                    val_1 = codes[codes[current_index + 1]]

                if parameter_2_mode == 1:
                    val_2 = codes[current_index + 2]
                else:
                    val_2 = codes[codes[current_index + 2]]

                if val_1 < val_2:
                    val_3 = 1
                else:
                    val_3 = 0

                # "Parameters that an instruction writes to will never be in immediate mode"
                codes[codes[current_index + 3]] = val_3
                current_index += 4
            case 8:
                if parameter_1_mode == 1:
                    val_1 = codes[current_index + 1]
                else:
                    val_1 = codes[codes[current_index + 1]]

                if parameter_2_mode == 1:
                    val_2 = codes[current_index + 2]
                else:
                    val_2 = codes[codes[current_index + 2]]

                if val_1 == val_2:
                    val_3 = 1
                else:
                    val_3 = 0

                # "Parameters that an instruction writes to will never be in immediate mode"
                codes[codes[current_index + 3]] = val_3
                current_index += 4
            case 99:
                running = False
                break
            case _:
                raise ValueError(
                    f"Invalid opcode {current_opcode} encountered in processing loop at index {current_index}. Codes: {codes}"
                )

    return codes, current_index, output
