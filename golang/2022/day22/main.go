package main

import (
	_ "embed"
	"math"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

// facing: 0 right, 1 down, 2 left, 3 up (as in the puzzle)
var dRow = [4]int{0, 1, 0, -1}
var dCol = [4]int{1, 0, -1, 0}

type board struct {
	grid  []string // padded with spaces to equal width
	moves []int    // >0 steps, -1 turn right, -2 turn left
}

type pos struct{ r, c, f int }

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}

func parse(in string) board {
	parts := strings.SplitN(in, "\n\n", 2)
	lines := strings.Split(parts[0], "\n")
	width := 0
	for _, l := range lines {
		width = max(width, len(l))
	}
	for i, l := range lines {
		lines[i] = l + strings.Repeat(" ", width-len(l))
	}

	var moves []int
	n := 0
	for _, ch := range strings.TrimSpace(parts[1]) {
		switch ch {
		case 'R', 'L':
			if n > 0 {
				moves = append(moves, n)
				n = 0
			}
			if ch == 'R' {
				moves = append(moves, -1)
			} else {
				moves = append(moves, -2)
			}
		default:
			n = n*10 + int(ch-'0')
		}
	}
	if n > 0 {
		moves = append(moves, n)
	}
	return board{grid: lines, moves: moves}
}

func (b board) at(r, c int) byte {
	if r < 0 || r >= len(b.grid) || c < 0 || c >= len(b.grid[0]) {
		return ' '
	}
	return b.grid[r][c]
}

// walk follows the path; next returns the position one step ahead
// (including any wrapping, which may change the facing).
func (b board) walk(next func(p pos) pos) int {
	p := pos{0, strings.IndexByte(b.grid[0], '.'), 0}
	for _, m := range b.moves {
		switch m {
		case -1:
			p.f = (p.f + 1) % 4
		case -2:
			p.f = (p.f + 3) % 4
		default:
			for range m {
				np := next(p)
				if b.at(np.r, np.c) == '#' {
					break
				}
				p = np
			}
		}
	}
	return 1000*(p.r+1) + 4*(p.c+1) + p.f
}

// part1: off the board, wrap to the opposite side of the same row/column.
func part1(in string) int {
	b := parse(in)
	rows, cols := len(b.grid), len(b.grid[0])
	return b.walk(func(p pos) pos {
		r, c := p.r, p.c
		for {
			r = (r + dRow[p.f] + rows) % rows
			c = (c + dCol[p.f] + cols) % cols
			if b.at(r, c) != ' ' {
				return pos{r, c, p.f}
			}
		}
	})
}

// ---------------------------------------------------------------- cube

type vec struct{ x, y, z int }

func (a vec) dot(b vec) int { return a.x*b.x + a.y*b.y + a.z*b.z }
func (a vec) neg() vec      { return vec{-a.x, -a.y, -a.z} }
func (a vec) add(b vec) vec { return vec{a.x + b.x, a.y + b.y, a.z + b.z} }
func (a vec) mul(k int) vec { return vec{a.x * k, a.y * k, a.z * k} }

// frame places a face of the net on the cube: n points out of the face,
// right is the 3D direction of increasing column, down of increasing row.
type frame struct {
	fr, fc         int // face position in the net (in units of faces)
	n, right, down vec
}

// dir returns the 3D direction of 2D facing f on this face.
func (fr frame) dir(f int) vec {
	switch f {
	case 0:
		return fr.right
	case 1:
		return fr.down
	case 2:
		return fr.right.neg()
	}
	return fr.down.neg()
}

// fold assigns a 3D frame to every face by walking the net: stepping onto
// the face to the right tilts the cube so the old "right" becomes the new
// normal, and so on for the other directions. This works for any net, so
// the example (4x4 faces) and the input (50x50) share the same code instead
// of hand-written edge tables.
func fold(b board, size int) map[[2]int]frame {
	start := [2]int{0, strings.IndexByte(b.grid[0], '.') / size}
	faces := map[[2]int]frame{start: {start[0], start[1], vec{0, 0, 1}, vec{1, 0, 0}, vec{0, 1, 0}}}
	queue := [][2]int{start}
	for len(queue) > 0 {
		k := queue[0]
		queue = queue[1:]
		f := faces[k]
		next := []frame{
			{f.fr, f.fc + 1, f.right, f.n.neg(), f.down}, // right
			{f.fr + 1, f.fc, f.down, f.right, f.n.neg()}, // down
			{f.fr, f.fc - 1, f.right.neg(), f.n, f.down}, // left
			{f.fr - 1, f.fc, f.down.neg(), f.right, f.n}, // up
		}
		for _, nf := range next {
			nk := [2]int{nf.fr, nf.fc}
			if _, ok := faces[nk]; ok || b.at(nf.fr*size, nf.fc*size) == ' ' {
				continue
			}
			faces[nk] = nf
			queue = append(queue, nk)
		}
	}
	if len(faces) != 6 {
		panic("net does not have 6 faces")
	}
	return faces
}

// part2: the board is folded into a cube. Points are tracked in doubled 3D
// coordinates (cube spans -size..size) so tile centers are integers.
// Leaving face F in 3D direction v lands on the face whose normal is v; the
// new point keeps its position along the shared edge, the new facing is -n(F).
func part2(in string) int {
	b := parse(in)
	tiles := 0
	for _, l := range b.grid {
		tiles += len(strings.ReplaceAll(l, " ", ""))
	}
	size := int(math.Sqrt(float64(tiles / 6)))
	faces := fold(b, size)
	byNormal := map[vec]frame{}
	for _, f := range faces {
		byNormal[f.n] = f
	}

	return b.walk(cubeNext(b, size, faces, byNormal))
}

// cubeNext returns the step function for walking on the folded cube.
func cubeNext(b board, size int, faces map[[2]int]frame, byNormal map[vec]frame) func(p pos) pos {
	return func(p pos) pos {
		r, c := p.r+dRow[p.f], p.c+dCol[p.f]
		if b.at(r, c) != ' ' {
			return pos{r, c, p.f}
		}
		f := faces[[2]int{p.r / size, p.c / size}]
		i, j := p.r%size, p.c%size
		// 3D point of the current tile (doubled coordinates)
		pt := f.n.mul(size).add(f.right.mul(2*j - (size - 1))).add(f.down.mul(2*i - (size - 1)))
		v := f.dir(p.f)
		g := byNormal[v]
		// set component along v to size (new face), along old normal to size-1 (edge row)
		pt = pt.add(v.mul(size - pt.dot(v))).add(f.n.mul(size - 1 - pt.dot(f.n)))
		ni := (pt.dot(g.down) + size - 1) / 2
		nj := (pt.dot(g.right) + size - 1) / 2
		nd := f.n.neg()
		nf := 0
		for k := range 4 {
			if g.dir(k) == nd {
				nf = k
			}
		}
		return pos{g.fr*size + ni, g.fc*size + nj, nf}
	}
}
