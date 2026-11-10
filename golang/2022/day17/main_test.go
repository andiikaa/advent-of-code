package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != 3068 {
			t.Errorf("Got %d, expected %d", got, 3068)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(example); got != 1514285714288 {
			t.Errorf("Got %d, expected %d", got, 1514285714288)
		}
	})
}

func TestPart1(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		expected := 3200
		p1 := part1(input)
		if p1 != expected {
			t.Errorf("Got %d, expected %d", p1, expected)
		}
	})

	t.Run("part2", func(t *testing.T) {
		expected := 1584927536247
		p2 := part2(input)
		if p2 != expected {
			t.Errorf("Got %d, expected %d", p2, expected)
		}
	})
}
