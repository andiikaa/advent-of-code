# 2022 · Day 15: Beacon Exclusion Zone

https://adventofcode.com/2022/day/15

## Part 1

Solved earlier in Java (`java/2022/d15`). Re-implemented in Go; reproduces the accepted answer 5073496.

**Approach:** Each sensor covers a diamond of radius `r` (Manhattan distance to its beacon). On row `y` that is the interval `x ± (r - |sy - y|)`. Sort and merge the intervals, sum their lengths, subtract beacons lying on the row. O(n log n) independent of the coordinate size (values are in the millions, so a cell set would be slow and memory-hungry).

**Struggles:** none. The row number differs between example (10) and input (2000000), so it is a parameter.

**Result:** < 0.01 s, ~4.5 MB.

## Part 2

**Approach:** Brute force over the 4,000,001² area would be ~1.6·10¹³ cell checks (not attempted). Since exactly one cell is uncovered, it must touch the border of some diamond, i.e. lie at distance `r+1` from a sensor. Those borders are the diagonal lines `y = x + a` and `y = -x + b`. Intersect every `a` line with every `b` line (n = 30 sensors → 3600 candidates), keep those inside the area, and return the first one no sensor covers. Tuning frequency `x*4000000 + y`.

**Struggles:** none in practice; the scaling trap was visible from the numbers before writing code. Alternative that also works: the part 1 interval merge for each of the 4M rows (~1–2 s).

**Result:** 13081194638237 (x=3270298, y=2638237). < 0.01 s, ~4.5 MB. Accepted on the first submission.
