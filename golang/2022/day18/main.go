package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type cube struct {
	x, y, z int
}

func main() {
	p1 := part1(input)
	println("part 1:", p1)
	p2 := part2(input)
	println("part 2:", p2)
}

func part1(in string) int {
	cubes := toCubes(in)
	return calculateSurfaceArea(cubes)
}

func part2(in string) int {
	cubes := toCubes(in)
	return calculateExteriorSurfaceArea(cubes)
}

func calculateSurfaceArea(cubes []cube) int {
	cubeSet := make(map[cube]struct{})
	for _, c := range cubes {
		cubeSet[c] = struct{}{}
	}

	surfaceArea := 0
	directions := []cube{
		{1, 0, 0}, {-1, 0, 0}, // left, right
		{0, 1, 0}, {0, -1, 0}, // up, down
		{0, 0, 1}, {0, 0, -1}, // front, back
	}

	for _, c := range cubes {
		for _, d := range directions {
			neighbor := cube{c.x + d.x, c.y + d.y, c.z + d.z}
			if _, exists := cubeSet[neighbor]; !exists {
				surfaceArea++
			}
		}
	}

	return surfaceArea
}

func calculateExteriorSurfaceArea(cubes []cube) int {
	cubeSet := make(map[cube]struct{})
	for _, c := range cubes {
		cubeSet[c] = struct{}{}
	}

	// Determine the bounding box
	minX, minY, minZ := cubes[0].x, cubes[0].y, cubes[0].z
	maxX, maxY, maxZ := cubes[0].x, cubes[0].y, cubes[0].z
	for _, c := range cubes {
		if c.x < minX {
			minX = c.x
		}
		if c.y < minY {
			minY = c.y
		}
		if c.z < minZ {
			minZ = c.z
		}
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
		if c.z > maxZ {
			maxZ = c.z
		}
	}

	// Flood fill from outside the bounding box to mark reachable air cubes
	reachableAir := make(map[cube]struct{})
	queue := []cube{{minX - 1, minY - 1, minZ - 1}}
	directions := []cube{
		{1, 0, 0}, {-1, 0, 0}, // left, right
		{0, 1, 0}, {0, -1, 0}, // up, down
		{0, 0, 1}, {0, 0, -1}, // front, back
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if _, exists := reachableAir[current]; exists {
			continue
		}
		if _, exists := cubeSet[current]; exists {
			continue
		}
		if current.x < minX-1 || current.x > maxX+1 || current.y < minY-1 || current.y > maxY+1 || current.z < minZ-1 || current.z > maxZ+1 {
			continue
		}

		reachableAir[current] = struct{}{}
		for _, d := range directions {
			neighbor := cube{current.x + d.x, current.y + d.y, current.z + d.z}
			queue = append(queue, neighbor)
		}
	}

	// Calculate the exterior surface area
	surfaceArea := 0
	for _, c := range cubes {
		for _, d := range directions {
			neighbor := cube{c.x + d.x, c.y + d.y, c.z + d.z}
			if _, exists := reachableAir[neighbor]; exists {
				surfaceArea++
			}
		}
	}

	return surfaceArea
}

func toCubes(in string) []cube {
	split := strings.Split(in, "\n")
	cubes := []cube{}
	for _, line := range split {
		s := strings.Split(line, ",")
		c := cube{
			x: parse(s[0]),
			y: parse(s[1]),
			z: parse(s[2]),
		}
		cubes = append(cubes, c)
	}
	return cubes
}

func parse(in string) int {
	if i, err := strconv.Atoi(in); err == nil {
		return i
	} else {
		panic(err)
	}
}
