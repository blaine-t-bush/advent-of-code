package day6

import (
	"fmt"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2020/day6/input.txt"
)

// get all characters from given line
func parseVote(line string) []rune {
	return []rune(line)
}

// get unique characters from given lines
func parseVotes(lines []string) []rune {
	votes := []rune{}
	for _, line := range lines {
		votes = append(votes, parseVote(line)...)
	}

	return util.UniqueSlice(votes)
}

// get common characters from given lines
func parseCommonVotes(lines []string) []rune {
	ballots := [][]rune{}
	for _, line := range lines {
		ballots = append(ballots, []rune(line))
	}

	return util.CommonElements(ballots)
}

func parseGroups(lines []string) [][]string {
	groups := [][]string{}
	group := []string{}

	for _, line := range lines {
		if line != "" {
			group = append(group, line)
		} else {
			groups = append(groups, group)
			group = []string{}
		}
	}

	groups = append(groups, group)
	return groups
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	groups := parseGroups(input)
	votes := [][]rune{}
	for _, group := range groups {
		votes = append(votes, parseVotes(group))
	}

	sum := 0
	for _, vote := range votes {
		sum += len(vote)
	}

	fmt.Println(sum)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	groups := parseGroups(input)
	votes := [][]rune{}
	for _, group := range groups {
		votes = append(votes, parseCommonVotes(group))
	}

	sum := 0
	for _, vote := range votes {
		sum += len(vote)
	}

	fmt.Println(sum)
}
