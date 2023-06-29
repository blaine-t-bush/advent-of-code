import re


def map_orbits():
    orbits = {}
    with open('./inputs/day6.txt') as f:
        for line in f.readlines():
            split = line.split(")")
            orbits[split[1].rstrip('\n')] = split[0]
    return orbits


def count_orbits(orbits):
    return count_direct_orbits(orbits) + count_indirect_orbits(orbits)


def count_direct_orbits(orbits):
    return len(orbits)


def count_indirect_orbits(orbits):
    count = 0
    for key in orbits:
        count += count_indirect_orbit(key, orbits)
    return count


def count_indirect_orbit(start, orbits):
    count = count_chain(start, orbits)
    if count > 0:
        return count - 1
    return 0


def count_chain(start, orbits):
    count = 0
    if start in orbits:
        count += 1
        return count + count_chain(orbits[start], orbits)
    else:
        return count


def get_parent(start, orbits):
    if start in orbits:
        return orbits[start]
    return None


def get_parents(start, orbits):
    parents = []
    while True:
        parent = get_parent(start, orbits)
        if parent is not None:
            parents.append(parent)
            start = parent
        else:
            break
    return parents


def find_common_ancestor(list_a, list_b):
    for a in list_a:
        for b in list_b:
            if a == b:
                return a
    return None


def get_distance(end, list):
    count = 0
    for item in list:
        if item == end:
            break
        count += 1
    return count


orbits = map_orbits()
# Part 1
print(count_orbits(orbits))
# Part 2
parents_you = get_parents("YOU", orbits)
parents_san = get_parents("SAN", orbits)
ancestor = find_common_ancestor(parents_you, parents_san)
print(get_distance(ancestor, parents_you) +
      get_distance(ancestor, parents_san))
