package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Crate struct {
	Label  string
	Column int
	Row    int
}

type Move struct {
	Count       int
	Source      int
	Destination int
}

func parseCrates(lines []string) []Crate {
	// find line with sequence of only numbers and spaces. these are the column labels.
	// all lines above these are the initial crate layout.
	// iterate in reverse order

	r, err := regexp.Compile(`\[\w\] |    |\[\w\]|   `)
	if err != nil {
		log.Fatal(err)
	}

	r2, err := regexp.Compile(`(?:\w)`)
	if err != nil {
		log.Fatal(err)
	}

	crates := []Crate{}
	for i := 0; i < len(lines)-1; i++ {
		m := r.FindAllString(lines[i], -1)
		for j, captured := range m {
			m2 := r2.FindAllString(captured, -1)
			if len(m2) != 0 {
				crates = append(crates, Crate{
					Label:  m2[0],
					Column: j + 1,
					Row:    len(lines) - i - 1,
				})
			}
		}
	}

	return crates
}

func parseColumns(line string) []int {
	columns := []int{}
	r, err := regexp.Compile(`(\d+)`)
	if err != nil {
		log.Fatal(err)
	}

	m := r.FindAllString(line, -1)
	for _, char := range m {
		val, err := strconv.Atoi(char)
		if err != nil {
			log.Fatal(err)
		}
		columns = append(columns, val)
	}

	return columns
}

func parseMoves(lines []string) map[int]Move {
	moves := map[int]Move{}
	for index, line := range lines {
		moves[index] = parseMove(line)
	}

	return moves
}

func parseMove(line string) Move {
	r, err := regexp.Compile(`^move (?P<count>\d+) from (?P<source>\d+) to (?P<destination>\d+)`)
	if err != nil {
		log.Fatal(err)
	}

	m := r.FindStringSubmatch(line)

	count, err := strconv.Atoi(m[1])
	if err != nil {
		log.Fatal(err)
	}

	source, err := strconv.Atoi(m[2])
	if err != nil {
		log.Fatal(err)
	}

	destination, err := strconv.Atoi(m[3])
	if err != nil {
		log.Fatal(err)
	}

	return Move{
		Count:       count,
		Source:      source,
		Destination: destination,
	}
}

func getMaxMove(moves map[int]Move) int {
	max := 0
	for num := range moves {
		if num > max {
			max = num
		}
	}

	return max
}

func getMaxRowInColumn(crates []Crate, column int) int {
	maxRow := 0
	for _, crate := range crates {
		if crate.Column == column && crate.Row > maxRow {
			maxRow = crate.Row
		}
	}

	return maxRow
}

func moveCrate(crates []Crate, move Move) []Crate {
	for i := 0; i < move.Count; i++ {
		// get crate to be moved
		var targetCrate Crate
		var targetIndex int
		for index, crate := range crates {
			// get crate to be moved
			if crate.Column == move.Source && crate.Row == getMaxRowInColumn(crates, move.Source) {
				targetCrate = crate
				targetIndex = index
			}
		}

		// remove it from list
		crates = append(crates[:targetIndex], crates[targetIndex+1:]...)

		// update it
		targetCrate.Column = move.Destination
		targetCrate.Row = getMaxRowInColumn(crates, move.Destination) + 1

		// re-add it to list
		crates = append(crates, targetCrate)
	}

	return crates
}

func moveCrate9001(crates []Crate, move Move) []Crate {
	for i := 0; i < move.Count; i++ {
		// get crate to be moved
		var targetCrate Crate
		var targetIndex int
		for index, crate := range crates {
			// get crate to be moved
			if crate.Column == move.Source && crate.Row == getMaxRowInColumn(crates, move.Source)-move.Count+i+1 {
				targetCrate = crate
				targetIndex = index
			}
		}

		// remove it from list
		crates = append(crates[:targetIndex], crates[targetIndex+1:]...)

		// update it
		targetCrate.Column = move.Destination
		targetCrate.Row = getMaxRowInColumn(crates, move.Destination) + 1

		// re-add it to list
		crates = append(crates, targetCrate)
	}

	return crates
}

func getTopCrate(crates []Crate, column int) string {
	var topCrate string
	for _, crate := range crates {
		if crate.Column == column && crate.Row == getMaxRowInColumn(crates, column) {
			topCrate = crate.Label
		}
	}

	return topCrate
}

func getInput() ([]int, []Crate, map[int]Move) {
	f, err := os.Open("./2022/day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var textCrates, textMoves []string
	cratesSection := true
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			cratesSection = false
			continue
		}

		if cratesSection {
			textCrates = append(textCrates, text)
		} else {
			textMoves = append(textMoves, text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return parseColumns(textCrates[len(textCrates)-1]), parseCrates(textCrates), parseMoves(textMoves)
}

func SolvePartOne() string {
	columns, crates, moves := getInput()

	for i := 0; i <= getMaxMove(moves); i++ {
		crates = moveCrate(crates, moves[i])
	}

	output := ""
	for _, column := range columns {
		output = output + getTopCrate(crates, column)
	}

	fmt.Println(output)
	return output
}

func SolvePartTwo() string {
	columns, crates, moves := getInput()

	for i := 0; i <= getMaxMove(moves); i++ {
		crates = moveCrate9001(crates, moves[i])
	}

	output := ""
	for _, column := range columns {
		output = output + getTopCrate(crates, column)
	}

	fmt.Println(output)
	return output
}
