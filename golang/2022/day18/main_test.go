package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		expected := 4504
		p1 := part1(input)
		if p1 != expected {
			t.Errorf("Got %d, expected %d", p1, expected)
		}
	})

	t.Run("part2", func(t *testing.T) {
		expected := 4504
		p2 := part2(input)
		if p2 != expected {
			t.Errorf("Got %d, expected %d", p2, expected)
		}
	})
}
