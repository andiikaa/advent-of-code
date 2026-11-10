package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != 1651 {
			t.Errorf("Got %d, expected %d", got, 1651)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(example); got != 1707 {
			t.Errorf("Got %d, expected %d", got, 1707)
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input); got != 2183 {
			t.Errorf("Got %d, expected %d", got, 2183)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(input); got != 2911 {
			t.Errorf("Got %d, expected %d", got, 2911)
		}
	})
}
