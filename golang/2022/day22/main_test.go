package main

import "testing"

func TestPart1(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		expected := 6032
		actual := part1()
		if actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("part2", func(t *testing.T) {

	})
}
