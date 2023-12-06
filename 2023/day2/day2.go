package day2

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

type set struct {
	r int
	g int
	b int
}

func isPossible(observed set, contents set) bool {
	return observed.r <= contents.r && observed.g <= contents.g && observed.b <= contents.b
}

func countTotals(sets []set) set {
	r, g, b := 0, 0, 0
	for _, s := range sets {
		r += s.r
		g += s.g
		b += s.b
	}
	return set{r: r, g: g, b: b}
}

func countHighest(sets []set) set {
	r, g, b := 0, 0, 0
	for _, s := range sets {
		if s.r > r {
			r = s.r
		}
		if s.g > g {
			g = s.g
		}
		if s.b > b {
			b = s.b
		}
	}
	return set{r: r, g: g, b: b}
}

func parseSetsInGame(line string) []set {
	sets := []set{}
	re := regexp.MustCompile(`((\d+) (red|green|blue))`)
	for _, substring := range strings.Split(line, ";") {
		r, g, b := 0, 0, 0
		matches := re.FindAllStringSubmatch(substring, -1)
		for _, match := range matches {
			count, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal("Could not parse count")
			}
			switch match[3] {
			case "red":
				r += count
			case "green":
				g += count
			case "blue":
				b += count
			default:
				log.Fatal("Parsed non-RGB color value")
			}
		}
		sets = append(sets, set{r: r, g: g, b: b})
	}
	return sets
}

func parseId(line string) int {
	re := regexp.MustCompile(`Game (\d+): .+`)
	matches := re.FindStringSubmatch(line)
	id, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal("Could not parse ID")
	}
	return id
}


func SolvePartOne(inputFile string) int {
	lines := util.ReadInput(inputFile)
	contents := set{r: 12, g: 13, b: 14}
	sum := 0
	for _, line := range lines {
		id := parseId(line)
		sets := parseSetsInGame(line)
		match := true
		for _, s := range sets {
			if !isPossible(s, contents) {
				match = false
			}
		}
		if match {
			sum += id
		}
	}
	return sum
}

func SolvePartTwo(inputFile string) int {
	lines := util.ReadInput(inputFile)
	sum := 0
	for _, line := range lines {
		s := countHighest(parseSetsInGame(line))
		power := s.r * s.g * s.b
		sum += power
	}
	return sum
}