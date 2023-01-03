package day22

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	facingNorth = 3
	facingEast  = 0
	facingSouth = 1
	facingWest  = 2
	cubeSize    = 4
)

var (
	fieldWidth int
)

type coord struct {
	x int
	y int
}

type instruction struct {
	isRotate        bool // true if instruction is rotate instead of translate
	rotateClockwise bool // true if clockwise (R), false if counterclockwise (L) or isRotate is false
	steps           int  // nil if isRotate is true, number of steps to move if isRotate is false
}

func parseField(lines []string) map[coord]bool {
	field := map[coord]bool{}
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '.':
				field[coord{x, y}] = false
			case '#':
				field[coord{x, y}] = true
			case ' ':
				continue
			default:
				log.Fatal("parseMap: invalid character encountered")
			}
		}
	}

	return field
}

func parseNextInstruction(line string) (instruction, string) {
	var i instruction
	var sliceIndex int // where to stop the slice. non-inclusive.
	var isRotate, rotateClockwise bool
	firstIndexR := strings.Index(line, "R")
	firstIndexL := strings.Index(line, "L")

	if firstIndexR == -1 && firstIndexL == -1 {
		// case: neither R nor L found. only a single translate instruction remains.
		isRotate = false
		sliceIndex = len(line)
	} else if firstIndexR == 0 {
		// case: R found in first index. next instruction is rotate.
		isRotate = true
		rotateClockwise = true
		sliceIndex = 1
	} else if firstIndexL == 0 {
		// case: L found in first index. next instruction is rotate.
		isRotate = true
		rotateClockwise = false
		sliceIndex = 1
	} else if firstIndexL == -1 || (firstIndexR != -1 && firstIndexR < firstIndexL) {
		// case: L not found or R found before L. next instruction is translate.
		isRotate = false
		sliceIndex = firstIndexR
	} else if firstIndexR == -1 || (firstIndexL != -1 && firstIndexL < firstIndexR) {
		// case: R not found or L found before R. next instruction is translate.
		isRotate = false
		sliceIndex = firstIndexL
	}

	if isRotate {
		i.isRotate = true
		i.rotateClockwise = rotateClockwise
	} else {
		steps, err := strconv.Atoi(line[:sliceIndex])
		util.CheckErr(err)
		i.isRotate = false
		i.steps = steps
	}

	return i, line[sliceIndex:]
}

func parseInstructions(line string) []instruction {
	instructions := []instruction{}
	for {
		next, remainder := parseNextInstruction(line)
		instructions = append(instructions, next)
		if len(remainder) == 0 {
			break
		} else {
			line = remainder
		}
	}

	return instructions
}

// take the current facing and proposed rotation to get the new facing.
func doRotate(rotateClockwise bool, currentFacing int) int {
	if rotateClockwise {
		switch currentFacing {
		case facingNorth:
			return facingEast
		case facingEast:
			return facingSouth
		case facingSouth:
			return facingWest
		case facingWest:
			return facingNorth
		default:
			log.Fatal("doRotate: invalid currentFacing")
			return -1
		}
	} else {
		switch currentFacing {
		case facingNorth:
			return facingWest
		case facingEast:
			return facingNorth
		case facingSouth:
			return facingEast
		case facingWest:
			return facingSouth
		default:
			log.Fatal("doRotate: invalid currentFacing")
			return -1
		}
	}
}

func getProposedCoord(currentFacing int, currentCoord coord) coord {
	proposedCoord := currentCoord
	switch currentFacing {
	case facingNorth:
		proposedCoord.y--
	case facingEast:
		proposedCoord.x++
	case facingSouth:
		proposedCoord.y++
	case facingWest:
		proposedCoord.x--
	}

	return proposedCoord
}

func getPeriodicCoord(currentFacing int, currentCoord coord, field map[coord]bool, cubeWrapping bool) coord {
	periodicCoord := currentCoord

	// TODO update wrapping rules for cube-style periodic boundary conditions
	if cubeWrapping {
		switch currentFacing {
		case facingNorth:
			// get coord that exists with same x and highest y
			for c := range field {
				if c.x == currentCoord.x && c.y > periodicCoord.y {
					periodicCoord = c
				}
			}
		case facingEast:
			// get coord that exists with same y and lowest x
			for c := range field {
				if c.y == currentCoord.y && c.x < periodicCoord.x {
					periodicCoord = c
				}
			}
		case facingSouth:
			// get coord that exists with same x and lowest y
			for c := range field {
				if c.x == currentCoord.x && c.y < periodicCoord.y {
					periodicCoord = c
				}
			}
		case facingWest:
			// get coord that exists with same y and highest x
			for c := range field {
				if c.y == currentCoord.y && c.x > periodicCoord.x {
					periodicCoord = c
				}
			}
		}
	} else {
		switch currentFacing {
		case facingNorth:
			// get coord that exists with same x and highest y
			for c := range field {
				if c.x == currentCoord.x && c.y > periodicCoord.y {
					periodicCoord = c
				}
			}
		case facingEast:
			// get coord that exists with same y and lowest x
			for c := range field {
				if c.y == currentCoord.y && c.x < periodicCoord.x {
					periodicCoord = c
				}
			}
		case facingSouth:
			// get coord that exists with same x and lowest y
			for c := range field {
				if c.x == currentCoord.x && c.y < periodicCoord.y {
					periodicCoord = c
				}
			}
		case facingWest:
			// get coord that exists with same y and highest x
			for c := range field {
				if c.y == currentCoord.y && c.x > periodicCoord.x {
					periodicCoord = c
				}
			}
		}
	}

	return periodicCoord
}

