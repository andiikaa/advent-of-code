package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != 13 {
			t.Errorf("Got %d, expected %d", got, 13)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(example); got != 1 {
			t.Errorf("Got %d, expected %d", got, 1)
		}
		if got := part2(example2); got != 36 {
			t.Errorf("Got %d, expected %d", got, 36)
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input); got != 6311 {
			t.Errorf("Got %d, expected %d", got, 6311)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(input); got != 2482 {
			t.Errorf("Got %d, expected %d", got, 2482)
		}
	})
}
