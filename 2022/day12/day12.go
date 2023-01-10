package day12

import (
	"fmt"
	"log"

	util "github.com/blaine-t-bush/advent-of-code/util"
	"github.com/dominikbraun/graph"
)

const (
	searchCount = 100000
	dirNorth    = iota
	dirEast
	dirSouth
	dirWest
)

var (
	mapWidth  int
	mapHeight int
)

type coord struct {
	x int
	y int
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

func getMapDimensions(heightmap map[coord]int) (int, int) {
	minX, minY := 1000000, 1000000
	maxX, maxY := 0, 0
	for c := range heightmap {
		if c.x < minX {
			minX = c.x
		}

		if c.x > maxX {
			maxX = c.x
		}

		if c.y < minY {
			minY = c.y
		}

		if c.y > maxY {
			maxY = c.y
		}
	}

	return maxX - minX, maxY - minY
}

func coordHash(c coord) string {
	return fmt.Sprintf("%d.%d", c.x, c.y)
}

func coordUnhash(i int) coord {
	y := i / mapWidth
	x := i - y
	return coord{x, y}
}

func createGraph(heightmap map[coord]int) graph.Graph[string, coord] {
	g := graph.New(coordHash, graph.Directed())

	// add vertices
	for c := range heightmap {
		_ = g.AddVertex(c)
	}

	// add edges
	dirs := []int{dirNorth, dirEast, dirSouth, dirWest}
	for c := range heightmap {
		for _, dir := range dirs {
			valid, neighbor := c.move(dir, heightmap)
			if valid {
				_ = g.AddEdge(coordHash(c), coordHash(neighbor))
			}
		}
	}

	return g
}

func getShortestPath(originalHeightmap map[coord]int, start coord, end coord, cutTop int, cutRight int, cutBottom int, cutLeft int) int {
	// truncate heightmap
	heightmap := map[coord]int{}
	for c, h := range originalHeightmap {
		if c.x >= cutLeft && c.x <= mapWidth-cutRight && c.y >= cutTop && c.y <= mapHeight-cutBottom {
			heightmap[c] = h
		}
	}

	newWidth, newHeight := getMapDimensions(heightmap)
	fmt.Printf("  truncated heightmap from %d by %d to %d by %d\n", mapWidth, mapHeight, newWidth, newHeight)

	g := createGraph(heightmap)

	shortestLength := 1000000
	var path, shortestPath []string
	fmt.Println("  beginning path searches")
	for i := 1; i <= searchCount; i++ {
		path, _ = graph.ShortestPath[string, coord](g, coordHash(start), coordHash(end))
		length := len(path) - 1
		if length > 0 && length < shortestLength {
			shortestPath = path
			shortestLength = length
		}

		if i%(searchCount/20) == 0 {
			fmt.Printf("    performed search %d of %d\n", i, searchCount)
			fmt.Printf("    current shortest path %d\n", shortestLength)
		}
	}

	fmt.Printf("  final shortest path %d\n", shortestLength)
	fmt.Printf("  path: %v\n", shortestPath)
	return shortestLength
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)
	heightmap, _, _ := parseHeightmap(input)
	mapWidth, mapHeight = getMapDimensions(heightmap)
	fmt.Println(heightmap[coord{0, 0}])

	// define custom start and end so we can test
	//   0, 20 början
	//  77, 18 bergspass
	// 142, 40 söderbacken
	// 139, 11 bergshalv
	// 137, 20 bergstopp
	// cBörjan := coord{0, 20}
	// cBergspass := coord{77, 18}
	// cSöderbacken := coord{142, 40}
	cBergshalv := coord{139, 11}
	cBergstopp := coord{137, 20}

	// run multiple times
	// fmt.Println("searching början to bergspass...")
	// path1 := getShortestPath(heightmap, cBörjan, cBergspass, 17, 84, 7, 0) // 5, 84, 5, 0

	// // fmt.Println("searching bergspass to söderbacken...")
	// path2 := getShortestPath(heightmap, cBergspass, cSöderbacken, 17, 18, 0, 74) // 15, 10, 0, 77

	// fmt.Println("searching söderbacken to bergshalv...")
	// path3 := getShortestPath(heightmap, cSöderbacken, cBergshalv, 6, 8, 0, 125) // 0, 0, 0, 125

	fmt.Println("searching bergshalv to bergstopp...")
	getShortestPath(heightmap, cBergstopp, cBergshalv, 9, 8, 6, 128) // 0, 0, 0, 125

	// fmt.Println()
	// fmt.Printf("overall shortest path: %d\n", path1+path2+path3+path4)
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
