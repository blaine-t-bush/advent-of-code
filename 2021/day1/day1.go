package day1

import (
	"fmt"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2021/day1/input.txt"
)

func countIncreases(nums []int) int {
	count := 0
	var previousNum int
	for i, num := range nums {
		if i != 0 && num > previousNum {
			count++
		}

		previousNum = num
	}

	return count
}

func stringsToWindows(strings []string) []int {
	ints := make([]int, len(strings)-2)
	for i := 0; i < len(strings)-2; i++ {
		num1, err := strconv.Atoi(strings[i])
		util.CheckErr(err)
		num2, err := strconv.Atoi(strings[i+1])
		util.CheckErr(err)
		num3, err := strconv.Atoi(strings[i+2])
		util.CheckErr(err)
		ints[i] = num1 + num2 + num3
	}

	return ints
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	fmt.Println(countIncreases(util.StringsToInts(input)))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	fmt.Println(countIncreases(stringsToWindows(input)))
}
