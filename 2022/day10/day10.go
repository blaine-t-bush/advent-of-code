package day10

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2022/day10/input.txt"
)

type command struct {
	cycles        int
	registerShift int
}

func parseCommand(line string) command {
	cmd := command{}
	switch line[0:4] {
	case "noop":
		cmd.cycles = 1
		cmd.registerShift = 0
	case "addx":
		r, err := regexp.Compile(`addx (-?)(\d+)`)
		util.CheckErr(err)

		m := r.FindStringSubmatch(line)
		if len(m) != 3 {
			log.Fatal("could not parse register shift command")
		}

		num, err := strconv.Atoi(m[2])
		util.CheckErr(err)

		if m[1] == "-" {
			num = -num
		}

		cmd.cycles = 2
		cmd.registerShift = num
	default:
		log.Fatal("could not parse command")
	}

	return cmd
}

func parseCommands(lines []string) []command {
	cmds := []command{}
	for _, line := range lines {
		cmds = append(cmds, parseCommand(line))
	}

	return cmds
}

func runCommands(cmds []command) map[int]int {
	// create map of cycle number to register value
	states := map[int]int{}
	totalCycles := 1
	register := 1

	for _, cmd := range cmds {
		for i := 0; i < cmd.cycles; i++ {
			states[totalCycles] = register
			totalCycles++
			if i == cmd.cycles-1 {
				register = register + cmd.registerShift
			}
		}
	}

	return states
}

func draw(cmds []command) {
	h = 6
	w = 40
	// initialize screen
	screen := make([][]string, h)
	for i := range screen {
		screen[i] = make([]string, w)
	}

	for i, row := range screen {
		for ii := range row {
			screen[i][ii] = "."
		}
	}

	// get map of cycle number to register value
	states := runCommands(cmds)

	// loop through states to see when cycle number and register value
	// match according to sprite dimensions
	var left, middle, right int
	currentCol := 0
	currentRow := 0
	for cycle, register := range states {
		currentCol = (cycle - 1) % 40
		currentRow = (cycle - 1) / 40
		left, middle, right = register-1, register, register+1

		spritePositions := []int{left, middle, right}
		if util.IntInSlice(currentCol, spritePositions) {
			screen[currentRow][currentCol] = "#"
		}
	}

	for _, row := range screen {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	cmds := parseCommands(input)
	states := runCommands(cmds)

	signalsOfInterest := []int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, index := range signalsOfInterest {
		sum += states[index] * index
	}
	fmt.Println(sum)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	cmds := parseCommands(input)
	draw(cmds)
}
