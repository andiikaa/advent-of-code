package main

import "testing"

func TestExample(t *testing.T) {
	bps := parse(example)
	t.Run("blueprints", func(t *testing.T) {
		for i, want := range []int{9, 12} {
			if got := maxGeodes(bps[i], 24); got != want {
				t.Errorf("blueprint %d: got %d, expected %d", i+1, got, want)
			}
		}
	})

	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != 33 {
			t.Errorf("Got %d, expected %d", got, 33)
		}
	})

	t.Run("blueprints 32 minutes", func(t *testing.T) {
		for i, want := range []int{56, 62} {
			if got := maxGeodes(bps[i], 32); got != want {
				t.Errorf("blueprint %d: got %d, expected %d", i+1, got, want)
			}
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(example); got != 56*62 {
			t.Errorf("Got %d, expected %d", got, 56*62)
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input); got != 1389 {
			t.Errorf("Got %d, expected %d", got, 1389)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(input); got != 3003 {
			t.Errorf("Got %d, expected %d", got, 3003)
		}
	})
}
