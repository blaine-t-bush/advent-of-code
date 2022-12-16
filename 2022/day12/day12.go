package day12

import (
	"fmt"

	util "github.com/blaine-t-bush/advent-of-code/util"
	"github.com/dominikbraun/graph"
)

const (
	inputFile   = "./2022/day12/input.txt"
	maxSearches = 100000
)

var (
	mapWidth  int
	mapHeight int
)

type coord struct {
	x int
	y int
}

func hash(c coord) int {
	return c.x + c.y*mapWidth
}

func unhash(i int) coord {
	return coord{i % mapWidth, i / mapWidth}
}

func createHeightmap(lines []string) (map[coord]int, coord, coord) {
	// define overall dimensions
	mapHeight = len(lines)
	mapWidth = len(lines[0])

	// create slice of locations (position and height)
	// and get start/end coordinates
	heightmap := map[coord]int{}
	var start, end coord
	for y, row := range lines {
		for x, char := range row {
			var h int
			var c coord
			if char == 'S' {
				h = 0
				c = coord{x, y}
				start = c
			} else if char == 'E' {
				h = 25
				c = coord{x, y}
				end = c
			} else {
				h = int(char - 97)
				c = coord{x, y}
			}

			heightmap[c] = h
		}
	}

	return heightmap, start, end
}

func createGraph(heightmap map[coord]int, start, end coord) graph.Graph[int, coord] {
	// convert map to graph
	g := graph.New(hash, graph.Directed())

	// add vertices
	for c := range heightmap {
		_ = g.AddVertex(c)
	}

	// add edges
	for c, h := range heightmap {
		var adjacent coord
		// add above
		adjacent = coord{c.x, c.y - 1}
		if c.y != 0 && heightmap[adjacent] <= h+1 {
			_ = g.AddEdge(hash(c), hash(adjacent))
		}

		// add below
		adjacent = coord{c.x, c.y + 1}
		if c.y != mapHeight-1 && heightmap[adjacent] <= h+1 {
			_ = g.AddEdge(hash(c), hash(adjacent))
		}

		// add right
		adjacent = coord{c.x + 1, c.y}
		if c.x != mapWidth-1 && heightmap[adjacent] <= h+1 {
			_ = g.AddEdge(hash(c), hash(adjacent))
		}

		// add left
		adjacent = coord{c.x - 1, c.y}
		if c.x != 0 && heightmap[adjacent] <= h+1 {
			_ = g.AddEdge(hash(c), hash(adjacent))
		}
	}

	return g
}

func findShortestPathLen(heightmap map[coord]int, start, end coord) int {
	g := createGraph(heightmap, start, end)
	steps := 1000000
	var shortest []int
	for i := 0; i < maxSearches; i++ {
		path, _ := graph.ShortestPath[int, coord](g, hash(start), hash(end))

		if len(path) < steps {
			steps = len(path)
			shortest = path
		}

		if i%(maxSearches/1000) == 0 {
			fmt.Printf("iteration %d of %d; shortest path length: %d\n", i, maxSearches, steps-1)
		}
	}

	fmt.Println(shortest)
	return steps - 1
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	heightmap, start, end := createHeightmap(input)
	findShortestPathLen(heightmap, start, end)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
