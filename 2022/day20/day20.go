package day20

import (
	"fmt"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2022/day20/example_input.txt"
)

func parseSequenceToList(lines []string) util.List {
	sequence := util.List{}
	for i := len(lines) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(lines[i])
		util.CheckErr(err)
		sequence.Insert(num)
	}

	return sequence
}

func parseSequenceToSlice(lines []string) []int {
	sequence := make([]int, len(lines))
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		util.CheckErr(err)
		sequence[i] = num
	}

	return sequence
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	list := parseSequenceToList(input)
	slice := parseSequenceToSlice(input)
	fmt.Println(len(slice))
	list.Display()
	list.InsertAt(2, 0)
	list.Display()
	list.Delete(4)
	list.Display()
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
