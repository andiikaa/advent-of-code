from aoc23.inputhelper import read_line_split
import operator

data = read_line_split("resources/day2-input.txt")
max_values = (12, 13, 14)

def round_sum(current: str) -> tuple:
    v = int(current.split(" ")[1].strip())
    if current.endswith("blue"):
        return 0, 0, v
    elif current.endswith("green"):
        return 0, v, 0
    elif current.endswith("red"):
        return v, 0, 0
    raise Exception("Unknown color")

def is_valid(rgb: tuple) -> bool:
    return all([rgb[i] <= max_values[i] for i in range(3)])

valid = []
game_nr = 0
for game in data:
    game = game[game.index(":") + 1:]
    current_is_valid = True
    game_nr += 1

    for pull in game.rsplit(";"):
        rgb = (0, 0, 0)
        for current in pull.split(","):
            rgb = tuple(map(operator.add, rgb, round_sum(current)))
        if not is_valid(rgb):
            current_is_valid = False
            break

    if current_is_valid:
        valid.append(game_nr)

print(sum(valid))