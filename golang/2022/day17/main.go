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
	chamberWidth := 7
	rockPositions := make(map[Point]struct{})
	maxHeight := 0

	for i := 0; i < numRocks; i++ {
		rock := rockShapes[i%len(rockShapes)]
		offset := 4
		if i == 0 {
			offset = 3
		}
		if i%1000000 == 0 {
			fmt.Println("Processing rock", i)
		}
		//printPos(rockPositions, chamberWidth)
		//fmt.Println()
		rockPos := Point{2, maxHeight + offset}

		for {
			// Apply jet push
			jet := jetPattern[jetIndex%len(jetPattern)]
			jetIndex++
			if jet == '>' {
				if canMove(rock, rockPos, Point{1, 0}, rockPositions, chamberWidth) {
					rockPos.x++
				}
			} else if jet == '<' {
				if canMove(rock, rockPos, Point{-1, 0}, rockPositions, chamberWidth) {
					rockPos.x--
				}
			} else {
				panic("invalid jet")
			}

			// Apply gravity
			if canMove(rock, rockPos, Point{0, -1}, rockPositions, chamberWidth) {
				rockPos.y--
			} else {
				break
			}
		}

		// Add rock to positions
		for _, p := range rock.shape {
			newPos := Point{rockPos.x + p.x, rockPos.y + p.y}
			rockPositions[newPos] = struct{}{}
			if newPos.y > maxHeight {
				maxHeight = newPos.y
			}
		}

		truncate(rockPositions)
	}

	return maxHeight + 1 // +1 to account for the height starting from 0
}

// only slightly optimizes the processing
func truncate(points map[Point]struct{}) int {
	// we expect that after this amount of rocks no one will reach the bottom
	// so we clean up the map a bit
	if len(points) < 10000 {
		return 0
	}

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

	//fmt.Println("Truncating elements", min)

	return min
}

func getMaxHeight(rockPositions map[Point]struct{}) int {
	maxHeight := 0
	for p := range rockPositions {
		if p.y > maxHeight {
			maxHeight = p.y
		}
	}
	return maxHeight
}

func canMove(rock Rock, pos, delta Point, rockPositions map[Point]struct{}, chamberWidth int) bool {
	for _, p := range rock.shape {
		newPos := Point{pos.x + p.x + delta.x, pos.y + p.y + delta.y}
		if newPos.x < 0 || newPos.x >= chamberWidth || newPos.y < 0 {
			return false
		}
		if _, exists := rockPositions[newPos]; exists {
			return false
		}
	}
	return true
}

func printPos(positions map[Point]struct{}, chamberWidth int) {
	maxHeight := getMaxHeight(positions)
	for y := maxHeight; y >= 0; y-- {
		for x := 0; x < chamberWidth; x++ {
			if _, exists := positions[Point{x, y}]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
