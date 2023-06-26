def solve_part_one():
    # opcode 1: add values at a and b and store at c
    # opcode 2: multiply values at a and b and store at c
    # opcode 99: cease

    # Get codes
    codes = []
    with open('./inputs/day2.txt') as f:
        for line in f.readlines():
            values = line.split(',')
            for value in values:
              codes.append(int(value))
    
    # Replace position 1 with value 12 and position 2 with value 2
    codes[1] = 12
    codes[2] = 2

    # Process codes
    return run_program(codes)


def solve_part_two():
    # Get codes
    codes_original = []
    with open('./inputs/day2.txt') as f:
        for line in f.readlines():
            values = line.split(',')
            for value in values:
              codes_original.append(int(value))

    for i in range(0, 100):
        for j in range(0, 100):
            codes = codes_original.copy()
            codes[1] = i
            codes[2] = j
            try:
              if run_program(codes) == 19690720:
                  return i, j, 100 * i + j
            except ValueError:
                continue
            
    return None, None, None


def run_program(codes):
    current_index = 0
    running = True
    while running:
        current_opcode = codes[current_index]
        match current_opcode:
            case 1:
                new_val = codes[codes[current_index+1]] + codes[codes[current_index+2]]
                codes[codes[current_index+3]] = new_val
            case 2:
                new_val = codes[codes[current_index+1]] * codes[codes[current_index+2]]
                codes[codes[current_index+3]] = new_val
            case 99:
                running = False
                break
            case _:
                raise ValueError(f'Invalid opcode {current_opcode} encountered in processing loop at index {current_index}')
        current_index += 4

    return codes[0]
    

print(solve_part_one())
print(solve_part_two())