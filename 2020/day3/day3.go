package day3

import (
	"fmt"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2020/day3/input.txt"
)

func parseTreeArray(lines []string) [][]bool {
	// create empty array
	treeArray := make([][]bool, len(lines))
	for i, line := range lines {
		treeArray[i] = make([]bool, len(line))
	}

	// populate: true for tree, false for empty
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				treeArray[y][x] = true
			} else {
				treeArray[y][x] = false
			}
		}
	}

	return treeArray
}

func checkForTree(x, y int, treeArray [][]bool) bool {
	return treeArray[y][x]
}

func slide(x, y, right, down int, treeArray [][]bool) (bool, int, int) {
	newX, newY := (x+right)%len(treeArray[0]), y+down
	return checkForTree(newX, newY, treeArray), newX, newY
}

func slides(right, down int, treeArray [][]bool) int {
	count := 0
	x, y := 0, 0
	var collide bool

	for {
		collide, x, y = slide(x, y, right, down, treeArray)

		if collide {
			count++
		}

		if y >= len(treeArray)-1 {
			break
		}
	}

	return count
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	treeArray := parseTreeArray(input)
	fmt.Println(slides(3, 1, treeArray))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	treeArray := parseTreeArray(input)
	product := 1
	product *= slides(1, 1, treeArray)
	product *= slides(3, 1, treeArray)
	product *= slides(5, 1, treeArray)
	product *= slides(7, 1, treeArray)
	product *= slides(1, 2, treeArray)
	fmt.Println(product)
}
