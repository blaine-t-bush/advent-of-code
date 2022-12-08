package day2

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile  = "./2021/day2/input.txt"
	dirForward = iota
	dirDown
	dirUp
)

type coord struct {
	horizontal int
	depth      int
}

type command struct {
	dir   int
	steps int
}

func parseCommand(line string) command {
	r, err := regexp.Compile(`(forward|down|up) (\d+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 3 {
		log.Fatal("could not parse command")
	}

	// parse direction
	var dir int
	switch m[1] {
	case "forward":
		dir = dirForward
	case "down":
		dir = dirDown
	case "up":
		dir = dirUp
	default:
		log.Fatal("could not parse command direction")
	}

	// parse steps
	steps, err := strconv.Atoi(m[2])
	util.CheckErr(err)

	return command{dir, steps}
}

func parseCommands(lines []string) []command {
	commands := make([]command, len(lines))
	for i, line := range lines {
		commands[i] = parseCommand(line)
	}

	return commands
}

func calcPosition(start coord, commands []command) coord {
	for _, command := range commands {
		switch command.dir {
		case dirForward:
			start.horizontal += command.steps
		case dirDown:
			start.depth += command.steps
		case dirUp:
			start.depth -= command.steps
		}
	}

	return start
}

func calcPositionWithAim(start coord, commands []command) coord {
	aim := 0
	for _, command := range commands {
		switch command.dir {
		case dirForward:
			start.horizontal += command.steps
			start.depth += aim * command.steps
		case dirDown:
			aim += command.steps
		case dirUp:
			aim -= command.steps
		}
	}

	return start
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	commands := parseCommands(input)
	final := calcPosition(coord{0, 0}, commands)
	fmt.Println(final.horizontal * final.depth)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	commands := parseCommands(input)
	final := calcPositionWithAim(coord{0, 0}, commands)
	fmt.Println(final.horizontal * final.depth)
}
