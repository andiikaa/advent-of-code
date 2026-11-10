# 2022 · Day 25: Full of Hot Air

https://adventofcode.com/2022/day/25

## Part 1

**Approach:** SNAFU is base 5 with digits −2…2 (`=`, `-`, `0`, `1`, `2`). Parsing: `n = n*5 + digit`. Converting back: take `n mod 5`; digits 3 and 4 are written as `=` and `-` with a carry of 1 into the next place (3 = 5 − 2, 4 = 5 − 1). The answer is a string.

**Struggles:** none. Tested against the full conversion table from the puzzle in both directions.

**Result:** `2----0=--1122=0=0021`. < 0.01 s. Accepted on the first submission.

## Part 2

No puzzle: with the other 49 stars, the "Start the Blender" button posts answer `0` for level 2. `aoc submit` normally refuses `0`, so day 25 part 2 is whitelisted. The helper didn't recognise the final congratulation text as a verdict, so the log line was added by hand.
