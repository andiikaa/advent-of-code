from aoc23.inputhelper import read_line_split


def replace_words(line: str):
    line = line.replace('one', 'o1e')
    line = line.replace('two', 't2o')
    line = line.replace('three', 't3e')
    line = line.replace('four', 'f4r')
    line = line.replace('five', 'f5e')
    line = line.replace('six', 's6x')
    line = line.replace('seven', 's7n')
    line = line.replace('eight', 'e8t')
    line = line.replace('nine', 'n9e')
    return line


data = read_line_split("resources/day1-in.txt")
numbers = []
for line in data:
    line = replace_words(line)
    n = ''
    for char in line:
        if char.isdigit():
            n += char
    numbers.append(int(n[0] + n[-1]))

print(sum(numbers))
