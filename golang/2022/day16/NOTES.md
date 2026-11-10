# 2022 · Day 16: Proboscidea Volcanium

https://adventofcode.com/2022/day/16

## Part 1

Solved earlier in Java (`java/2022/d16`). Re-implemented in Go; reproduces the accepted answer 2183.

**Approach:** Only 15 of 54 valves have a positive flow rate. Compress the graph: Floyd–Warshall over all valves, keep the distance matrix between AA and the 15 useful valves. A DFS then only chooses "which valve to open next" (travel + 1 minute), adding `remainingTime * rate` immediately when a valve opens. The DFS records `best[openedMask]` = most pressure released by any route that opens exactly that set.

**Struggles:** none, the graph compression keeps the search small (the raw minute-by-minute search over 54 valves would explode).

**Result:** < 0.03 s CPU, ~5.7 MB.

## Part 2

**Approach:** Two actors moving at the same time would square the state space. Instead: you and the elephant open **disjoint** sets of valves, and each actor's best for a given set is independent of the other. So run the same DFS once for 26 minutes to get `best[mask]`, compute `sub[mask]` = max of `best` over all subsets of `mask` (subset-max DP, 15 · 2¹⁵ steps), and take `max(best[m] + sub[full ^ m])`.

**Struggles:** none; the joint search was avoided by design (the known trap of this puzzle).

**Result:** 2911. 0.02 s CPU, ~5.7 MB, single core. Accepted on the first submission.
