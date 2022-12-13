package day5

import (
	"fmt"
	"log"
	"regexp"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2020/day5/input.txt"
)

func parseToParts(line string) (string, string) {
	r, err := regexp.Compile(`^([FB]+)([RL]+)$`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 3 {
		log.Fatal("could not parse boarding pass to parts")
	}

	return m[1], m[2]
}

func parseRow(part string) int {
	min, max := 0, util.PowInt(2, len(part))-1
	for i := len(part) - 1; i > 0; i-- {
		if part[i] == 'F' {
			max = max - (max-min)/2 - 1
		} else if part[i] == 'B' {
			min = min + (max-min)/2 + 1
		}
	}

	if max != min {
		log.Fatal("could not parse row")
	}

	return max
}

func parseCol(part string) int {
	min, max := 0, util.PowInt(2, len(part))-1
	for i := len(part) - 1; i > 0; i-- {
		if part[i] == 'L' {
			max = max - (max-min)/2 - 1
		} else if part[i] == 'R' {
			min = min + (max-min)/2 + 1
		}
	}

	if max != min {
		log.Fatal("could not parse row")
	}

	return max
}

func parseSeat(line string) (int, int) {
	partRow, partCol := parseToParts(line)
	return parseRow(partRow), parseCol(partCol)
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
