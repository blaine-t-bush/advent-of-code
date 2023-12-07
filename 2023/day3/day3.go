package day3

import (
	"log"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

type schematicPart struct {
	symbol rune
	x int
	y int
}

type schematicNum struct {
	val int
	x int
	y int
}

func (n *schematicNum) x1() int {
	return n.x
}

func (n *schematicNum) x2() int {
	return n.x + n.w() - 1
}

func (n *schematicNum) w() int {
	return len(strconv.Itoa(n.val))
}

func (n *schematicNum) span() []int {
	span := []int{}
	for x := n.x1(); x <= n.x2(); x++ {
		span = append(span, x)
	}
	return span
}

func (n *schematicNum) isAdjacentTo(p schematicPart) bool {
	if n.y == p.y {
		// On same row.
		if n.x2() == p.x-1 || n.x1() == p.x+1 {
			return true
		}
	} else if n.y == p.y-1 {
		// On row above.
		for _, x := range n.span() {
			if x <= p.x+1 && x >= p.x-1 {
				return true
			}
		}
	} else if n.y == p.y+1 {
		// On row below.
		for _, x := range n.span() {
			if x <= p.x+1 && x >= p.x-1 {
				return true
			}
		}
	}

	return false
}

type schematic struct {
	w int
	h int
	parts []schematicPart
	nums []schematicNum
}

func parseSchematic(lines []string) schematic {
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	nums := []schematicNum{}
	parts := []schematicPart{}
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			char := rune(line[x])
			if char == '.' {
				// Blank space. Onwards and upwards!
				continue
			} else if util.InSlice[rune](char, digits) {
				// A number! Parse until the end.
				intAsString := string(char)
				for x2 := x+1; x2 < len(line); x2++ {
					char := rune(line[x2])
					if util.InSlice[rune](char, digits) {
						intAsString = intAsString + string(char)
					} else {
						break
					}
				}
				num, err := strconv.Atoi(intAsString)
				if err != nil {
					log.Fatal("Could not parse number")
				}
				nums = append(nums, schematicNum{val: num, x: x, y: y})
				x += len(intAsString)-1
			} else {
				// A symbol! These only have width 1.
				parts = append(parts, schematicPart{symbol: char, x: x, y: y})
			}
		}
	}

	return schematic{
		w: len(lines[0]),
		h: len(lines),
		parts: parts,
		nums: nums}
}

func SolvePartOne(inputFile string) int {
	lines := util.ReadInput(inputFile)
	schematic := parseSchematic(lines)
	sum := 0
	for _, num := range schematic.nums {
		isPartNumber := false
		for _, part := range schematic.parts {
			if num.isAdjacentTo(part) {
				isPartNumber = true
				break
			}
		}
		// fmt.Printf("Found potential part number: %d at {x: %d, y: %d}\n", num.val, num.x, num.y)

		if isPartNumber {
			sum += num.val
		}
	}
	return sum
}

func SolvePartTwo(inputFile string) int {
	lines := util.ReadInput(inputFile)
	schematic := parseSchematic(lines)
	sum := 0
	for _, part := range schematic.parts {
		if part.symbol == '*' {
			adjacentNumbers := []int{}
			for _, num := range schematic.nums {
				if num.isAdjacentTo(part) {
					adjacentNumbers = append(adjacentNumbers, num.val)
				}
			}
	
			if len(adjacentNumbers) == 2 {
				// fmt.Printf("Found gear %c at {%d, %d} with numbers %d, %d\n", part.symbol, part.x, part.y, adjacentNumbers[0], adjacentNumbers[1])
				sum += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}
	return sum
}