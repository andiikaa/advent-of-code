from aoc23.inputhelper import read_line_split

data = read_line_split("resources/day4-in.txt")

games = []
for card in data:
    lr = card.split(":")[1].split("|")
    wins = list(map(lambda v: int(v.strip()), filter(lambda v: len(v.strip()) != 0, lr[0].split(" "))))
    my = list(map(lambda v: int(v.strip()), filter(lambda v: len(v.strip()) != 0, lr[1].split(" "))))
    games.append((wins, my))

score = []
for game in games:
    win = game[0]
    my = game[1]
    value = 0

    for win_nr in win:
        if win_nr in my:
            value = 1 if value == 0 else value * 2

    score.append(value)

# 26218
print(sum(score))
