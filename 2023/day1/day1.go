package day1

import (
	"log"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

func wordToInt(word string) int {
	switch word {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		log.Fatal("Could not parse word to integer")
	}
	return -1
}

func wordsToInts(words []string) []int {
	ints := []int{}
	for _, word := range words {
		num, err := strconv.Atoi(word)
		if err != nil {
			num = wordToInt(word)
		}
		ints = append(ints, num)
	}
	return ints
}

func intsToDoubleDigit(ints []int) int {
	if (len(ints) < 1) {
		log.Fatal("Did not receive any integers")
	}

	if (len(ints) == 1) {
		return 10*ints[0] + ints[0]
	} else {
		return 10*ints[0] + ints[len(ints)-1]
	}
}

func findFirstAndLast(line string, needles []string) []string {
	// Find first
	index := len(line) - 1
	string1 := ""
	for _, needle := range needles {
		newIndex := strings.Index(line, needle)
		if (newIndex > -1 && newIndex <= index) {
			index = newIndex
			string1 = needle
		}
	}

	// Find last
	index = 0
	string2 := ""
	for _, needle := range needles {
		newIndex := strings.LastIndex(line, needle)
		if (newIndex > -1 && newIndex >= index) {
			index = newIndex
			string2 = needle
		}
	}

	return []string{string1, string2}
}

func SolvePartOne(inputFile string) int {
	lines := util.ReadInput(inputFile)
	total := 0
	needles := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, line := range lines {
		matches := findFirstAndLast(line, needles)
		ints := wordsToInts(matches)
		total += intsToDoubleDigit(ints)
	}
	return total
}

func SolvePartTwo(inputFile string) int {
	lines := util.ReadInput(inputFile)
	total := 0
	needles := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, line := range lines {
		matches := findFirstAndLast(line, needles)
		ints := wordsToInts(matches)
		total += intsToDoubleDigit(ints)
	}
	return total
}