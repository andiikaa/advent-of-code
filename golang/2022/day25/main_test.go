package main

import "testing"

func TestConversion(t *testing.T) {
	// table from the puzzle
	table := map[int]string{
		1: "1", 2: "2", 3: "1=", 4: "1-", 5: "10", 6: "11", 7: "12", 8: "2=", 9: "2-", 10: "20",
		15: "1=0", 20: "1-0", 2022: "1=11-2", 12345: "1-0---0", 314159265: "1121-1110-1=0",
	}
	for dec, snafu := range table {
		if got := toSNAFU(dec); got != snafu {
			t.Errorf("toSNAFU(%d) = %s, expected %s", dec, got, snafu)
		}
		if got := fromSNAFU(snafu); got != dec {
			t.Errorf("fromSNAFU(%s) = %d, expected %d", snafu, got, dec)
		}
	}
}

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(example); got != "2=-1=0" {
			t.Errorf("Got %s, expected %s", got, "2=-1=0")
		}
	})
}

// TestInput locks in the accepted answers.
func TestInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		if got := part1(input); got != "2----0=--1122=0=0021" {
			t.Errorf("Got %s, expected %s", got, "2----0=--1122=0=0021")
		}
	})
}
