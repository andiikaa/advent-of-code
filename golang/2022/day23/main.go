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

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}

func parse(in string) map[pt]struct{} {
	elves := map[pt]struct{}{}
	for r, line := range strings.Split(in, "\n") {
		for c, ch := range line {
			if ch == '#' {
				elves[pt{r, c}] = struct{}{}
			}
		}
	}
	return elves
}

// checks[d] are the three cells that must be free to move in direction d;
// the first entry is the move itself. Order: N, S, W, E.
var checks = [4][3]pt{
	{{-1, 0}, {-1, 1}, {-1, -1}},
	{{1, 0}, {1, 1}, {1, -1}},
	{{0, -1}, {-1, -1}, {1, -1}},
	{{0, 1}, {-1, 1}, {1, 1}},
}

// round runs one round starting with direction index first and returns the
// new positions and whether any elf moved.
func round(elves map[pt]struct{}, first int) (map[pt]struct{}, bool) {
	has := func(p pt) bool { _, ok := elves[p]; return ok }
	proposals := make(map[pt]pt, len(elves)) // elf -> target
	count := make(map[pt]int, len(elves))    // target -> number of proposers
	for e := range elves {
		alone := true
		for dr := -1; dr <= 1 && alone; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if (dr != 0 || dc != 0) && has(pt{e.r + dr, e.c + dc}) {
					alone = false
					break
				}
			}
		}
		if alone {
			continue
		}
		for k := range 4 {
			d := checks[(first+k)%4]
			free := true
			for _, o := range d {
				if has(pt{e.r + o.r, e.c + o.c}) {
					free = false
					break
				}
			}
			if free {
				t := pt{e.r + d[0].r, e.c + d[0].c}
				proposals[e] = t
				count[t]++
				break
			}
		}
	}

	moved := false
	next := make(map[pt]struct{}, len(elves))
	for e := range elves {
		if t, ok := proposals[e]; ok && count[t] == 1 {
			next[t] = struct{}{}
			moved = true
		} else {
			next[e] = struct{}{}
		}
	}
	return next, moved
}

func part1(in string) int {
	elves := parse(in)
	for i := range 10 {
		elves, _ = round(elves, i%4)
	}
	minR, minC, maxR, maxC := 1<<31, 1<<31, -1<<31, -1<<31
	for e := range elves {
		minR, maxR = min(minR, e.r), max(maxR, e.r)
		minC, maxC = min(minC, e.c), max(maxC, e.c)
	}
	return (maxR-minR+1)*(maxC-minC+1) - len(elves)
}

// part2 runs rounds until nobody moves.
func part2(in string) int {
	elves := parse(in)
	for i := 0; ; i++ {
		var moved bool
		elves, moved = round(elves, i%4)
		if !moved {
			return i + 1
		}
	}
}
