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

type seat struct {
	row int
	col int
}

func (s seat) getId() int {
	return s.row*8 + s.col
}

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
	for i := 0; i < len(part); i++ {
		if part[i] == 'F' {
			max = max - (max-min)/2 - 1
		} else if part[i] == 'B' {
			min = min + (max-min)/2 + 1
		}
	}

	if max != min {
		fmt.Printf("%s, got min %d and max %d\n", part, min, max)
		log.Fatal("could not parse row")
	}

	return max
}

func parseCol(part string) int {
	min, max := 0, util.PowInt(2, len(part))-1
	for i := 0; i < len(part); i++ {
		if part[i] == 'L' {
			max = max - (max-min)/2 - 1
		} else if part[i] == 'R' {
			min = min + (max-min)/2 + 1
		}
	}

	if max != min {
		fmt.Printf("%s, got min %d and max %d\n", part, min, max)
		log.Fatal("could not parse col")
	}

	return max
}

func parseSeat(line string) seat {
	partRow, partCol := parseToParts(line)
	return seat{
		row: parseRow(partRow),
		col: parseCol(partCol),
	}
}

func parseSeats(lines []string) []seat {
	seats := make([]seat, len(lines))
	for i, line := range lines {
		seats[i] = parseSeat(line)
	}

	return seats
}

func getMaxId(seats []seat) int {
	maxId := 0
	for _, s := range seats {
		if s.getId() > maxId {
			maxId = s.getId()
		}
	}

	return maxId
}

func getIds(seats []seat) []int {
	ids := make([]int, len(seats))
	for i, s := range seats {
		ids[i] = s.getId()
	}

	return ids
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	seats := parseSeats(input)
	fmt.Println(getMaxId(seats))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	seats := parseSeats(input)
	maxId := getMaxId(seats)
	ids := getIds(seats)
	var id int
	for i := 0; i < maxId; i++ {
		if !util.IntInSlice(i, ids) && util.IntInSlice(i-1, ids) && util.IntInSlice(i+1, ids) {
			id = i
		}
	}

	fmt.Println(id)
}
