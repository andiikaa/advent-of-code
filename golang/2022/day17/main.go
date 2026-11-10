package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

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

// state identifies a moment between two rocks. If the same state is seen
// twice, everything that follows repeats with the same period.
type state struct {
	rockNum  int
	jetIndex int
	profile  [7]int // depth of each column's top below the tower top
}

type seen struct {
	rock   int // number of rocks dropped when the state was seen
	height int
}

// simulateFallingRocks drops numRocks rocks and returns the tower height.
// 10^12 rocks can't be simulated one by one, so once a state repeats, whole
// periods are skipped arithmetically and only the remainder is simulated.
func simulateFallingRocks(jetPattern string, numRocks int) int {
	jetPattern = strings.TrimSpace(jetPattern)
	jetIndex := 0
	top := -1 // y of the highest rock, -1 = floor only
	colMax := [7]int{-1, -1, -1, -1, -1, -1, -1}
	rockPositions := make(map[Point]struct{})
	states := make(map[state]seen)
	skipped := 0 // height added by skipping periods

	for i := 0; i < numRocks; i++ {
		rockNum := i % len(rockShapes)

		if skipped == 0 {
			st := state{rockNum: rockNum, jetIndex: jetIndex, profile: profile(colMax, top)}
			if prev, ok := states[st]; ok {
				period := i - prev.rock
				periods := (numRocks - i) / period
				skipped = periods * (top - prev.height)
				i += periods * period
				if i >= numRocks {
					break
				}
			} else {
				states[st] = seen{rock: i, height: top}
			}
		}

		rock := rockShapes[rockNum]
		rockPos := Point{2, top + 4}
		for {
			// Apply jet push
			jet := jetPattern[jetIndex]
			jetIndex = (jetIndex + 1) % len(jetPattern)
			switch jet {
			case '>':
				if canMove(rock, rockPos, Point{1, 0}, rockPositions) {
					rockPos.x++
				}
			case '<':
				if canMove(rock, rockPos, Point{-1, 0}, rockPositions) {
					rockPos.x--
				}
			default:
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
			top = max(top, newPos.y)
			colMax[newPos.x] = max(colMax[newPos.x], newPos.y)
		}

		// keep the map small: rows far below the lowest column top can no
		// longer be reached by a falling rock
		if i%1000 == 0 {
			truncate(rockPositions, colMax)
		}
	}

	return top + 1 + skipped
}

// profile normalizes the column tops relative to the tower top, so the same
// surface shape at different heights compares equal.
func profile(colMax [7]int, top int) [7]int {
	out := [7]int{}
	for i, v := range colMax {
		out[i] = top - v
	}
	return out
}

func truncate(points map[Point]struct{}, colMax [7]int) {
	low := colMax[0]
	for _, v := range colMax {
		low = min(low, v)
	}
	// margin for rocks that slide under overhangs
	low -= 50
	for p := range points {
		if p.y < low {
			delete(points, p)
		}
	}
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
