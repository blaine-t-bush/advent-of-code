from functools import cmp_to_key

def get_pairs():
  pairs = []
  pair = []

  with open("./2022/day13/input.txt", "r") as input_file:
    while True:
      line = input_file.readline()
      if not line:
        pairs.append(pair)
        break
      elif line == "\n":
        pairs.append(pair)
        pair = []
      else:
        pair.append(eval(line))

  return pairs

def combine_pairs(pairs):
  combined = []
  for pair in pairs:
    combined.append(pair[0])
    combined.append(pair[1])
  return combined

def compare_ints(left, right):
  if left < right:
    return -1
  elif left > right:
    return 1
  else:
    return 0

def compare_lists(left, right):
  for i in range(0, min(len(left), len(right))):
    subleft = left[i]
    subright = right[i]

    if isinstance(subleft, list) and isinstance(subright, list):
      val = compare_lists(subleft, subright)
    elif ~isinstance(subleft, list) and isinstance(subright, list):
      val = compare_lists([subleft], subright)
    elif isinstance(subleft, list) and ~isinstance(subright, list):
      val = compare_lists(subleft, [subright])
    elif ~isinstance(subleft, list) and ~isinstance(subright, list):
      val = compare_ints(subleft, subright)
    
    if val == -1 or val == 1:
      return val
  
  if len(left) < len(right):
    return -1
  elif len(left) > len(right):
    return 1
  else:
    return 0

# part 1
def solve_part1():
  pairs = get_pairs()
  indices = []
  for i in range(0, len(pairs)):
    if compare_lists(pairs[i][0], pairs[i][1]) == 1:
      indices.append(i+1)
  print(sum(indices))

# part 2
def solve_part2():
  pairs = combine_pairs(get_pairs())
  pairs.append([[2]])
  pairs.append([[6]])
  ordered = sorted(pairs, key=cmp_to_key(compare_lists))

  for i in range(0, len(ordered)):
    if ordered[i] == [[2]]:
      packet1 = i+1
      print("packet found")
    elif ordered[i] == [[6]]:
      packet2 = i+1
      print("packet found")

  print(packet1*packet2)

solve_part2()