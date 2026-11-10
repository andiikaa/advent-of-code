# 2022 · Day 19: Not Enough Minerals

https://adventofcode.com/2022/day/19

## Part 1

**Approach:** DFS per blueprint. Instead of branching every minute (build ore/clay/obsidian/geode robot or wait → 5^24 paths), each step chooses **which robot to build next** and jumps forward to the minute it gets built (wait for resources, +1 minute to build). A geode robot built with `t` minutes left is credited `t` geodes immediately, so the state needs no geode-robot count.

Pruning:
- Never own more robots of a kind than the most of that resource any recipe needs per minute (the factory builds one robot per minute, so extra production can't be spent).
- Upper bound: even building one geode robot every remaining minute only adds `t(t-1)/2`; cut the branch if that can't beat the best found so far.

Parsing collects all numbers and takes 7 per blueprint, so the multi-line example and the one-line input share one parser.

**Struggles:** none. The "jump to next build" formulation plus the two pruning rules was enough from the start; no naive per-minute version was tried.

**Result:** 1389. 0.01 s CPU, ~4.6 MB. Accepted on the first submission.

## Part 2

**Approach:** Same DFS, 32 minutes, first three blueprints, product of the geode counts. The longer horizon is where per-minute searches usually blow up; the bound prune keeps it fast.

**Struggles:** none.

**Result:** 3003. Whole program (both parts) 0.20 s wall, 0.01 s CPU, ~4.6 MB. Accepted on the first submission.
