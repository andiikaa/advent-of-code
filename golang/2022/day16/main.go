package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}

// cave holds only the valves with a positive flow rate (index 0..n-1) plus
// the start AA, with the shortest travel times between them.
type cave struct {
	rates []int
	dist  [][]int // dist[i][j] between useful valves
	start []int   // distance from AA to each useful valve
}

var lineRe = regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)

func parse(in string) cave {
	idx := map[string]int{}
	var names []string
	var rates []int
	var edges [][]string
	for _, line := range strings.Split(in, "\n") {
		m := lineRe.FindStringSubmatch(line)
		if m == nil {
			panic("bad line: " + line)
		}
		r, _ := strconv.Atoi(m[2])
		idx[m[1]] = len(names)
		names = append(names, m[1])
		rates = append(rates, r)
		edges = append(edges, strings.Split(m[3], ", "))
	}

	// Floyd-Warshall over all valves (n is ~60, so n^3 is tiny).
	n := len(names)
	const inf = 1 << 30
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			if i != j {
				d[i][j] = inf
			}
		}
		for _, e := range edges[i] {
			d[i][idx[e]] = 1
		}
	}
	for k := range n {
		for i := range n {
			for j := range n {
				if d[i][k]+d[k][j] < d[i][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}

	var useful []int
	for i, r := range rates {
		if r > 0 {
			useful = append(useful, i)
		}
	}
	c := cave{}
	for _, u := range useful {
		c.rates = append(c.rates, rates[u])
		c.start = append(c.start, d[idx["AA"]][u])
		row := make([]int, len(useful))
		for j, v := range useful {
			row[j] = d[u][v]
		}
		c.dist = append(c.dist, row)
	}
	return c
}

// bestPerSet explores every reachable order of opening valves within the
// time limit and returns, for each set of opened valves (bitmask), the most
// pressure any single actor can release by opening exactly that set.
func bestPerSet(c cave, minutes int) []int {
	best := make([]int, 1<<len(c.rates))
	var dfs func(pos, timeLeft, opened, pressure int)
	dfs = func(pos, timeLeft, opened, pressure int) {
		if pressure > best[opened] {
			best[opened] = pressure
		}
		for next := range c.rates {
			if opened&(1<<next) != 0 {
				continue
			}
			var travel int
			if pos < 0 {
				travel = c.start[next]
			} else {
				travel = c.dist[pos][next]
			}
			t := timeLeft - travel - 1 // move there, then one minute to open
			if t <= 0 {
				continue
			}
			dfs(next, t, opened|1<<next, pressure+t*c.rates[next])
		}
	}
	dfs(-1, minutes, 0, 0)
	return best
}

func part1(in string) int {
	best := 0
	for _, p := range bestPerSet(parse(in), 30) {
		best = max(best, p)
	}
	return best
}

// part2: simulating you and the elephant together blows up the state space.
// Instead, both open disjoint sets of valves, so the answer is the best
// pair of disjoint sets from the single-actor table for 26 minutes.
// sub[m] = best over all subsets of m makes the pairing O(2^n).
func part2(in string) int {
	c := parse(in)
	best := bestPerSet(c, 26)
	n := len(c.rates)
	sub := make([]int, len(best))
	copy(sub, best)
	for bit := range n {
		for m := range sub {
			if m&(1<<bit) != 0 {
				sub[m] = max(sub[m], sub[m^(1<<bit)])
			}
		}
	}
	full := len(best) - 1
	answer := 0
	for m, p := range best {
		answer = max(answer, p+sub[full^m])
	}
	return answer
}
