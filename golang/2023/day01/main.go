package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	p1 := part1(input)
	println("part 1:", p1)
}

func part1(in string) int {
	split := strings.Split(in, "\n")
	sum := 0
	for _, s := range split {
		var first, second int
		for i := 0; i < len(s); i++ {
			if strings.ContainsAny(s[i:i+1], "0123456789") {
				first, _ = strconv.Atoi(s[i : i+1])
				break
			}
		}

		for i := len(s) - 1; i >= 0; i-- {
			if strings.ContainsAny(s[i:i+1], "0123456789") {
				second, _ = strconv.Atoi(s[i : i+1])
				break
			}
		}
		sum += first*10 + second
	}

	return sum
}
