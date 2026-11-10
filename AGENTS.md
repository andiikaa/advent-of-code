# AGENTS.md

Advent of Code solutions. Older years are in `java/` and `python/`; new work is in **Go** under `golang/` (module `github.com/andiikaa/advent-of-code`, Go stdlib only).

Current goal: finish **2022**. Run `go run ./cmd/aoc status 2022` to see which days are open:
- `*` (silver): part 1 was solved earlier (often in Java under `java/2022/`); only part 2 is missing. Solve both parts in Go anyway: part 2 builds on part 1, and the part 1 answer (shown in `puzzle.md` as "Your puzzle answer was …") is a free check of the Go code. Do not submit part 1 again.
- no star: solve both parts.
- Day 25 part 2 has no puzzle: once all other 49 stars are collected, submit `aoc submit <year> 25 2 0` (that is what the "Start the Blender" button posts).

## Layout

```
golang/
  cmd/aoc/            helper: login, status, fetch, submit
  2022/dayNN/
    main.go           package main; //go:embed input.txt and example.txt; part1(in string) int, part2(in string) int
    main_test.go      TestExample (example answers) + TestInput (accepted answers, regression)
    example.txt       example from the puzzle text
    input.txt         personal puzzle input (no trailing newline)
    puzzle.md         puzzle text (gitignored, refetched on every fetch)
    submissions.log   every submitted answer and its verdict (committed, used by the submit guard)
    NOTES.md          write-up: approach and struggles per part (committed)
```

Follow the style of the existing days (`golang/2022/day18` is a good reference): one `main` package per day, plain functions, no shared framework. Small helpers may live in the day's package. Return types other than `int` are fine when the answer is not a number (e.g. a string).

## Helper (`golang/cmd/aoc`)

Run from `golang/`:

```sh
go run ./cmd/aoc status 2022              # stars per day, list of open days
go run ./cmd/aoc fetch 2022 19            # creates 2022/day19, writes puzzle.md, example.txt, input.txt, main.go/main_test.go skeletons
go run ./cmd/aoc submit 2022 19 1 1234    # submit answer for part 1
```

- Auth: session cookie (just the value) from `~/.config/aoc/session`, shared by all worktrees, or `$AOC_SESSION`. `go run ./cmd/aoc login` copies it from Safari to `~/.config/aoc/session`, which only works if the terminal app has Full Disk Access. Never print, log or commit it. If requests fail with a redirect / 400 / 500, the cookie probably expired: ask the user for a new one.
- `fetch` never overwrites existing files except `puzzle.md`. Run it again after solving part 1 to get the part 2 text.
- `example.txt` is the first code block after "For example". Check it against `puzzle.md`; some puzzles have several examples or the relevant one is further down. Fix it by hand if needed (or add `example2.txt`).
- The input is downloaded only once. Don't fetch it again or loop over the site; AoC asks automated tools to keep traffic low.

## Workflow per day

1. `fetch`, then read `puzzle.md` carefully: rules, edge cases, and the **expected example answer**.
2. TDD on the example: put the expected example answer in `TestExample`, implement until `go test ./2022/dayNN/` passes. Add more small tests for tricky pieces if it helps.
3. Run on the real input: `go run ./2022/dayNN/`. It should finish within a few seconds. If it doesn't, the approach is wrong; don't wait it out (see below).
4. Sanity check the result (plausible size, no overflow, not 0, not equal to the example answer).
5. Submit with `submit`. Then record the accepted answer in `TestInput` (remove the `t.Skip`).
6. Update `NOTES.md` for the part (see below).
7. Re-fetch for part 2 and repeat steps 1–6 for part 2.
8. Commit the day (`dayNN: part 1` / `dayNN: part 2`).

## Measuring and documenting

Measure the real-input run on a built binary, so compile time doesn't count:

```sh
go build -o /tmp/aoc-dayNN ./2022/dayNN && /usr/bin/time -l /tmp/aoc-dayNN
```

`real`/`user`/`sys` give wall time and CPU (user+sys > real means more than one core, e.g. "CPU 200%"), and `maximum resident set size` is peak memory in bytes. Never let a run go unbounded: run attempts in the background or with `go test -timeout 60s`, and kill anything that passes ~60 s or several GB of memory, then write down how far it got.

`NOTES.md` (skeleton created by `fetch`) is the honest log of the day, one section per part:
- **Approach**: the idea and algorithm in a few sentences, with complexity where useful.
- **Struggles**: every attempt that failed or didn't scale, with numbers: e.g. "naive BFS over all states: killed after 60 s at 9 GB RSS, CPU 200%", "OOM in part 2 because the grid was stored as a map of cells", "wrong answer 1234 (too low): off-by-one in the range merge". Also misreadings of the puzzle and wrong submissions with their cause. Write "none" if it went straight through.
- **Result**: final runtime, peak memory, CPU of the accepted solution.

## Submitting: rate limit

Wrong answers trigger an increasing cooldown (1 min, then 5 min, …). Treat every submission as expensive:
- Only submit after the example test passes **and** the real-input run is sanity checked. Never guess or try variations.
- `submit` refuses: an already solved part, an answer submitted before, answers outside known too-high/too-low bounds, and submissions during an active cooldown (`golang/.aoc/cooldown`). Do not work around these guards.
- If an answer is wrong: stop, re-read the puzzle, look for the misunderstanding, add a test for the case you missed. Don't submit again until you've found a concrete bug.
- If a part gets two wrong answers, stop and report to the user instead of continuing.

## Performance

Part 2 is often the same problem at a scale where the naive solution explodes (exponential search, huge iteration counts, giant grids). Typical fixes:
- Shortest paths: BFS (unweighted), Dijkstra (`container/heap`), A* with an admissible heuristic.
- Search over states: memoization keyed on a compact state, pruning with upper bounds (branch and bound), bitmasks for sets of visited things.
- Huge iteration counts (e.g. "after 1000000000000 rounds"): detect a cycle in the state, then skip ahead arithmetically.
- Large coordinates: work with ranges/intervals instead of cells, or compress coordinates.
- Modular arithmetic (LCM / product of moduli) to keep numbers small.
Use `int` (64-bit). Note the chosen trick in a short comment so it's clear why the code looks the way it does.
