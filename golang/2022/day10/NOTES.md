# 2022 · Day 10: Cathode-Ray Tube

https://adventofcode.com/2022/day/10

## Part 1

Solved earlier in Java (`java/2022/d10`). Re-implemented in Go; reproduces the accepted answer 17940.

**Approach:** A small `run` function walks the program and calls a callback once per cycle with the current `X` (`addx` takes two cycles, `noop` one). Part 1 sums `cycle * X` at cycles 20, 60, 100, ...

**Struggles:** none.

**Result:** < 0.01 s, ~5 MB.

## Part 2

**Approach:** Same `run` callback: during cycle `c` the CRT draws column `(c-1) % 40`; the pixel is lit if it is within ±1 of `X`. The answer is the 8 capital letters on the rendered 40x6 screen, read by eye: **ZCBAJFJZ**.

**Struggles:** The answer isn't a number, so `part2` returns the screen as a string; the example test compares the full rendered screen. Reading letters is the only manual step; rendering `.` as spaces makes them much easier to read. Accepted on the first submission.

**Result:** < 0.01 s, ~5 MB.
