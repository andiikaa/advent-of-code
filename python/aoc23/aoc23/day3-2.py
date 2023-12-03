from aoc23.inputhelper import read_line_split

#data = read_line_split("resources/day3-example.txt")
# data = read_line_split("resources/day3-example2.txt")
data = read_line_split("resources/day3-in.txt")


# (row_min, row_max, col_min, col_max)
def symbol_range(sym: tuple) -> tuple:
    row_min = sym[0] - 1 if sym[0] > 0 else 0
    row_max = sym[0] + 1 if sym[0] < len(data) - 1 else len(data) - 1
    col_min = sym[1] - 1 if sym[1] > 0 else 0
    col_max = sym[1] + 1 if sym[1] < len(data[0]) - 1 else len(data[0]) - 1
    return row_min, row_max, col_min, col_max


# (row, start, end)
digits = set()
# (row, col)
symbols = set()
for i, row in enumerate(data):
    # (row, start, end)
    digit_pos = None
    for j, c in enumerate(row):

        if c.isdigit() and digit_pos is None:
            digit_pos = (i, j, j)
        elif not c.isdigit() and digit_pos is not None:
            digit_pos = (digit_pos[0], digit_pos[1], j)
            digits.add(digit_pos)
            digit_pos = None

        if c != "." and not c.isdigit():
            symbols.add((i, j))
        if j == len(row) - 1 and digit_pos is not None:
            digit_pos = (digit_pos[0], digit_pos[1], j + 1)
            digits.add(digit_pos)
            digit_pos = None


def is_in_range(pos: tuple, sym: tuple) -> bool:
    rowrange = pos[0] in range(sym[0], sym[1] + 1)
    allowed_cols = range(sym[2], sym[3] + 1)
    pos_cols = range(pos[1], pos[2])
    return any(pos_col in allowed_cols for pos_col in pos_cols) and rowrange


found = []
for s in symbols:
    if data[s[0]][s[1]] != '*':
        continue
    r = symbol_range(s)
    filtered = list(filter(lambda d: is_in_range(d, r), digits))
    if len(filtered) == 2:
        mapped = list(map(lambda pos: int(data[pos[0]][pos[1]:pos[2]]), filtered))
        found.append(mapped[0] * mapped[1])

print(sum(found))
