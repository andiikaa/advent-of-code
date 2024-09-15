package utils

import "testing"

func TestOverlap(t *testing.T) {
	r1 := Range{1, 5}
	r2 := Range{3, 7}
	r3 := Range{5, 7}

	if !r1.Overlaps(&r2) {
		t.Errorf("Expected overlap")
	}

	if !r2.Overlaps(&r1) {
		t.Errorf("Expected overlap")
	}

	if !r1.Overlaps(&r3) {
		t.Errorf("Expected overlap")
	}

	if !r3.Overlaps(&r2) {
		t.Errorf("Expected overlap")
	}
}

func TestMerge(t *testing.T) {
	r1 := Range{1, 5}
	r2 := Range{3, 7}
	r3 := Range{5, 7}

	merged, ok := r1.Merge(&r2)
	if !ok {
		t.Errorf("Expected merge")
	}

	if merged.Start != 1 {
		t.Errorf("Expected start to be 1")
	}

	if merged.End != 7 {
		t.Errorf("Expected end to be 7")
	}

	merged, ok = r2.Merge(&r3)
	if !ok {
		t.Errorf("Expected merge")
	}

	if merged.Start != 3 {
		t.Errorf("Expected start to be 3")
	}

	if merged.End != 7 {
		t.Errorf("Expected end to be 7")
	}

	merged, ok = r1.Merge(&r3)
	if !ok {
		t.Errorf("Expected merge")
	}

	if merged.Start != 1 {
		t.Errorf("Expected start to be 1")
	}

	if merged.End != 7 {
		t.Errorf("Expected end to be 7")
	}

	merged, ok = r3.Merge(&r1)
	if !ok {
		t.Errorf("Expected merge")
	}

	if merged.Start != 1 {
		t.Errorf("Expected start to be 1")
	}

	if merged.End != 7 {
		t.Errorf("Expected end to be 7")
	}
}
