package main

import "testing"

func TestPart1(t *testing.T) {
	expected := 54990
	p1 := part1(input)
	if p1 != expected {
		t.Errorf("Got %d, expected %d", p1, expected)
	}

}
