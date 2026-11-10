package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

//go:embed example2.txt
var example2 string

type pos struct{ x, y int }

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}

func part1(in string) int { return simulate(in, 2) }

func part2(in string) int { return simulate(in, 10) }

// simulate moves a rope of n knots and counts the distinct tail positions.
func simulate(in string, n int) int {
	dirs := map[string]pos{"R": {1, 0}, "L": {-1, 0}, "U": {0, 1}, "D": {0, -1}}
	knots := make([]pos, n)
	visited := map[pos]struct{}{{}: {}}
	for _, line := range strings.Split(in, "\n") {
		f := strings.Fields(line)
		d := dirs[f[0]]
		steps, err := strconv.Atoi(f[1])
		if err != nil {
			panic(err)
		}
		for range steps {
			knots[0].x += d.x
			knots[0].y += d.y
			for i := 1; i < n; i++ {
				knots[i] = follow(knots[i], knots[i-1])
			}
			visited[knots[n-1]] = struct{}{}
		}
	}
	return len(visited)
}

// follow moves the knot t one step towards h if they no longer touch.
// With longer ropes a knot can be two steps away diagonally, which sign()
// handles the same way as the straight and knight-move cases.
func follow(t, h pos) pos {
	dx, dy := h.x-t.x, h.y-t.y
	if abs(dx) <= 1 && abs(dy) <= 1 {
		return t
	}
	return pos{t.x + sign(dx), t.y + sign(dy)}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func sign(v int) int {
	switch {
	case v > 0:
		return 1
	case v < 0:
		return -1
	}
	return 0
}
