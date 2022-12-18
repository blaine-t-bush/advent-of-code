package day17

import (
	"fmt"
	"log"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

// rocks spawn in the following order, repeating after the end is reached,
// with their leftmost edge two spaces away from the left edge of the cave.
// after spawning, rocks are pushed 1 unit according to the input,
// then fall one unit, then repeat.
// 1
//   ####
// 2
//    #
//   ###
//    #
// 3
//     #
//     #
//   ###
// 4
//   #
//   #
//   #
//   #
// 5
//   ##
//   ##

const (
	inputFile = "./2022/day17/input.txt"
	rockType1 = iota
	rockType2
	rockType3
	rockType4
	rockType5
	moveLeft = iota
	moveRight
	left  = 0
	right = 6
	floor = 0
)

var (
	rockTypeCoords = map[int][]coord{
		rockType1: {
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},
		rockType2: {
			{1, 0},
			{0, -1},
			{1, -1},
			{2, -1},
			{1, -2},
		},
		rockType3: {
			{2, 0},
			{2, -1},
			{0, -2},
			{1, -2},
			{2, -2},
		},
		rockType4: {
			{0, 0},
			{0, -1},
			{0, -2},
			{0, -3},
		},
		rockType5: {
			{0, 0},
			{1, 0},
			{0, 1},
			{1, 1},
		},
	}

	rockTypeBounds = map[int]bound{
		rockType1: {
			top:    getTopmost(rockTypeCoords[rockType1]),
			right:  getRightmost(rockTypeCoords[rockType1]),
			bottom: getBottommost(rockTypeCoords[rockType1]),
			left:   getLeftmost(rockTypeCoords[rockType1]),
		},
		rockType2: {
			top:    getTopmost(rockTypeCoords[rockType2]),
			right:  getRightmost(rockTypeCoords[rockType2]),
			bottom: getBottommost(rockTypeCoords[rockType2]),
			left:   getLeftmost(rockTypeCoords[rockType2]),
		},
		rockType3: {
			top:    getTopmost(rockTypeCoords[rockType3]),
			right:  getRightmost(rockTypeCoords[rockType3]),
			bottom: getBottommost(rockTypeCoords[rockType3]),
			left:   getLeftmost(rockTypeCoords[rockType3]),
		},
		rockType4: {
			top:    getTopmost(rockTypeCoords[rockType4]),
			right:  getRightmost(rockTypeCoords[rockType4]),
			bottom: getBottommost(rockTypeCoords[rockType4]),
			left:   getLeftmost(rockTypeCoords[rockType4]),
		},
		rockType5: {
			top:    getTopmost(rockTypeCoords[rockType5]),
			right:  getRightmost(rockTypeCoords[rockType5]),
			bottom: getBottommost(rockTypeCoords[rockType5]),
			left:   getLeftmost(rockTypeCoords[rockType5]),
		},
	}
)

type coord struct {
	x int
	y int
}

type bound struct {
	top    int
	right  int
	bottom int
	left   int
}

type rock struct {
	rockType int
	coords   []coord
	bounds   bound
	topLeft  coord
}

func (r rock) height() int {
	return r.bounds.top - r.bounds.bottom + 1
}

func addCoords(c1 coord, c2 coord) coord {
	return coord{c1.x + c2.x, c1.y + c2.y}
}

func getLeftmost(coords []coord) int {
	xs := []int{}
	for _, c := range coords {
		xs = append(xs, c.x)
	}

	return util.MinIntsSlice(xs)
}

func getRightmost(coords []coord) int {
	xs := []int{}
	for _, c := range coords {
		xs = append(xs, c.x)
	}

	return util.MaxIntsSlice(xs)
}

func getTopmost(coords []coord) int {
	ys := []int{}
	for _, c := range coords {
		ys = append(ys, c.y)
	}

	return util.MaxIntsSlice(ys)
}

func getBottommost(coords []coord) int {
	ys := []int{}
	for _, c := range coords {
		ys = append(ys, c.y)
	}

	return util.MinIntsSlice(ys)
}

func spawnRock(rockType int, topmost int) rock {
	// leftmost point of rock is two from cave left edge.
	// topmost point of rock is at top of cave.
	r := rock{
		rockType: rockType,
		coords:   rockTypeCoords[rockType],
		bounds:   rockTypeBounds[rockType],
	}

	r.topLeft = coord{2, topmost + 3 + r.height()}
	fmt.Println(r.topLeft)

	return r
}

func moveRock(r rock, dir int, occupied map[coord]bool) (rock, bool) {
	// check x movement
	newX := r.topLeft.x
	switch dir {
	case moveLeft:
		if newX-1 >= left {
			newX--
		}
	case moveRight:
		if newX+1 <= right {
			newX++
		}
	default:
		log.Fatal("moveRock: invalid move direction")
	}

	// check y movement
	origY := r.topLeft.y
	newY := r.topLeft.y
	bottom := r.topLeft.y + r.height()
	fmt.Println(bottom)
	if _, exists := occupied[coord{newX, bottom}]; !exists {
		newY--
	}

	r.topLeft.x = newX
	r.topLeft.y = newY

	return r, newY == origY
}

func parseJets(input string) []int {
	jets := make([]int, len(input))
	for i, char := range input {
		switch char {
		case '>':
			jets[i] = moveRight
		case '<':
			jets[i] = moveLeft
		default:
			log.Fatal("parseJets: could not parse jets")
		}
	}

	return jets
}

func nextJet(jets []int, currentIndex int) (int, int) {
	if currentIndex+1 >= len(jets) {
		return jets[0], 0
	} else {
		return jets[currentIndex+1], currentIndex + 1
	}
}

func freezeRock(r rock, occupied map[coord]bool) map[coord]bool {
	coordsToAdd := []coord{}
	for _, c := range r.coords {
		coordsToAdd = append(coordsToAdd, addCoords(r.topLeft, c))
	}

	fmt.Println(coordsToAdd)

	for _, c := range coordsToAdd {
		if _, exists := occupied[c]; exists {
			draw(occupied)
			log.Fatal("freezeRock: coordinate is already occupied")
		} else {
			occupied[c] = true
		}
	}

	return occupied
}

func height(occupied map[coord]bool) int {
	h := 0
	for c := range occupied {
		if c.y > h {
			h = c.y
		}
	}

	return h
}

func draw(occupied map[coord]bool) {
	h := height(occupied)

	for y := h; y >= 1; y-- {
		fmt.Printf("|")
		for x := left; x <= right; x++ {
			if _, exists := occupied[coord{x, y}]; exists {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("|")
		fmt.Printf("\n")
	}
	fmt.Println("+-------+")
}

func SolvePartOne() {
	// get jets
	input := util.ReadInput(inputFile)
	jets := parseJets(input[0])
	topmost := floor
	currentIndex := len(jets)
	var currentDir, currentRockType int
	var stopped bool
	occupied := map[coord]bool{
		{0, 0}: true,
		{1, 0}: true,
		{2, 0}: true,
		{3, 0}: true,
		{4, 0}: true,
		{5, 0}: true,
		{6, 0}: true,
	}

	// spawn first rock (topmost begins at floor level)

	for i := 0; i < 10; i++ {
		currentRockType = (i + 1) % 5
		r := spawnRock(currentRockType, topmost)
		fmt.Printf("spawning and moving rock with index %d, type %d, height %d\n", i, currentRockType, r.height())

		for {
			// get attempted direction
			currentDir, currentIndex = nextJet(jets, currentIndex)

			// move rock until moveRock returns true
			r, stopped = moveRock(r, currentDir, occupied)

			if stopped {
				// update map of coords
				occupied = freezeRock(r, occupied)
				topmost = height(occupied)
				break
			}
		}
	}

	fmt.Println(height(occupied))
	draw(occupied)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
