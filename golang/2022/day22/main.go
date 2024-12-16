package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

type facing int
type cellType int

const (
	// facing
	faceUp    = facing(0)
	faceRight = facing(1)
	faceDown  = facing(2)
	faceLeft  = facing(3)

	// type of cell
	emptyCell = cellType(0)
	wallCell  = cellType(1)
	tileCell  = cellType(2)
)

func main() {
	part1()
}

func part1() int {
	//colMax := 150
	// colMax := 16
	in := strings.Split(example, "\n\n")
	// '-1' for ' ', '0' for '.', (move) '1' for '#' (wall)
	boardMap := toArray(in[0])
	// -1 for R, -2 for L, positive number for number of steps
	path := parseMoves(in[1])
	f := faceRight

	// moves
	for _, p := range path {
		col := 0
		row := 0

		if p == -1 {
			// turn right
			f = facing((f + 1) % 4)
		} else if p == -2 {
			// turn left
			f = facing((f - 1) % 4)
		} else {
			// move forward
			if f == faceUp {
				row -= p
			} else if f == faceRight {
				for i := 0; i < p; i++ {
					col++

					// jump over empty cells
					for {
						if len(boardMap[row]) <= col {
							col = 0
						}
						if boardMap[row][col] == emptyCell {
							col++
						} else {
							break
						}
					}

					if boardMap[row][col] == wallCell {
						fmt.Println("hit wall at", row, col)
					} else if boardMap[row][col] == tileCell {
						fmt.Println("found tile at", row, col)
					}
				}
			} else if f == faceDown {
				row += p
			} else if f == faceLeft {
				col -= p
			} else {
				panic("invalid facing")
			}
		}
	}
	return 0
}

func toArray(board string) [][]cellType {
	row := strings.Split(board, "\n")
	intBoard := [][]cellType{}
	for _, b := range row {
		e := []cellType{}
		for _, c := range b {
			if c == '#' {
				e = append(e, wallCell)
			} else if c == '.' {
				e = append(e, tileCell)
			} else {
				e = append(e, emptyCell)
			}
		}
		intBoard = append(intBoard, e)
	}
	return intBoard
}

func parseMoves(path string) []int {
	out := make([]int, 0)
	for i := 0; i < len(path); i++ {
		c := string(path[i])
		if c == "R" {
			out = append(out, -1)
		} else if c == "L" {
			out = append(out, -2)
		} else {
			lastInt := -1
			check := c
			for {
				if v, err := strconv.Atoi(check); err != nil {
					if lastInt == -1 {
						panic("failed to parse input")
					}
					i--
					out = append(out, lastInt)
					break
				} else {
					lastInt = v
				}
				i++
				if i >= len(path) {
					out = append(out, lastInt)
					break
				}
				check = check + string(path[i])
			}
		}
	}
	return out
}
