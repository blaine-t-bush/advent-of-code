package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Assignment struct {
	ID    int
	Start int
	End   int
}

type Pair struct {
	ID          int
	Assignment1 Assignment
	Assignment2 Assignment
}

func getInput() []Pair {
	f, err := os.Open("./2022/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	pairCount := 0
	pairs := []Pair{}
	for scanner.Scan() {
		text := scanner.Text()
		pairs = append(pairs, parsePair(text, pairCount))
		pairCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pairs
}

func parsePair(text string, pairCount int) Pair {
	r, err := regexp.Compile(`^(?P<num1>\d+)-(?P<num2>\d+),(?P<num3>\d+)-(?P<num4>\d+)`)
	if err != nil {
		log.Fatal(err)
	}

	m := r.FindStringSubmatch(text)

	start1, err := strconv.Atoi(m[1])
	if err != nil {
		log.Fatal(err)
	}

	end1, err := strconv.Atoi(m[2])
	if err != nil {
		log.Fatal(err)
	}

	start2, err := strconv.Atoi(m[3])
	if err != nil {
		log.Fatal(err)
	}

	end2, err := strconv.Atoi(m[4])
	if err != nil {
		log.Fatal(err)
	}

	return Pair{
		ID: pairCount,
		Assignment1: Assignment{
			ID:    pairCount * 2,
			Start: start1,
			End:   end1,
		},
		Assignment2: Assignment{
			ID:    pairCount*2 + 1,
			Start: start2,
			End:   end2,
		},
	}
}

func SolvePartOne() int {
	totalOverlapCount := 0
	for _, pair := range getInput() {
		if (pair.Assignment1.Start >= pair.Assignment2.Start && pair.Assignment1.End <= pair.Assignment2.End) || (pair.Assignment1.Start <= pair.Assignment2.Start && pair.Assignment1.End >= pair.Assignment2.End) {
			totalOverlapCount++
		}
	}

	fmt.Println(totalOverlapCount)
	return totalOverlapCount
}

func SolvePartTwo() int {
	partialOverlapCount := 0
	for _, pair := range getInput() {
		if !(pair.Assignment1.End < pair.Assignment2.Start || pair.Assignment2.End < pair.Assignment1.Start) {
			partialOverlapCount++
		}
	}

	fmt.Println(partialOverlapCount)
	return partialOverlapCount
}
