package main

import (
	_ "embed"
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

	cycleHight := 0

	for i := 0; i < numRocks; i++ {
		cyclestart := cycleHight
		rockNum := i % len(rockShapes)
		rock := rockShapes[rockNum]
		offset := 4
		if i == 0 {
			offset = 3
		}

		// cycle check
		// 0-6 top positions normalized
		// 7 rock shape
		// 8 jet index
		pattern := [9]int{}
		dist := distances(colMax)
		copy(pattern[:7], dist[:])
		pattern[7] = rockNum
		pattern[8] = jetIndex % len(jetPattern)
		if c, exists := cycles[pattern]; exists {
			c.hit++
			cycles[pattern] = c
			cycleHight += c.height
			continue
		}

		// simulate rock falling
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

		cycles[pattern] = Cycle{height: cycleHight - cyclestart}

		truncate(rockPositions, colMax)
	}

	return cycleHight + 1 // +1 to account for the height starting from 0
}

func distances(colMax [7]int) [7]int {
	out := [7]int{}
	min := math.MaxInt64
	for _, v := range colMax {
		if v < min {
			min = v
		}
	}
	// normalize so that we have a "pattern"
	for i, v := range colMax {
		out[i] = v - min
	}
	return out
}

func truncate(points map[Point]struct{}, colMax [7]int) int {
	min := math.MaxInt64
	for i := 0; i < 7; i++ {
		if colMax[i] < min {
			min = colMax[i]
		}
	}

	// because we start with an offset end would remove too much
	min = min - 4

	// if the point is below the global min, remove it
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
