package utils

type Range struct {
	Start int
	End   int
}

func (r *Range) Contains(i int) bool {
	return i >= r.Start && i <= r.End
}

func (r *Range) ContainsOther(other *Range) bool {
	return r.Start <= other.Start && r.End >= other.End
}

// inclusive overlap
// Range{1, 5} overlaps Range{5, 7}
func (r *Range) Overlaps(other *Range) bool {
	return r.Start <= other.End && r.End >= other.Start
}

// merges two ranges together into one if they overlap
func (r *Range) Merge(other *Range) (Range, bool) {
	if r.Overlaps(other) {
		newRange := Range{
			Start: r.Start,
			End:   r.End,
		}

		if r.Start > other.Start {
			newRange.Start = other.Start
		}

		if r.End < other.End {
			newRange.End = other.End
		}

		return newRange, true
	}
	return *r, false
}
