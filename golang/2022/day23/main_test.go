package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != 110 {
			t.Errorf("Got %d, expected %d", got, 110)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(example); got != 20 {
			t.Errorf("Got %d, expected %d", got, 20)
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input); got != 4158 {
			t.Errorf("Got %d, expected %d", got, 4158)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(input); got != 1014 {
			t.Errorf("Got %d, expected %d", got, 1014)
		}
	})
}
