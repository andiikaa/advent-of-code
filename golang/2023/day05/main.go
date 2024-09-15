package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

//go:embed input.txt
var input string

func main() {
	seeds := seeds(input)
	parsed := parseMaps(input)

	p1 := part1(seeds, parsed)
	fmt.Println("part 1:", p1)

	p2 := part2(seeds, parsed)
	fmt.Println("part 2:", p2)
}

func part2(seeds []int, parsed [][][]int) int {
	// todo this is too slow
	lowest := math.MaxInt64
	mu := sync.Mutex{}
	wait := sync.WaitGroup{}
	for i := 0; i < len(seeds); i = i + 2 {
		wait.Add(1)
		go func(k int) {
			tmpLowest := math.MaxInt64
			for j := seeds[k]; j < seeds[k]+seeds[k+1]; j++ {
				v := part1([]int{j}, parsed)
				if v < tmpLowest {
					tmpLowest = v
				}
			}
			mu.Lock()
			if tmpLowest < lowest {
				lowest = tmpLowest
			}
			mu.Unlock()
			wait.Done()
		}(i)
	}

	wait.Wait()

	return lowest
}

func part1(seeds []int, parsed [][][]int) int {
	lowest := math.MaxInt64

	for _, seed := range seeds {
		next := seed

		// this is a category (e.g soil)
		for _, m := range parsed {
			for _, row := range m {
				if next >= row[1] && next < row[1]+row[2] {
					next = row[0] + (next - row[1])
					break
				}
			}
		}
		if next < lowest {
			lowest = next
		}
	}
	return lowest
}

func parseMaps(in string) [][][]int {
	result := make([][][]int, 0)

	split := strings.Split(in, "map:")
	for i, v := range split {
		if i == 0 {
			continue
		}

		entry := make([][]int, 0)

		rows := strings.Split(v, "\n\n")
		rows = strings.Split(rows[0], "\n")
		for _, row := range rows {
			if intRow := stringsToIntList(row); len(intRow) > 0 {
				entry = append(entry, intRow)
			}
		}
		result = append(result, entry)
	}
	return result
}

func seeds(in string) []int {
	s := strings.Split(in, "\n")
	seedS, _ := strings.CutPrefix(s[0], "seeds: ")
	return stringsToIntList(seedS)
}

func stringsToIntList(s string) []int {
	intStr := strings.Split(s, " ")

	result := make([]int, 0)

	for _, v := range intStr {
		if i, err := strconv.Atoi(v); err == nil {
			result = append(result, i)
		}
	}

	return result
}
