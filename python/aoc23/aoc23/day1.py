from aoc23.inputhelper import read_line_split

data = read_line_split("resources/day1-in.txt")

numbers = []
for line in data:
    n = ''
    for char in line:
        if char.isdigit():
            n += char
    numbers.append(int(n[0] + n[-1]))

print(sum(numbers))
