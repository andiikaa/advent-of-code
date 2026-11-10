# 2022 · Day 22: Monkey Map

https://adventofcode.com/2022/day/22

## Part 1

**Approach:** Pad all map rows to the same width, parse the path into steps and turns. A shared `walk` follows the path and asks a `next(pos)` function for the tile ahead; walls stop the current move. Flat wrapping: keep stepping in the facing direction modulo the board size until a non-space tile is hit.

**Struggles:** A hand-written start existed (`cc504f8 basic move right`, returned 0 for the example). Replaced with a full implementation rather than extended, the old one is in git history.

**Result:** 56372. < 0.01 s, ~4.7 MB. Accepted on the first submission.

## Part 2

**Approach:** Fold the net into a cube instead of hard-coding edge pairs (the example and the input have different nets, so a hand-written edge table would be needed twice and is very error-prone). BFS over the faces of the net assigns each face a 3D frame (normal, "right", "down"); stepping onto a neighbour face tilts the frame (e.g. going right: new normal = old right, new right = −old normal). Positions are 3D points in doubled coordinates so tile centres are integers. Leaving face F in 3D direction `v` lands on the face whose normal is `v`: set the point's `v`-component to the cube surface and its old-normal component to the edge row, convert back with the new face's frame; the new facing is −normal(F). Face size = √(tiles/6), so the same code handles 4×4 and 50×50.

**Struggles:** Wrong edge mappings are the classic failure here and the example doesn't exercise all edges. Added `TestCubeWrapRoundTrip`: step off every boundary tile in every outward direction and step back with reversed facing; must return to the start. 56 edge steps (example) and 700 (input, 14 edges × 50) all round-trip before submitting.

**Result:** 197047. < 0.01 s, ~4.7 MB. Accepted on the first submission.
