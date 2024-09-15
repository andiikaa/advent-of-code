package main

import "testing"

func Test(t *testing.T) {
	seeds := seeds(input)
	parsed := parseMaps(input)

	t.Run("part1", func(t *testing.T) {
		expected := 462648396
		p1 := part1(seeds, parsed)
		if p1 != expected {
			t.Errorf("Got %d, expected %d", p1, expected)
		}
	})

	t.Run("part2", func(t *testing.T) {
		expected := 2520479
		p2 := part2(seeds, parsed)
		if p2 != expected {
			t.Errorf("Got %d, expected %d", p2, expected)
		}
	})
}
