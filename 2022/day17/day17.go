package day17

import (
	"fmt"
	"log"
	"reflect"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

// chamber is seven units wide
// rock spawns with two units between its left edge and the left wall,
// and three units between its bottom edge and the topmost edge of existing
// rocks, or floor if no rocks

type coord struct {
	x int
	y int
}

const (
	dirLeft               = -1
	dirRight              = 1
	chamberWidth          = 7
	spawnHorizontalOffset = 2
	spawnVerticalOffset   = 3
	rockCount             = 1000000
)

func parseJets(line string) []int {
	jets := []int{}
	for _, char := range line {
		switch char {
		case '<':
			jets = append(jets, dirLeft)
		case '>':
			jets = append(jets, dirRight)
		default:
			log.Fatal("parseJets: invalid direction")
		}
	}

	return jets
}

func initRocks() map[coord]bool {
	rocks := map[coord]bool{}
	for x := 0; x < chamberWidth; x++ {
		rocks[coord{x, 0}] = true
	}
	return rocks
}

func topmost(rocks map[coord]bool) int {
	top := 0
	for c := range rocks {
		if c.y > top {
			top = c.y
		}
	}
	return top
}

func topmostFloor(rocks map[coord]bool) int {
	top := topmost(rocks)
	for y := top; y >= 0; y-- {
		isFloor := true
		for x := 0; x < chamberWidth; x++ {
			if _, exists := rocks[coord{x, y}]; !exists {
				isFloor = false
				break
			}
		}

		if isFloor {
			return y
		}
	}

	return 0
}

func truncateToTopmostFloor(rocks map[coord]bool) map[coord]bool {
	truncated := map[coord]bool{}
	topmostFloor := topmostFloor(rocks)
	for c := range rocks {
		if c.y >= topmostFloor {
			truncated[c] = true
		}
	}

	return truncated
}

func collides(r rock, offset coord, rocks map[coord]bool) bool {
	for _, c := range r {
		if _, exists := rocks[coord{c.x + offset.x, c.y + offset.y}]; exists {
			return true
		}

		if c.x+offset.x < 0 || c.x+offset.x >= chamberWidth {
			return true
		}
	}

	return false
}

func getSpawnOffset(rocks map[coord]bool) coord {
	return coord{spawnHorizontalOffset, topmost(rocks) + spawnVerticalOffset + 1}
}

func freezeRock(r rock, offset coord, rocks map[coord]bool) map[coord]bool {
	for _, c := range r {
		rocks[coord{c.x + offset.x, c.y + offset.y}] = true
	}
	return rocks
}

func moveRock(r rock, offset coord, rocks map[coord]bool, dir int) (coord, bool) {
	new := offset
	done := false

	// check if jet movement is allowed
	withJet := coord{new.x + dir, new.y}
	if !collides(r, withJet, rocks) {
		new = withJet
	}

	// check if gravity movement is allowed
	withGravity := coord{new.x, new.y - 1}
	if !collides(r, withGravity, rocks) {
		new = withGravity
	} else {
		// if rock cannot move down, it's in its final spot
		done = true
	}

	return new, done
}

func moveRocks(jets []int, rocks map[coord]bool) map[coord]bool {
	jetIndex := 0
	for rockIndex := 0; rockIndex < rockCount; rockIndex++ {
		// select next rock type
		var r rock
		switch rockIndex % 5 {
		case 0:
			r = rock1
		case 1:
			r = rock2
		case 2:
			r = rock3
		case 3:
			r = rock4
		case 4:
			r = rock5
		default:
			log.Fatal("moveRocks: invalid rock phase")
		}

		// determine rock starting position
		offset := getSpawnOffset(rocks)

		// get truncated version of current rock structure
		truncated := truncateToTopmostFloor(rocks)

		// move rock until it stops, advancing the jet index each time
		done := false
		for {
			if jetIndex%10000 == 0 {
				fmt.Printf("current iteration: %d\n", jetIndex)
			}

			offset, done = moveRock(r, offset, truncated, jets[jetIndex%len(jets)])
			jetIndex++

			if done {
				break
			}
		}

		including, excluding := checkForPattern(rocks)

		if including {
			fmt.Printf("pattern detected, including floor! height %d, rock type %d, iteration %d\n", topmost(rocks), rockIndex%5, jetIndex)
		}

		if excluding {
			fmt.Printf("pattern detected, excluding floor! height %d, rock type %d, iteration %d\n", topmost(rocks), rockIndex%5, jetIndex)

		}

		// finally, add rock to map
		rocks = freezeRock(r, offset, rocks)
	}

	return rocks
}

func checkForPattern(rocks map[coord]bool) (bool, bool) {
	includingFloorEqual := false
	excludingFloorEqual := false
	h := topmost(rocks)

	if h%2 != 0 {
		bottomHalf := map[coord]bool{}
		topHalf := map[coord]bool{}

		for c := range rocks {
			if c.y <= (h+1)/2 {
				bottomHalf[c] = true
			} else {
				topHalf[c] = true
			}
		}

		includingFloorEqual = reflect.DeepEqual(topHalf, bottomHalf)
	} else {
		bottomHalf := map[coord]bool{}
		topHalf := map[coord]bool{}

		for c := range rocks {
			if c.y <= h/2 {
				bottomHalf[c] = true
			} else {
				topHalf[c] = true
			}
		}

		excludingFloorEqual = reflect.DeepEqual(topHalf, bottomHalf)
	}

	return includingFloorEqual, excludingFloorEqual
}

func draw(rocks map[coord]bool) {
	for y := topmost(rocks); y >= 0; y-- {
		fmt.Printf("|")
		for x := 0; x < chamberWidth; x++ {
			if _, exists := rocks[coord{x, y}]; exists {
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("|\n")
	}

	fmt.Println()
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)
	jets := parseJets(input[0])
	fmt.Printf("jet count: %d\n", len(jets))
	rocks := initRocks()
	final := moveRocks(jets, rocks)
	// draw(final)
	fmt.Printf("final height: %d\n", topmost(final))
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
