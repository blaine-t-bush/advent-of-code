def read_data():
    with open("./inputs/day8.txt") as f:
        for line in f.readlines():
            data = line.rstrip("\n")
    return data


def count_layers(w, h, data):
    layer_length = w * h
    return int(len(data) / layer_length)


def get_layers(w, h, data):
    layer_count = count_layers(w, h, data)
    layers = []
    for i in range(layer_count):
        layers.append(data[i * w * h : (i + 1) * w * h])

    return layers


def count_digit(string, digit):
    count = 0
    for char in string:
        if char == digit:
            count += 1

    return count


def find_layer_with_fewest_zeros(w, h, layers):
    fewest_zeros_layer = None
    fewest_zeros = w * h
    for layer in layers:
        count = count_digit(layer, "0")
        if count < fewest_zeros:
            fewest_zeros = count
            fewest_zeros_layer = layer
    return fewest_zeros_layer


def stack_image(w, h, layers):
    stacked_image = []
    for y in range(h):
        for x in range(w):
            color = "2"
            for layer in layers:
                if color == "2" and layer[y * w + x] != "2":
                    color = layer[y * w + x]
                    break
            stacked_image.append(color)

    return stacked_image


def solve_part_one():
    width = 25
    height = 6
    data = read_data()
    layers = get_layers(width, height, data)
    layer = find_layer_with_fewest_zeros(width, height, layers)
    return count_digit(layer, "1") * count_digit(layer, "2")


def solve_part_two():
    width = 25
    height = 6
    data = read_data()
    layers = get_layers(width, height, data)
    stacked_image = stack_image(width, height, layers)

    for y in range(height):
        for x in range(width):
            pixel = stacked_image[y * width + x]
            if pixel == "0":
                print("█", end="")
            else:
                print(" ", end="")
        print()


# print(solve_part_one())
solve_part_two()