// take the current facing and coordinate, proposed steps, and field to get the new position.
func doTranslate(steps int, currentFacing int, currentCoord coord, field map[coord]bool, cubeWrapping bool) coord {
	// attempt to move one step at a time.
	for i := 0; i < steps; i++ {
		// if coordinate does not exist, determine periodic coordinate and check that instead.
		target := getProposedCoord(currentFacing, currentCoord)
		if !coordExists(field, target) {
			target = getPeriodicCoord(currentFacing, currentCoord, field, cubeWrapping)
		}

		if coordExistsAndOpen(field, target) {
			// if coordinate exists and is open, update currentCoord, reduce steps by 1, and recurse.
			currentCoord = target
		} else if coordExists(field, target) {
			// if coordinate exists but is not open, done.
			break
		} else {
			log.Fatal("doTranslate: could not find proposed coordinate")
		}
	}

	return currentCoord
}

func doInstruction(inst instruction, currentFacing int, currentCoord coord, field map[coord]bool, cubeWrapping bool) (int, coord) {
	if inst.isRotate {
		return doRotate(inst.rotateClockwise, currentFacing), currentCoord
	} else {
		return currentFacing, doTranslate(inst.steps, currentFacing, currentCoord, field, cubeWrapping)
	}
}

func doInstructions(instructions []instruction, currentFacing int, currentCoord coord, field map[coord]bool, cubeWrapping bool) (int, coord) {
	for _, inst := range instructions {
		currentFacing, currentCoord = doInstruction(inst, currentFacing, currentCoord, field, cubeWrapping)
	}

	return currentFacing, currentCoord
}

func printInstructions(instructions []instruction) {
	fmt.Println("<instructions>")
	for i, inst := range instructions {
		fmt.Printf("%4d: ", i)
		if inst.isRotate && inst.rotateClockwise {
			fmt.Print("rot  CW\n")
		} else if inst.isRotate && !inst.rotateClockwise {
			fmt.Print("rot CCW\n")
		} else {
			fmt.Printf("mov %3d\n", inst.steps)
		}
	}
	fmt.Println("</instructions>")
}

func coordExists(field map[coord]bool, target coord) bool {
	_, exists := field[target]
	return exists
}

func coordExistsAndOpen(field map[coord]bool, target coord) bool {
	return coordExists(field, target) && !field[target]
}

func getStartingCoord(field map[coord]bool) coord {
	y := 0
	for x := 0; x < fieldWidth; x++ {
		target := coord{x, y}
		if coordExistsAndOpen(field, target) {
			return target
		}
	}

	log.Fatal("getStartingCoord: could not find open coord")
	return coord{}
}

func getFieldDimensions(field map[coord]bool) (int, int) {
	width, height := 0, 0
	for c := range field {
		if c.x+1 > width {
			width = c.x + 1
		}
		if c.y+1 > height {
			height = c.y + 1
		}
	}

	return width, height
}

func calcPassword(currentFacing int, currentCoord coord) int {
	return 1000*(currentCoord.y+1) + 4*(currentCoord.x+1) + currentFacing
}

func SolvePartOne(inputFile string) {
	// parse field
	input := util.ReadInput(inputFile)
	field := parseField(input[:len(input)-1])
	fieldWidth, _ = getFieldDimensions(field)

	// parse instructions
	instructions := parseInstructions(input[len(input)-1])
	// printInstructions(instructions)

	// determine initial conditions
	currentFacing := facingEast
	currentCoord := getStartingCoord(field)

	// run through instructions and calculate answer
	finalFacing, finalCoord := doInstructions(instructions, currentFacing, currentCoord, field, false)
	fmt.Printf("col: %d, row: %d, facing: %d\n", finalCoord.x, finalCoord.y, finalFacing)
	fmt.Printf("answer: %d\n", calcPassword(finalFacing, finalCoord))
}

func SolvePartTwo(inputFile string) {
	// parse field
	input := util.ReadInput(inputFile)
	field := parseField(input[:len(input)-1])
	fieldWidth, _ = getFieldDimensions(field)

	// parse instructions
	instructions := parseInstructions(input[len(input)-1])
	// printInstructions(instructions)

	// determine initial conditions
	currentFacing := facingEast
	currentCoord := getStartingCoord(field)

	// run through instructions and calculate answer
	finalFacing, finalCoord := doInstructions(instructions, currentFacing, currentCoord, field, true)
	fmt.Printf("col: %d, row: %d, facing: %d\n", finalCoord.x, finalCoord.y, finalFacing)
	fmt.Printf("answer: %d\n", calcPassword(finalFacing, finalCoord))
}
