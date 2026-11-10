package main

import "testing"

const exampleScreen = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != 13140 {
			t.Errorf("Got %d, expected %d", got, 13140)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(example); got != exampleScreen {
			t.Errorf("Got\n%s\nexpected\n%s", got, exampleScreen)
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input); got != 17940 {
			t.Errorf("Got %d, expected %d", got, 17940)
		}
	})

	// Accepted answer: ZCBAJFJZ (read from the rendered screen).
	t.Run("part2", func(t *testing.T) {
		expected := `####..##..###...##....##.####...##.####.
...#.#..#.#..#.#..#....#.#.......#....#.
..#..#....###..#..#....#.###.....#...#..
.#...#....#..#.####....#.#.......#..#...
#....#..#.#..#.#..#.#..#.#....#..#.#....
####..##..###..#..#..##..#.....##..####.`
		if got := part2(input); got != expected {
			t.Errorf("Got\n%s\nexpected\n%s", got, expected)
		}
	})
}
