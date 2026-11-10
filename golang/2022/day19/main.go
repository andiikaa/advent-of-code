package main

import (
	_ "embed"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

// blueprint costs: ore robot (ore), clay robot (ore), obsidian robot (ore,
// clay), geode robot (ore, obsidian)
type blueprint struct {
	id                      int
	oreOre, clayOre         int
	obsOre, obsClay         int
	geodeOre, geodeObsidian int
	maxOre                  int // most ore any robot costs
}

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}

var numRe = regexp.MustCompile(`\d+`)

// parse works for the multi-line example and the one-line input alike.
func parse(in string) []blueprint {
	var nums []int
	for _, s := range numRe.FindAllString(in, -1) {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	var out []blueprint
	for i := 0; i+7 <= len(nums); i += 7 {
		n := nums[i : i+7]
		b := blueprint{n[0], n[1], n[2], n[3], n[4], n[5], n[6], 0}
		b.maxOre = max(b.oreOre, b.clayOre, b.obsOre, b.geodeOre)
		out = append(out, b)
	}
	return out
}

func part1(in string) int {
	sum := 0
	for _, b := range parse(in) {
		sum += b.id * maxGeodes(b, 24)
	}
	return sum
}

func part2(in string) int {
	bps := parse(in)
	product := 1
	for _, b := range bps[:min(3, len(bps))] {
		product *= maxGeodes(b, 32)
	}
	return product
}

// maxGeodes runs a DFS that decides which robot to build next and jumps
// straight to the minute it is built (instead of branching every minute).
// Pruning:
//   - never build more robots of a kind than can be spent per minute (the
//     factory builds one robot per minute, so more ore/clay/obsidian per
//     minute than the most expensive recipe needs is useless)
//   - a geode robot built at time t yields t geodes; stop a branch if even
//     building a geode robot every remaining minute can't beat the best
func maxGeodes(b blueprint, minutes int) int {
	best := 0
	var dfs func(t, ore, clay, obs, rOre, rClay, rObs, geodes int)
	dfs = func(t, ore, clay, obs, rOre, rClay, rObs, geodes int) {
		best = max(best, geodes)
		if geodes+t*(t-1)/2 <= best {
			return
		}

		// wait returns the minutes until the cost can be paid, or -1 if never.
		wait := func(have, rate, cost int) int {
			if have >= cost {
				return 0
			}
			if rate == 0 {
				return -1
			}
			return (cost - have + rate - 1) / rate
		}
		try := func(w int, build func(dt int)) {
			if w < 0 {
				return
			}
			dt := w + 1 // wait, then one minute to build
			if dt >= t {
				return
			}
			build(dt)
		}

		// geode robot
		try(max(wait(ore, rOre, b.geodeOre), waitOrNever(wait(obs, rObs, b.geodeObsidian))), func(dt int) {
			nt := t - dt
			dfs(nt, ore+rOre*dt-b.geodeOre, clay+rClay*dt, obs+rObs*dt-b.geodeObsidian, rOre, rClay, rObs, geodes+nt)
		})
		// obsidian robot
		if rObs < b.geodeObsidian {
			try(max(wait(ore, rOre, b.obsOre), waitOrNever(wait(clay, rClay, b.obsClay))), func(dt int) {
				dfs(t-dt, ore+rOre*dt-b.obsOre, clay+rClay*dt-b.obsClay, obs+rObs*dt, rOre, rClay, rObs+1, geodes)
			})
		}
		// clay robot
		if rClay < b.obsClay {
			try(wait(ore, rOre, b.clayOre), func(dt int) {
				dfs(t-dt, ore+rOre*dt-b.clayOre, clay+rClay*dt, obs+rObs*dt, rOre, rClay+1, rObs, geodes)
			})
		}
		// ore robot
		if rOre < b.maxOre {
			try(wait(ore, rOre, b.oreOre), func(dt int) {
				dfs(t-dt, ore+rOre*dt-b.oreOre, clay+rClay*dt, obs+rObs*dt, rOre+1, rClay, rObs, geodes)
			})
		}
	}
	dfs(minutes, 0, 0, 0, 1, 0, 0, 0)
	return best
}

// waitOrNever maps "never" (-1) to a value that makes max() report never.
func waitOrNever(w int) int {
	if w < 0 {
		return 1 << 30
	}
	return w
}
