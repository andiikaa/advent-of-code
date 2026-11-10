# 2022 · Day 9: Rope Bridge

https://adventofcode.com/2022/day/9

## Part 1

Solved earlier in Java (`java/2022/d9`). Re-implemented in Go as the base for part 2; it reproduces the accepted answer 6311.

**Approach:** Simulate head and tail step by step; the tail moves one step towards the head (`sign(dx), sign(dy)`) whenever they stop touching. Count distinct tail positions in a set.

**Struggles:** none.

**Result:** < 0.01 s CPU, ~5 MB peak memory, single core.

## Part 2

**Approach:** Same simulation with 10 knots, each following the one in front. Because the follow rule uses `sign()` on both axes, the new case that only appears with longer ropes (a knot two steps away diagonally) needs no special handling.

**Struggles:** The fetched `example.txt` was the first code block of the puzzle (an illustration, not the motions), so both examples had to be written by hand. Led to an improvement in `aoc fetch` (prefer the block after "For example"). Accepted on the first submission.

**Result:** 2482. < 0.01 s CPU, ~5 MB peak memory, single core.
