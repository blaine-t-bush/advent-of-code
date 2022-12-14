package day14

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile   = "./2022/day14/input.txt"
	terrainRock = iota
	terrainSand
	extraSpaceLeft  = 1000
	extraSpaceRight = 1000
)

type coord struct {
	x int
	y int
}

func (c coord) moveDown() coord {
	return coord{c.x, c.y + 1}
}

func (c coord) moveDiagonalLeft() coord {
	return coord{c.x - 1, c.y + 1}
}

func (c coord) moveDiagonalRight() coord {
	return coord{c.x + 1, c.y + 1}
}

func parsePaths(lines []string) [][]coord {
	// get slices of coords
	pathsStrings := make([][]string, len(lines))
	for i, line := range lines {
		pathsStrings[i] = strings.Split(line, " -> ")
	}

	paths := make([][]coord, len(lines))
	for i, pathStrings := range pathsStrings {
		path := make([]coord, len(pathStrings))
		for ii, pathString := range pathStrings {
			split := strings.Split(pathString, ",")
			if len(split) != 2 {
				log.Fatal("could not parse rock coord")
			}

			x, err := strconv.Atoi(split[0])
			util.CheckErr(err)

			y, err := strconv.Atoi(split[1])
			util.CheckErr(err)

			path[ii] = coord{x, y}
		}
		paths[i] = path
	}

	return paths
}

func fillPaths(paths [][]coord) [][]coord {
	filledPaths := make([][]coord, len(paths))
	for i, path := range paths {
		filledPath := []coord{}
		for ii, point := range path {
			if ii == len(path)-1 {
				// last point has no continuation
				continue
			} else {
				nextPoint := path[ii+1]
				if point.x == nextPoint.x && point.y == nextPoint.y {
					// next point is same point. do nothing.
					continue
				} else if point.x == nextPoint.x {
					// filledPath is vertical
					if point.y < nextPoint.y {
						// next point is below (y increases downward)
						for y := point.y; y <= nextPoint.y; y++ {
							filledPath = append(filledPath, coord{x: point.x, y: y})
						}
					} else {
						// next point is above (y increases downward)
						for y := point.y; y >= nextPoint.y; y-- {
							filledPath = append(filledPath, coord{x: point.x, y: y})
						}
					}
				} else if point.y == nextPoint.y {
					// filledPath is horizontal
					if point.x < nextPoint.x {
						// next point is right
						for x := point.x; x <= nextPoint.x; x++ {
							filledPath = append(filledPath, coord{x: x, y: point.y})
						}
					} else {
						// next point is below
						for x := point.x; x >= nextPoint.x; x-- {
							filledPath = append(filledPath, coord{x: x, y: point.y})
						}
					}
				} else {
					log.Fatal("projected path is not horizontal or vertical")
				}
			}
		}

		filledPaths[i] = filledPath
	}

	return filledPaths
}

func addFloor(rocksAndSand map[coord]int) map[coord]int {
	minX, maxX, _, maxY := getBounds(rocksAndSand)
	for x := minX - extraSpaceLeft; x <= maxX+extraSpaceRight; x++ {
		rocksAndSand[coord{x, maxY + 2}] = terrainRock
	}

	return rocksAndSand
}

func parseRocks(lines []string) map[coord]int {
	paths := parsePaths(lines)
	filledPaths := fillPaths(paths)

	rocks := map[coord]int{}
	for _, paths := range filledPaths {
		for _, coord := range paths {
			rocks[coord] = terrainRock
		}
	}

	return rocks
}

func moveSandOne(start coord, rocksAndSand map[coord]int) coord {
	// priority list:
	//   1. move down 1
	//   2. move down 1 and left 1
	//   3. move down 1 and right 1
	//   4. rest. never move again. create new sand at source.

	down := start.moveDown()
	diagLeft := start.moveDiagonalLeft()
	diagRight := start.moveDiagonalRight()

	if _, occupied := rocksAndSand[down]; !occupied {
		return down
	} else if _, occupied := rocksAndSand[diagLeft]; !occupied {
		return diagLeft
	} else if _, occupied := rocksAndSand[diagRight]; !occupied {
		return diagRight
	} else {
		return start
	}
}

func moveSand(start coord, rocksAndSand map[coord]int) (coord, bool) {
	// returns coord of sand's ending location and a bool.
	// bool is false if sand came to rest, true if sand kept moving after many rounds.
	var moved coord
	count := 0
	previous := start
	for {
		count++
		moved = moveSandOne(previous, rocksAndSand)

		if moved == start {
			return moved, true
		}

		if moved == previous {
			return moved, false
		}

		if count >= 10000 {
			return moved, true
		}

		previous = moved
	}
}

func countSand(rocksAndSand map[coord]int) int {
	count := 0
	for _, terrain := range rocksAndSand {
		if terrain == terrainSand {
			count++
		}
	}

	return count
}

func getBounds(rocksAndSand map[coord]int) (int, int, int, int) {
	minX, minY := 1000000, 1000000
	maxX, maxY := 0, 0
	for point := range rocksAndSand {
		if point.x > maxX {
			maxX = point.x
		} else if point.x < minX {
			minX = point.x
		}

		if point.y > maxY {
			maxY = point.y
		} else if point.y < minY {
			minY = point.y
		}
	}

	return minX, maxX, minY, maxY
}

func draw(rocksAndSand map[coord]int) {
	minX, maxX, minY, maxY := getBounds(rocksAndSand)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			terrain, exists := rocksAndSand[coord{x, y}]

			var char string
			if !exists {
				char = "."
			} else {
				switch terrain {
				case terrainRock:
					char = "#"
				case terrainSand:
					char = "o"
				}
			}

			fmt.Print(char)
		}
		fmt.Print("\n")
	}
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	rocksAndSand := parseRocks(input)
	defaultStart := coord{500, 0}
	for {
		restingCoord, trigger := moveSand(defaultStart, rocksAndSand)

		if trigger {
			break
		}

		rocksAndSand[restingCoord] = terrainSand
	}

	// draw(rocksAndSand)
	fmt.Println(countSand(rocksAndSand))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	rocksAndSand := parseRocks(input)
	rocksAndSand = addFloor(rocksAndSand)
	defaultStart := coord{500, 0}
	for {
		restingCoord, trigger := moveSand(defaultStart, rocksAndSand)

		if trigger {
			break
		}

		rocksAndSand[restingCoord] = terrainSand
	}

	rocksAndSand[defaultStart] = terrainSand

	// draw(rocksAndSand)
	fmt.Println(countSand(rocksAndSand))
}
