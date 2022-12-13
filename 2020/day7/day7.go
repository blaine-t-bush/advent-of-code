package day7

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2020/day7/input.txt"
)

type bag struct {
	color string
	rules map[string]int
}

func parseBagColor(line string) string {
	r, err := regexp.Compile(`^(\w+ \w+) bags contain .*$`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 2 {
		log.Fatal("could not parse bag color")
	}

	return m[1]
}

func parseBagContents(line string) map[string]int {
	// remove the color section
	r, err := regexp.Compile(`^\w+ \w+ bags contain (.*)\.$`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 2 {
		log.Fatal("could not separate bag color")
	}

	// split by comma-space and loop through results
	trimmed := strings.Trim(m[1], ".")
	split := strings.Split(trimmed, ", ")
	rules := map[string]int{}
	for _, rule := range split {
		if rule == "no other bags" {
			return rules
		}

		rContent, err := regexp.Compile(`^(\d+) (\w+ \w+) bags?$`)
		util.CheckErr(err)

		mContent := rContent.FindStringSubmatch(rule)
		if len(mContent) != 3 {
			log.Fatal("could not parse bag rule")
		}

		count, err := strconv.Atoi(mContent[1])
		util.CheckErr(err)

		rules[mContent[2]] = count
	}

	return rules
}

func parseBag(line string) bag {
	return bag{
		color: parseBagColor(line),
		rules: parseBagContents(line),
	}
}

func parseBags(lines []string) []bag {
	bags := make([]bag, len(lines))
	for i, line := range lines {
		bags[i] = parseBag(line)
	}

	return bags
}

func (b bag) canContainDirectly(color string) bool {
	for containerColor, containerCount := range b.rules {
		if containerColor == color && containerCount > 0 {
			return true
		}
	}

	return false
}

func getDirectContainers(color string, bags []bag) []bag {
	containers := []bag{}
	for _, bag := range bags {
		if bag.canContainDirectly(color) {
			containers = append(containers, bag)
		}
	}

	return containers
}

func getAllContainers(color string, bags []bag) []bag {
	// get direct containers
	containers := getDirectContainers(color, bags)

	// for each direct container, get its direct containers
	for _, bag := range containers {
		containers = append(containers, getAllContainers(bag.color, bags)...)
	}

	return containers
}

func getBagByColor(color string, bags []bag) bag {
	var found bag
	for _, bag := range bags {
		if bag.color == color {
			found = bag
		}
	}

	return found
}

func getContentsCount(color string, bags []bag) int {
	target := getBagByColor(color, bags)

	contents := 1
	for contentColor, contentCount := range target.rules {
		contents += contentCount * getContentsCount(contentColor, bags)
	}

	return contents
}

func getUniqueColors(bags []bag) []string {
	uniqueColors := []string{}
	for _, bag := range bags {
		if !util.InSlice(bag.color, uniqueColors) {
			uniqueColors = append(uniqueColors, bag.color)
		}
	}

	return uniqueColors
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	bags := parseBags(input)
	containers := getAllContainers("shiny gold", bags)
	fmt.Println(len(getUniqueColors(containers)))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	bags := parseBags(input)
	fmt.Println(getContentsCount("shiny gold", bags) - 1)
}
