from aoc23.inputhelper import read_line_split

data = read_line_split("resources/day5-ex.txt")


def get_numbers(values: str):
    return list(map(lambda v: int(v), filter(lambda v: v != '', map(lambda v: v.strip(), values.split(" ")))))


seeds = get_numbers(data[0].split(":")[1])
#seed_to_soil = []
#soil_to_fertilizer = []
#fertilizer_to_water = []
#water_to_light = []
#light_to_temperature = []
#temperature_to_humidity = []
#humidity_to_location = []

all_map = []

for i, row in enumerate(data):
    if len(row) != 0 and not row[0].isdigit() and not row.startswith("seeds:"):
        entry = []

        for k in range(i + 1, len(data)):
            if len(data[k]) != 0 and data[k][0].isdigit():
                entry.append(get_numbers(data[k]))
            else:
                break

        if len(entry) != 0:
            all_map.append(entry)

for entry in all_map:
    range(entry[0][0], entry[0][2] + 1)
    range(entry[0][1], entry[0][2] + 1)



print(all_map)
