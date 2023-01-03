package day24

import (
	"fmt"
	"log"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	dirUp = iota
	dirRight
	dirDown
	dirLeft
)

var (
	mapHeight int
	mapWidth  int
)

type coord struct {
	x int
	y int
}

type blizzard struct {
	coord coord
	dir   int
}

func (b blizzard) advance() blizzard {
	new := b
	switch new.dir {
	case dirUp:
		if new.coord.y == 0 {
			new.coord.y = mapHeight - 1
		} else {
			new.coord.y--
		}
	case dirRight:
		if new.coord.x == mapWidth-1 {
			new.coord.x = 0
		} else {
			new.coord.x++
		}
	case dirDown:
		if new.coord.y == mapHeight-1 {
			new.coord.y = 0
		} else {
			new.coord.y++
		}
	case dirLeft:
		if new.coord.x == 0 {
			new.coord.x = mapWidth - 1
		} else {
			new.coord.x--
		}
	default:
		log.Fatal("blizzard.advance: invalid direction")
	}

	return new
}

func advanceAll(blizzards []blizzard, steps int) []blizzard {
	new := blizzards
	for s := 0; s < steps; s++ {
		old := new
		new = make([]blizzard, len(blizzards))
		for i, b := range old {
			new[i] = b.advance()
		}
	}

	return new
}

func getMapDims(lines []string) (height, width int) {
	return len(lines) - 2, len(lines[0]) - 2
}

func getBlizzards(lines []string) []blizzard {
	blizzards := []blizzard{}
	for y, line := range lines {
		for x, char := range line {
			if char != '#' && char != '.' {
				newBlizzard := blizzard{coord: coord{x: x - 1, y: y - 1}}
				switch char {
				case '^':
					newBlizzard.dir = dirUp
				case '>':
					newBlizzard.dir = dirRight
				case 'v':
					newBlizzard.dir = dirDown
				case '<':
					newBlizzard.dir = dirLeft
				default:
					log.Fatal("getBlizzards: invalid direction character")
				}
				blizzards = append(blizzards, newBlizzard)
			}
		}
	}

	return blizzards
}

func draw(blizzards []blizzard) {
	for y := -1; y <= mapHeight; y++ {
		for x := -1; x <= mapWidth; x++ {
			// draw borders
			char := "."
			if (y == -1 && x != 0) || (y == mapHeight && x != mapWidth) || x == -1 || x == mapWidth {
				char = "#"
			} else {
				// draw either blizzard or empty space
				count := 0
				for _, b := range blizzards {
					if x == b.coord.x && y == b.coord.y {
						count++
						if count > 1 {
							char = strconv.Itoa(count)
						} else {
							switch b.dir {
							case dirUp:
								char = "^"
							case dirRight:
								char = ">"
							case dirDown:
								char = "v"
							case dirLeft:
								char = "<"
							default:
								log.Fatal("draw: invalid blizzard direction")
							}
						}
					}
				}
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)
	mapHeight, mapWidth = getMapDims(input)
	blizzards := getBlizzards(input)
	draw(blizzards)
	blizzards = advanceAll(blizzards, 1)
	draw(blizzards)
	blizzards = advanceAll(blizzards, 2)
	draw(blizzards)
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	mapHeight, mapWidth = getMapDims(input)
	fmt.Println(len(input))
}
