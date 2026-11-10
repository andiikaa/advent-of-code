package main

import "testing"

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != 18 {
			t.Errorf("Got %d, expected %d", got, 18)
		}
	})

	t.Run("part2", func(t *testing.T) {
		v := parse(example)
		t1 := v.trip(v.start, v.goal, 0)
		t2 := v.trip(v.goal, v.start, t1)
		t3 := v.trip(v.start, v.goal, t2)
		if t1 != 18 || t2-t1 != 23 || t3-t2 != 13 {
			t.Errorf("legs %d, %d, %d; expected 18, 23, 13", t1, t2-t1, t3-t2)
		}
		if got := part2(example); got != 54 {
			t.Errorf("Got %d, expected %d", got, 54)
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input); got != 221 {
			t.Errorf("Got %d, expected %d", got, 221)
		}
	})

	t.Run("part2", func(t *testing.T) {
		if got := part2(input); got != 739 {
			t.Errorf("Got %d, expected %d", got, 739)
		}
	})
}
