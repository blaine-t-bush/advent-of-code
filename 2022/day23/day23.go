package day23

import (
	"fmt"
	"log"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

// round
//   phase 1
//     if no other elf in one of 8 adjacent tiles, do nothing
//     if no elf N, NE, or NW, propose moving north 1
//     if no elf S, SE, or SE, propose moving south 1
//     if no elf W, SW, or NW, propose moving west 1
//     if no elf E, SE, or NE, propose moving east 1
//   phase 2
//     if elf was only one to propose moving to given position, move
//     otherwise, don't move
//   phase 3
//     rotate order of directions for consideration, e.g. start with south
//     and end with north.

const (
	dirN = iota
	dirNE
	dirE
	dirSE
	dirS
	dirSW
	dirW
	dirNW
)

var (
	dirCycles = map[int]dirInfo{
		0: {
			primary: dirN,
			all: []int{
				dirN,
				dirNE,
				dirNW,
			},
		},
		1: {
			primary: dirS,
			all: []int{
				dirS,
				dirSE,
				dirSW,
			},
		},
		2: {
			primary: dirW,
			all: []int{
				dirW,
				dirNW,
				dirSW,
			},
		},
		3: {
			primary: dirE,
			all: []int{
				dirE,
				dirNE,
				dirSE,
			},
		},
	}
)

type coord struct {
	x int
	y int
}

type proposal struct {
	start coord
	end   coord
}

type dirInfo struct {
	primary int
	all     []int
}

func (c coord) move(dir int, steps int) coord {
	moved := c
	switch dir {
	case dirN:
		moved.y -= steps
	case dirNE:
		moved.x += steps
		moved.y -= steps
	case dirE:
		moved.x += steps
	case dirSE:
		moved.x += steps
		moved.y += steps
	case dirS:
		moved.y += steps
	case dirSW:
		moved.x -= steps
		moved.y += steps
	case dirW:
		moved.x -= steps
	case dirNW:
		moved.x -= steps
		moved.y -= steps
	default:
		log.Fatal("move: invalid direction")
	}

	return moved
}

func (c coord) elfCount(steps int, dirs []int, elves map[coord]bool) int {
	count := 0
	for _, dir := range dirs {
		if elves[c.move(dir, steps)] {
			count++
		}
	}
	return count
}

func (c coord) elfCountAll(steps int, elves map[coord]bool) int {
	return c.elfCount(steps, []int{dirN, dirNE, dirE, dirSE, dirS, dirSW, dirW, dirNW}, elves)
}

func parseMap(lines []string) map[coord]bool {
	elves := map[coord]bool{}
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				elves[coord{x, y}] = true
			}
		}
	}

	return elves
}

func getProposals(elves map[coord]bool, phase int) []proposal {
	// use phase to determine appropriate directions for each check.
	// phase can be 0 thru 3. 0 means check north first, 1 means east first, etc.
	dir1 := phase
	dir2 := (phase + 1) % 4
	dir3 := (phase + 2) % 4
	dir4 := (phase + 3) % 4

	proposals := []proposal{}
	for c := range elves {
		// check if at least one elf in surrounding 8 tiles
		if c.elfCountAll(1, elves) == 0 {
			continue
		}

		var proposed coord
		if c.elfCount(1, dirCycles[dir1].all, elves) == 0 {
			proposed = c.move(dirCycles[dir1].primary, 1)
		} else if c.elfCount(1, dirCycles[dir2].all, elves) == 0 {
			proposed = c.move(dirCycles[dir2].primary, 1)
		} else if c.elfCount(1, dirCycles[dir3].all, elves) == 0 {
			proposed = c.move(dirCycles[dir3].primary, 1)
		} else if c.elfCount(1, dirCycles[dir4].all, elves) == 0 {
			proposed = c.move(dirCycles[dir4].primary, 1)
		} else {
			continue
		}

		proposals = append(proposals, proposal{start: c, end: proposed})
	}

	return proposals
}

func filterProposals(proposals []proposal) []proposal {
	countByEnd := map[coord]int{}
	for _, proposed := range proposals {
		if countByEnd[proposed.end] > 0 {
			countByEnd[proposed.end]++
		} else {
			countByEnd[proposed.end] = 1
		}
	}

	filteredEnds := []coord{}
	for end, count := range countByEnd {
		if count == 1 {
			filteredEnds = append(filteredEnds, end)
		}
	}

	filteredProposals := []proposal{}
	for _, proposed := range proposals {
		for _, end := range filteredEnds {
			if proposed.end == end {
				filteredProposals = append(filteredProposals, proposed)
			}
		}
	}

	return filteredProposals
}

func enactProposals(elves map[coord]bool, proposals []proposal) map[coord]bool {
	originCount := 0
	for _, proposed := range proposals {
		if proposed.end.x == 0 && proposed.end.y == 0 {
			originCount++
			// fmt.Printf("enactProposals: attempting to move to %v from %v\n", proposed.end, proposed.start)
		}

		if _, exists := elves[proposed.start]; !exists {
			fmt.Printf("enactProposals: key to delete does not exist (%v)\n", proposed.start)
		}

		lengthBeforeDeletion := len(elves)
		delete(elves, proposed.start)
		if len(elves) != lengthBeforeDeletion-1 {
			fmt.Println("enactProposals: delete did not delete entry")
		}

		if _, exists := elves[proposed.end]; exists {
			fmt.Printf("enactProposals: key to add already exists (%v)\n", proposed.end)
		}

		lengthBeforeAddition := len(elves)
		elves[proposed.end] = true
		if len(elves) != lengthBeforeAddition+1 {
			fmt.Println("enactProposals: map[key] did not create entry")
		}
	}

	if originCount > 1 {
		fmt.Printf("enactProposals: number of proposed moves to 0, 0 is %d\n", originCount)
	}

	return elves
}

func round(elves map[coord]bool, phase int) (map[coord]bool, bool) {
	proposals := getProposals(elves, phase)
	filteredProposals := filterProposals(proposals)
	if len(filteredProposals) == 0 {
		return elves, true
	} else {
		elves = enactProposals(elves, filteredProposals)
		return elves, false
	}
}

func rounds(elves map[coord]bool, rounds int) map[coord]bool {
	var phase int
	for i := 0; i < rounds; i++ {
		phase = i % 4
		elves, _ = round(elves, phase)
	}

	return elves
}

func getBounds(elves map[coord]bool) (int, int, int, int) {
	minX, minY := 100000, 100000
	maxX, maxY := -100000, -100000

	for c := range elves {
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

	return minX, maxX, minY, maxY
}

func countEmptyInBounds(elves map[coord]bool) int {
	elfCount := 0
	minX, maxX, minY, maxY := getBounds(elves)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if elves[coord{x, y}] {
				elfCount++
			}
		}
	}

	w, h := maxX-minX+1, maxY-minY+1
	total := w * h
	return total - elfCount
}

func draw(elves map[coord]bool) {
	minX, maxX, minY, maxY := getBounds(elves)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if elves[coord{x, y}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)
	elves := parseMap(input)
	elves = rounds(elves, 10)
	draw(elves)
	fmt.Println(countEmptyInBounds(elves))
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)

	elves := parseMap(input)
	roundCount := 0
	var phase int
	var done bool
	for {
		phase = roundCount % 4
		elves, done = round(elves, phase)

		if done {
			break
		}

		roundCount++
	}

	fmt.Println(roundCount + 1)
}
