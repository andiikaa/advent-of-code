from aoc23.inputhelper import read_line_split

data = read_line_split("resources/day4-in.txt")

games = []
for i, card in enumerate(data):
    lr = card.split(":")[1].split("|")
    wins = list(map(lambda v: int(v.strip()), filter(lambda v: len(v.strip()) != 0, lr[0].split(" "))))
    my = list(map(lambda v: int(v.strip()), filter(lambda v: len(v.strip()) != 0, lr[1].split(" "))))
    games.append((i, wins, my))

new_cards = []
total_cards = 0
loop_cards = games.copy()

while True:
    for game in loop_cards:
        game_nr = game[0]
        win = game[1]
        my = game[2]
        count = 0

        for win_nr in win:
            if win_nr in my:
                count += 1

        if count != 0:
            cards = list(range(game_nr + 1, game_nr + count + 1))
            new_cards.extend(cards)

    total_cards += len(loop_cards)

    if len(new_cards) == 0:
        break
    loop_cards.clear()
    for c in new_cards:
        if c < len(games):
            loop_cards.append(games[c])
    new_cards = []

# 9997537
print(total_cards)
