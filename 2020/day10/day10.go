package day10

import (
	"fmt"
	"log"
	"sort"
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

	// add device
	nums = append(nums, util.MaxInSlice[int](nums)+3)
	// sort
	sort.Ints(nums)

	return nums
}

func getDifferences(joltages []int) []int {
	differences := []int{}
	prev := 0
	for _, joltage := range joltages {
		differences = append(differences, joltage-prev)
		prev = joltage
	}

	return differences
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)
	joltages := parseNumbers(input)
	differences := getDifferences(joltages)
	fmt.Printf("1-jolt differences: %d\n", util.CountInSlice[int](1, differences))
	fmt.Printf("3-jolt differences: %d\n", util.CountInSlice[int](3, differences))
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
