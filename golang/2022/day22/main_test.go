package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != 6032 {
			t.Errorf("Got %d, expected %d", got, 6032)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(example); got != 5031 {
			t.Errorf("Got %d, expected %d", got, 5031)
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input); got != 56372 {
			t.Errorf("Got %d, expected %d", got, 56372)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(input); got != 197047 {
			t.Errorf("Got %d, expected %d", got, 197047)
		}
	})
}
