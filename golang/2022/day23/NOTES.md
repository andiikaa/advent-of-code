# 2022 · Day 23: Unstable Diffusion

https://adventofcode.com/2022/day/23

## Part 1

**Approach:** Elves as a set of points in a map (the area is unbounded). Each round: every elf with a neighbour proposes the first free direction (rotating start N, S, W, E), proposals are counted per target, and only unique targets are taken. Answer = bounding rectangle area − number of elves after 10 rounds.

**Struggles:** none.

**Result:** 4158. Accepted on the first submission.

## Part 2

**Approach:** Same `round`, which also reports whether any elf moved; loop until a round with no movement.

**Struggles:** none. ~2500 elves for ~1000 rounds with hash maps is fast enough, so no grid/bitset optimisation was needed.

**Result:** 1014. Both parts 0.59 s wall, 0.42 s CPU, ~10.7 MB. Accepted on the first submission.
