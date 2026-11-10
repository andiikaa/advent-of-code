package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example, 10); got != 26 {
			t.Errorf("Got %d, expected %d", got, 26)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(example, 20); got != 56000011 {
			t.Errorf("Got %d, expected %d", got, 56000011)
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input, 2000000); got != 5073496 {
			t.Errorf("Got %d, expected %d", got, 5073496)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(input, 4000000); got != 13081194638237 {
			t.Errorf("Got %d, expected %d", got, 13081194638237)
		}
	})
}
