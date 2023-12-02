from aoc23.inputhelper import read_line_split

data = read_line_split("resources/day2-in.txt")


def pull_tuple(current: str) -> list:
    v = int(current.split(" ")[1].strip())
    if current.endswith("blue"):
        return [0, 0, v]
    elif current.endswith("green"):
        return [0, v, 0]
    elif current.endswith("red"):
        return [v, 0, 0]
    raise Exception("Unknown color")


def map_to_max(left: list, right: list) -> list:
    rgb = [0, 0, 0]
    for n in range(3):
        if left[n] >= right[n]:
            rgb[n] = left[n]
        else:
            rgb[n] = right[n]
    return rgb


sum = 0
for game in data:
    game = game[game.index(":") + 1:]
    rgb_max = [0, 0, 0]

    for pull in game.rsplit(";"):
        for current in pull.split(","):
            rgb_max = map_to_max(rgb_max, pull_tuple(current))

    sum += rgb_max[0] * rgb_max[1] * rgb_max[2]

print(sum)
