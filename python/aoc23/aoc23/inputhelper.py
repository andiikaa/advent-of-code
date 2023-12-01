
def read_line_split(file):
    f = open(file, 'r')
    data = f.read()
    f.close()
    return data.splitlines()
