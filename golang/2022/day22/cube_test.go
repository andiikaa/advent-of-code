package main

import (
	"math"
	"strings"
	"testing"
)

// cubeStep builds the part 2 step function for a board.
func cubeStep(in string) (board, func(p pos) pos) {
	b := parse(in)
	tiles := 0
	for _, l := range b.grid {
		tiles += len(strings.ReplaceAll(l, " ", ""))
	}
	size := int(math.Sqrt(float64(tiles / 6)))
	faces := fold(b, size)
	byNormal := map[vec]frame{}
	for _, f := range faces {
		byNormal[f.n] = f
	}
	return b, cubeNext(b, size, faces, byNormal)
}

// TestCubeWrapRoundTrip: stepping off any edge tile and stepping back
// (reversed facing) must return to the same tile, on example and input.
func TestCubeWrapRoundTrip(t *testing.T) {
	for name, in := range map[string]string{"example": example, "input": input} {
		b, step := cubeStep(in)
		checked := 0
		for r, row := range b.grid {
			for c := range row {
				if b.at(r, c) == ' ' {
					continue
				}
				for f := range 4 {
					if b.at(r+dRow[f], c+dCol[f]) != ' ' {
						continue
					}
					np := step(pos{r, c, f})
					if b.at(np.r, np.c) == ' ' {
						t.Fatalf("%s: (%d,%d,%d) wrapped off the board to %v", name, r, c, f, np)
					}
					back := step(pos{np.r, np.c, (np.f + 2) % 4})
					if back.r != r || back.c != c || back.f != (f+2)%4 {
						t.Fatalf("%s: (%d,%d,%d) -> %v -> %v", name, r, c, f, np, back)
					}
					checked++
				}
			}
		}
		t.Logf("%s: %d edge steps round-trip", name, checked)
	}
}
