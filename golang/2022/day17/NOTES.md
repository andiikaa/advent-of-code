# 2022 · Day 17: Pyroclastic Flow

https://adventofcode.com/2022/day/17

## Part 1

Solved by hand in Go (accepted 3200). Kept as is: rocks are point sets, `canMove` checks walls/floor/settled rocks in a map, old rows are pruned from the map.

## Part 2

**Approach:** 10¹² rocks can't be simulated. Between two rocks, the future only depends on (rock shape index, jet index, shape of the surface). The surface is captured as the depth of each column's top below the tower top. When such a state repeats, the rocks in between form a period: skip `(N - i) / period` whole periods, add `periods × heightPerPeriod`, and simulate only the remaining rocks.

**Struggles:**
- The earlier hand-written version already had a cycle map, but on a hit it only added the height and `continue`d: the loop still ran 10¹² iterations (would never finish), and the skipped rock didn't advance the jet index, so the simulation went out of sync after the first repeat. The test expected 4504 for part 2, which is a copy-paste of day 18's answer.
- Surface profiles are a heuristic (two states can share a profile but differ deeper down). Verified by running the input with and without skipping for 5000 … 300000 rocks: identical heights in all cases.
- The map pruning ran over the whole map after every rock and used a margin of 4 rows below the lowest column top, which is too tight for rocks sliding under overhangs. Now runs every 1000 rocks with a 50-row margin.

**Result:** 1584927536247. < 0.01 s CPU, ~6.7 MB (the cycle shows up after a few thousand rocks). Accepted on the first submission.
