package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

type sensor struct {
	x, y, bx, by, r int // r = manhattan distance to the closest beacon
}

func main() {
	println("part 1:", part1(input, 2000000))
	println("part 2:", part2(input, 4000000))
}

func parse(in string) []sensor {
	var out []sensor
	for _, line := range strings.Split(in, "\n") {
		var s sensor
		if _, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.x, &s.y, &s.bx, &s.by); err != nil {
			panic(err)
		}
		s.r = abs(s.x-s.bx) + abs(s.y-s.by)
		out = append(out, s)
	}
	return out
}

// part1 merges the covered x-intervals on the row and subtracts beacons on it.
func part1(in string, row int) int {
	sensors := parse(in)
	type span struct{ lo, hi int }
	var spans []span
	for _, s := range sensors {
		w := s.r - abs(s.y-row)
		if w >= 0 {
			spans = append(spans, span{s.x - w, s.x + w})
		}
	}
	sort.Slice(spans, func(i, j int) bool { return spans[i].lo < spans[j].lo })

	covered := 0
	cur := spans[0]
	for _, sp := range spans[1:] {
		if sp.lo > cur.hi+1 {
			covered += cur.hi - cur.lo + 1
			cur = sp
		} else if sp.hi > cur.hi {
			cur.hi = sp.hi
		}
	}
	covered += cur.hi - cur.lo + 1

	beacons := map[int]bool{}
	for _, s := range sensors {
		if s.by == row && !beacons[s.bx] {
			beacons[s.bx] = true
			covered--
		}
	}
	return covered
}

// part2: the search area is 4M x 4M, far too big to scan cell by cell.
// The only uncovered cell must lie just outside the border of some sensor
// diamond (distance r+1). Those borders are diagonal lines y = x + a and
// y = -x + b, so candidate cells are the intersections of an "a" line with a
// "b" line - only O(n^2) points. Each candidate is checked against all sensors.
func part2(in string, max int) int {
	sensors := parse(in)
	var as, bs []int
	for _, s := range sensors {
		d := s.r + 1
		as = append(as, s.y-s.x+d, s.y-s.x-d)
		bs = append(bs, s.y+s.x+d, s.y+s.x-d)
	}
	for _, a := range as {
		for _, b := range bs {
			// y = x + a and y = -x + b => x = (b - a) / 2
			if (b-a)%2 != 0 {
				continue
			}
			x, y := (b-a)/2, (a+b)/2
			if x < 0 || y < 0 || x > max || y > max {
				continue
			}
			if !coveredByAny(sensors, x, y) {
				return x*4000000 + y
			}
		}
	}
	panic("no uncovered cell found")
}

func coveredByAny(sensors []sensor, x, y int) bool {
	for _, s := range sensors {
		if abs(s.x-x)+abs(s.y-y) <= s.r {
			return true
		}
	}
	return false
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
