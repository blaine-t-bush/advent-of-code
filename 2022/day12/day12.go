package day12

import (
	"fmt"
	"log"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	searchCount = 1000
	dirNorth    = iota
	dirEast
	dirSouth
	dirWest
)

type coord struct {
	x int
	y int
}

func (c coord) transform() util.Coord {
	return util.Coord{X: c.x, Y: c.y}
}

func (c coord) translate(dir int, steps int) coord {
	new := c
	switch dir {
	case dirNorth:
		new.y -= steps
	case dirEast:
		new.x += steps
	case dirSouth:
		new.y += steps
	case dirWest:
		new.x -= steps
	default:
		log.Fatal("coord.move: invalid direction")
	}

	return new
}

func (c coord) move(dir int, heightmap map[coord]int) (bool, coord) {
	new := c.translate(dir, 1)
	if newH, exists := heightmap[new]; exists && newH-heightmap[c] <= 1 {
		return true, new
	}

	return false, new
}

func parseHeightmap(lines []string) (map[coord]int, coord, coord) {
	heightmap := map[coord]int{}
	var start, end coord
	for y, line := range lines {
		for x, char := range line {
			c := coord{x, y}
			var h int
			switch char {
			case 'S':
				start = c
				h = 1
			case 'E':
				end = c
				h = 26
			default:
				h = int(char) - 96
			}
			heightmap[c] = h
		}
	}

	return heightmap, start, end
}

func createGraph(heightmap map[coord]int) util.DescendantMap {
	d := util.DescendantMap{}

	// add vertices and edges
	dirs := []int{dirNorth, dirEast, dirSouth, dirWest}
	for c := range heightmap {
		for _, dir := range dirs {
			valid, neighbor := c.move(dir, heightmap)
			if valid && len(d[c.transform()]) == 0 {
				d[c.transform()] = []util.Coord{neighbor.transform()}
			} else if valid {
				d[c.transform()] = append(d[c.transform()], neighbor.transform())
			}
		}
	}

	return d
}

func getShortestPath(heightmap map[coord]int, start util.Coord, end util.Coord) int {
	d := createGraph(heightmap)
	return int(d.ShortestPath(start, end))
}

func SolvePartOne(inputFile string) {
	fmt.Println("Part 1")
	input := util.ReadInput(inputFile)
	heightmap, start, end := parseHeightmap(input)

	fmt.Printf("Found shortest path with distance %d\n", getShortestPath(heightmap, start.transform(), end.transform()))
}

func SolvePartTwo(inputFile string) {
	fmt.Println("Part 2")
	input := util.ReadInput(inputFile)
	heightmap, _, end := parseHeightmap(input)
	possibleStarts := []coord{}
	for c, h := range heightmap {
		if h == 1 {
			possibleStarts = append(possibleStarts, c)
		}
	}

	var idealStart coord
	minDistance := 10000000
	for _, possibleStart := range possibleStarts {
		distance := getShortestPath(heightmap, possibleStart.transform(), end.transform())
		if distance != -1 {
			fmt.Printf("Path from %v has distance %d\n", possibleStart, distance)
			if distance < minDistance {
				minDistance = distance
				idealStart = possibleStart
			}
		}
	}

	fmt.Printf("Shortest path is from %v with distance %d\n", idealStart, minDistance)
}
