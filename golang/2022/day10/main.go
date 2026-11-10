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

func main() {
	println("part 1:", part1(input))
	println("part 2:")
	println(part2(input))
}

// run calls tick once per cycle with the cycle number (1-based) and the
// value of X during that cycle.
func run(in string, tick func(cycle, x int)) {
	x, cycle := 1, 0
	for _, line := range strings.Split(in, "\n") {
		f := strings.Fields(line)
		cycle++
		tick(cycle, x)
		if f[0] == "addx" {
			cycle++
			tick(cycle, x)
			v, err := strconv.Atoi(f[1])
			if err != nil {
				panic(err)
			}
			x += v
		}
	}
}

func part1(in string) int {
	sum := 0
	run(in, func(cycle, x int) {
		if (cycle-20)%40 == 0 {
			sum += cycle * x
		}
	})
	return sum
}

// part2 renders the 40x6 CRT. The answer is the capital letters it shows.
func part2(in string) string {
	var sb strings.Builder
	run(in, func(cycle, x int) {
		col := (cycle - 1) % 40
		if col >= x-1 && col <= x+1 {
			sb.WriteByte('#')
		} else {
			sb.WriteByte('.')
		}
		if col == 39 {
			sb.WriteByte('\n')
		}
	})
	return strings.TrimRight(sb.String(), "\n")
}
