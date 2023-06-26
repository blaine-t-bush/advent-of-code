package day9

import (
	"fmt"
	"log"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

func parseNumbers(lines []string) []int {
	nums := []int{}
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("Could not parse number.")
		}
		nums = append(nums, num)
	}

	return nums
}

func SolvePartOne(inputFile string) {
	lines := util.ReadInput(inputFile)
	nums := parseNumbers(lines)
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
