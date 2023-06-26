package day1

import (
	"fmt"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2020/day1/input.txt"
)

func calcProduct(nums []int) int {
	var product int
	for i, num1 := range nums {
		for ii, num2 := range nums {
			if i != ii && num1+num2 == 2020 {
				product = num1 * num2
			}
		}
	}

	return product
}

func calcProductPartTwo(nums []int) int {
	var product int
	for i, num1 := range nums {
		for ii, num2 := range nums {
			for iii, num3 := range nums {
				if i != ii && ii != iii && num1+num2+num3 == 2020 {
					product = num1 * num2 * num3
				}
			}
		}
	}

	return product
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	nums := util.StringsToInts(input)
	fmt.Println(calcProduct(nums))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	nums := util.StringsToInts(input)
	fmt.Println(calcProductPartTwo(nums))
}
