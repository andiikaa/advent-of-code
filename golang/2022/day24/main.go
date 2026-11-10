package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

type pt struct{ r, c int }

// valley is the inner area without walls: rows 0..h-1, cols 0..w-1.
// start is at (-1, 0), goal at (h, w-1).
type valley struct {
	grid  []string // inner rows, walls stripped
	h, w  int
	start pt
	goal  pt
}

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}

func parse(in string) valley {
	lines := strings.Split(in, "\n")
	var grid []string
	for _, l := range lines[1 : len(lines)-1] {
		grid = append(grid, l[1:len(l)-1])
	}
	h, w := len(grid), len(grid[0])
	return valley{
		grid:  grid,
		h:     h,
		w:     w,
		start: pt{-1, strings.IndexByte(lines[0], '.') - 1},
		goal:  pt{h, strings.IndexByte(lines[len(lines)-1], '.') - 1},
	}
}

func mod(a, m int) int { return ((a % m) + m) % m }

// free reports whether cell p is free of blizzards at minute t. Blizzards
// move independently and wrap, so instead of simulating them, look up
// where a blizzard would have to start to be here now.
func (v valley) free(p pt, t int) bool {
	if p == v.start || p == v.goal {
		return true
	}
	if p.r < 0 || p.r >= v.h || p.c < 0 || p.c >= v.w {
		return false
	}
	return v.grid[p.r][mod(p.c-t, v.w)] != '>' &&
		v.grid[p.r][mod(p.c+t, v.w)] != '<' &&
		v.grid[mod(p.r-t, v.h)][p.c] != 'v' &&
		v.grid[mod(p.r+t, v.h)][p.c] != '^'
}

// trip returns the first minute the goal is reached, starting at minute t0.
// BFS by minute: the frontier is the set of positions reachable at time t.
func (v valley) trip(from, to pt, t0 int) int {
	frontier := map[pt]struct{}{from: {}}
	moves := []pt{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for t := t0; ; t++ {
		next := make(map[pt]struct{}, len(frontier)*2)
		for p := range frontier {
			for _, m := range moves {
				q := pt{p.r + m.r, p.c + m.c}
				if q == to {
					return t + 1
				}
				if v.free(q, t+1) {
					next[q] = struct{}{}
				}
			}
		}
		if len(next) == 0 {
			panic("no way through")
		}
		frontier = next
	}
}

func part1(in string) int {
	v := parse(in)
	return v.trip(v.start, v.goal, 0)
}

// part2: goal, back to the start for the snacks, then goal again. Each leg
// is shortest on its own given its start time, since waiting at the start
// or goal is always allowed.
func part2(in string) int {
	v := parse(in)
	t := v.trip(v.start, v.goal, 0)
	t = v.trip(v.goal, v.start, t)
	return v.trip(v.start, v.goal, t)
}
