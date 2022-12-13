package day2

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2020/day2/input.txt"
)

type rule struct {
	min  int
	max  int
	char rune
}

type rule2 struct {
	pos1 int
	pos2 int
	char byte
}

func parseLine(line string) (rule, string) {
	r, err := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 5 {
		log.Fatal("could not parse line")
	}

	min, err := strconv.Atoi(m[1])
	util.CheckErr(err)

	max, err := strconv.Atoi(m[2])
	util.CheckErr(err)

	return rule{min, max, rune(m[3][0])}, m[4]
}

func parseLine2(line string) (rule2, string) {
	r, err := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 5 {
		log.Fatal("could not parse line")
	}

	min, err := strconv.Atoi(m[1])
	util.CheckErr(err)

	max, err := strconv.Atoi(m[2])
	util.CheckErr(err)

	return rule2{min, max, m[3][0]}, m[4]
}

func validatePassword(r rule, password string) bool {
	count := 0
	for _, char := range password {
		if char == rune(r.char) {
			count++
		}
	}

	return count <= r.max && count >= r.min
}

func validatePassword2(r rule2, password string) bool {
	chars := []byte{password[r.pos1-1], password[r.pos2-1]}
	return (chars[0] == r.char && chars[1] != r.char) || (chars[0] != r.char && chars[1] == r.char)
}

func countValidPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		if validatePassword(parseLine(line)) {
			count++
		}
	}

	return count
}

func countValidPasswords2(lines []string) int {
	count := 0
	for _, line := range lines {
		if validatePassword2(parseLine2(line)) {
			count++
		}
	}

	return count
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	fmt.Println(countValidPasswords(input))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	fmt.Println(countValidPasswords2(input))
}
