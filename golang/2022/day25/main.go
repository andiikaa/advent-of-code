package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

func main() {
	println("part 1:", part1(input))
}

// part1 sums the SNAFU numbers and writes the sum as SNAFU.
// (Day 25 has no second puzzle: its second star is given for all 49 others.)
func part1(in string) string {
	sum := 0
	for _, line := range strings.Split(in, "\n") {
		sum += fromSNAFU(line)
	}
	return toSNAFU(sum)
}

// fromSNAFU: base 5 with digits 2, 1, 0, - (-1), = (-2).
func fromSNAFU(s string) int {
	n := 0
	for _, ch := range s {
		n *= 5
		n += strings.IndexRune("=-012", ch) - 2
	}
	return n
}

// toSNAFU: take n mod 5; digits 3 and 4 become = and - with a carry of 1
// into the next place (3 = 5 - 2, 4 = 5 - 1).
func toSNAFU(n int) string {
	if n == 0 {
		return "0"
	}
	var out []byte
	for n > 0 {
		d := n % 5
		n /= 5
		if d > 2 {
			d -= 5
			n++
		}
		out = append(out, "=-012"[d+2])
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return string(out)
}
