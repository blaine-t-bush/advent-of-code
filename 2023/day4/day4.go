package day4

import (
	"log"
	"strconv"
	"strings"

	"github.com/blaine-t-bush/advent-of-code/util"
)

type card struct {
	id int
	winningNumbers []int
	revealedNumbers []int
}

func parseCard(line string) card {
	// Get card ID.
	slicesForId := strings.Split(line, ":")
	idAsString := strings.Fields(slicesForId[0])[1]
	id, err := strconv.Atoi(idAsString)
	if err != nil {
		log.Fatal("Could not parse card ID")
	}

	// Get winning numbers.
	slicesForNumbers := strings.Split(slicesForId[1], "|")
	winningNumbersAsStrings := strings.Fields(slicesForNumbers[0])
	winningNumbers := []int{}
	for _, winningNumberAsString := range winningNumbersAsStrings {
		winningNumber, err := strconv.Atoi(winningNumberAsString)
		if err != nil {
			log.Fatal("Could not parse winning number")
		}
		winningNumbers = append(winningNumbers, winningNumber)
	}

	// Get revealed numbers.
	revealedNumbersAsStrings := strings.Fields(slicesForNumbers[1])
	revealedNumbers := []int{}
	for _, revealedNumberAsString := range revealedNumbersAsStrings {
		revealedNumber, err := strconv.Atoi(revealedNumberAsString)
		if err != nil {
			log.Fatal("Could not parse revealed number")
		}
		revealedNumbers = append(revealedNumbers, revealedNumber)
	}

	return card{id: id, winningNumbers: winningNumbers, revealedNumbers: revealedNumbers}
}

func parseCards(lines []string) []card {
	cards := []card{}
	for _, line := range lines {
		cards = append(cards, parseCard(line))
	}
	return cards
}

func (c *card) countMatches() int {
	matches := 0
	for _, revealedNumber := range c.revealedNumbers {
		if util.InSlice[int](revealedNumber, c.winningNumbers) {
			matches++
		}
	}
	return matches
}

func (c *card) getScore() int {
	matches := c.countMatches()
	if matches == 0 {
		return 0
	}
	return util.IntsPow(2, c.countMatches()-1)
}

func getMaxId(cards []card) int {
	maxId := 0
	for _, card := range cards {
		if card.id > maxId {
			maxId = card.id
		}
	}
	return maxId
}

func SolvePartOne(inputFile string) int {
	lines := util.ReadInput(inputFile)
	cards := parseCards(lines)
	sum := 0
	for _, card := range cards {
		sum += card.getScore()
	}
	return sum
}

func SolvePartTwo(inputFile string) int {
	lines := util.ReadInput(inputFile)
	cards := parseCards(lines)

	// Establish original amounts
	copies := map[int]int{}
	for _, card := range cards {
		copies[card.id] = 1
	}

	// Iterate through cards
	maxId := getMaxId(cards)
	for _, card := range cards {
		for i := 0; i < copies[card.id]; i++ {
			matches := card.countMatches()
			if matches > 0 {
				for ii := card.id+1; ii <= card.id+matches && ii <= maxId; ii++ {
					copies[ii] += 1
				}
			}
		}
	}

	// Count total number of cards
	count := 0
	for _, n := range copies {
		count += n
	}

	return count
}