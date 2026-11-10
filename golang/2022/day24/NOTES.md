# 2022 · Day 24: Blizzard Basin

https://adventofcode.com/2022/day/24

## Part 1

**Approach:** BFS by minute: the frontier is the set of cells you can be in at minute `t`; each step tries wait/up/down/left/right and keeps cells that are free at `t+1`. Blizzards are never simulated: each moves independently and wraps inside the walls, so cell `(r, c)` is hit by a `>` blizzard at time `t` iff the initial map has `>` at `(r, (c - t) mod w)` (same for the other three directions). The frontier is a set, so it stays bounded by the valley size (~3500 cells) instead of growing with the number of paths.

**Struggles:** The fetched `example.txt` was the small two-blizzard illustration, not the "more complex example" with the expected answer, so it was extracted by hand. While doing that, found a bug in `aoc fetch`: `<` and `>` in code blocks were decoded and then stripped as HTML tags (`#<^<#` vanished from `puzzle.md`). Fixed in `toMarkdown` with a regression test; `example.txt` files of earlier days were not affected (they are extracted separately).

**Result:** 221. Accepted on the first submission.

## Part 2

**Approach:** Three BFS trips chained by start time: start → goal, goal → start, start → goal. Taking the earliest arrival for each leg is optimal because you can always wait at the start/goal cells (blizzards never reach them), so arriving earlier never hurts the next leg.

**Struggles:** none. The example test checks each leg (18, 23, 13), not only the total 54.

**Result:** 739. Both parts 0.23 s wall, 0.05 s CPU, ~9.7 MB. Accepted on the first submission.
