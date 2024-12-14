package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

type Rock struct {
	shape []Point
}

type Cycle struct {
	height    int
	rockIndex int
	jetIndex  int
	hit       int
}

var rockShapes = []Rock{
	{shape: []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},         // -
	{shape: []Point{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}}, // +
	{shape: []Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}}, // L
	{shape: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},         // |
	{shape: []Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}},         // square
}

func main() {
	p1 := part1(input)
	println("part 1:", p1)
	p2 := part2(input)
	println("part 2:", p2)
}

func part1(in string) int {
	return simulateFallingRocks(in, 2022)
}

func part2(in string) int {
	return simulateFallingRocks(in, 1_000_000_000_000)
}

func simulateFallingRocks(jetPattern string, numRocks int) int {
	jetPattern = strings.TrimSpace(jetPattern)
	jetIndex := 0
	colMax := [7]int{}
	rockPositions := make(map[Point]struct{})
	cycles := make(map[[9]int]Cycle)

	maxHeight := 0
	cycleHight := 0

	for i := 0; i < numRocks; i++ {
		rock := rockShapes[i%len(rockShapes)]
		offset := 4
		if i == 0 {
			offset = 3
		}

		rockPos := Point{2, cycleHight + offset}

		for {
			// Apply jet push
			jet := jetPattern[jetIndex%len(jetPattern)]
			jetIndex++
			if jet == '>' {
				if canMove(rock, rockPos, Point{1, 0}, rockPositions) {
					rockPos.x++
				}
			} else if jet == '<' {
				if canMove(rock, rockPos, Point{-1, 0}, rockPositions) {
					rockPos.x--
				}
			} else {
				panic("invalid jet")
			}

			// Apply gravity
			if canMove(rock, rockPos, Point{0, -1}, rockPositions) {
				rockPos.y--
			} else {
				break
			}
		}

		// Add rock to positions
		for _, p := range rock.shape {
			newPos := Point{rockPos.x + p.x, rockPos.y + p.y}
			rockPositions[newPos] = struct{}{}

			// the hightest in the cycle
			if newPos.y > cycleHight {
				cycleHight = newPos.y
			}

			// save the max postions for each column
			if colMax[newPos.x] < newPos.y {
				colMax[newPos.x] = newPos.y
			}
		}

		//cycle check
		pattern := [9]int{}
		copy(pattern[:7], colMax[:])
		min := math.MaxInt64
		for _, v := range pattern[:7] {
			if v < min {
				min = v
			}
		}

		// normalize so that we have a "pattern"
		for i, v := range pattern[:7] {
			pattern[i] = v - min
		}

		pattern[7] = (jetIndex - 1) % len(jetPattern)
		pattern[8] = i % len(rockShapes)

		if c, ok := cycles[pattern]; ok {
			c.hit++
			rocksFallen := i - c.rockIndex
			if rocksFallen+i < numRocks {
				hDiff := cycleHight - c.height
				cycleHight += hDiff
				i += rocksFallen
				fmt.Println("Cycle found", pattern, "hits", c.hit)

				var keysToDelete []Point
				for p := range rockPositions {
					keysToDelete = append(keysToDelete, p)
				}

				for _, p := range keysToDelete {
					delete(rockPositions, p)
				}

				for _, p := range keysToDelete {
					rockPositions[Point{p.x, p.y + hDiff}] = struct{}{}
				}

				for i, v := range colMax {
					colMax[i] = v + hDiff
				}

				truncate(rockPositions)

			}
		} else {
			cycles[pattern] = Cycle{height: maxHeight, rockIndex: i, jetIndex: jetIndex, hit: 0}
		}
	}

	return cycleHight + 1 // +1 to account for the height starting from 0
}

func truncate(points map[Point]struct{}) int {
	max := [7]int{}

	for p := range points {
		if max[p.x] < p.y {
			max[p.x] = p.y
		}
	}

	min := math.MaxInt64
	for i := 0; i < 7; i++ {
		if max[i] < min {
			min = max[i]
		}
	}

	if min < math.MaxInt64 {
		for p := range points {
			if p.y < min {
				delete(points, p)
			}
		}
	}

	return min
}

func canMove(rock Rock, pos, delta Point, rockPositions map[Point]struct{}) bool {
	for _, p := range rock.shape {
		newPos := Point{pos.x + p.x + delta.x, pos.y + p.y + delta.y}
		if newPos.x < 0 || newPos.x >= 7 || newPos.y < 0 {
			return false
		}
		if _, exists := rockPositions[newPos]; exists {
			return false
		}
	}
	return true
}
