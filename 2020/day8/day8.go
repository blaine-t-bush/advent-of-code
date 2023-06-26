package day8

import (
	"fmt"
	"log"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

type operationType int

type instruction struct {
	operation operationType
	argument  int
}

const (
	operationAcc operationType = 0
	operationJmp operationType = 1
	operationNop operationType = 2
)

func parseInstruction(line string) instruction {
	instr := instruction{}

	switch line[0:3] {
	case "acc":
		instr.operation = operationAcc
	case "jmp":
		instr.operation = operationJmp
	case "nop":
		instr.operation = operationNop
	default:
		log.Fatalf("Could not parse operation type. Line: %s\n", line)
	}

	arg, err := strconv.Atoi(line[5:])
	util.CheckErr(err)

	switch line[4] {
	case '-':
		arg *= -1
	case '+':
	default:
		log.Fatalf("Could not parse argument sign. Line: %s\n", line)
	}

	instr.argument = arg
	return instr
}

func parseInstructions(lines []string) []instruction {
	instructions := []instruction{}
	for _, line := range lines {
		instructions = append(instructions, parseInstruction(line))
	}

	return instructions
}

func runInstruction(instr instruction, currentInstrIndex int, accumulator int) (nextInstrIndex int, newAccumulator int) {
	switch instr.operation {
	case operationAcc:
		newAccumulator = accumulator + instr.argument
		nextInstrIndex = currentInstrIndex + 1
	case operationJmp:
		newAccumulator = accumulator
		nextInstrIndex = currentInstrIndex + instr.argument
	case operationNop:
		newAccumulator = accumulator
		nextInstrIndex = currentInstrIndex + 1
	default:
		log.Fatal("Invalid operation.")
	}

	return nextInstrIndex, newAccumulator
}

func runInstructionsPartOne(instructions []instruction) (accumulator int) {
	accumulator = 0
	currentInstrIndex := 0
	completedInstrIndices := []int{}

	for {
		if util.InSlice[int](currentInstrIndex, completedInstrIndices) {
			break
		}

		completedInstrIndices = append(completedInstrIndices, currentInstrIndex)
		currentInstrIndex, accumulator = runInstruction(instructions[currentInstrIndex], currentInstrIndex, accumulator)
	}

	return accumulator
}

func runInstructionsPartTwo(instructions []instruction) (accumulator int, found bool) {
	accumulator = 0
	currentInstrIndex := 0
	maxInstrIndex := len(instructions) - 1
	found = false

	for i := 0; i < 10000; i++ {
		if currentInstrIndex == maxInstrIndex+1 {
			found = true
			break
		}

		currentInstrIndex, accumulator = runInstruction(instructions[currentInstrIndex], currentInstrIndex, accumulator)
	}

	return accumulator, found
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)
	instructions := parseInstructions(input)
	accumulator := runInstructionsPartOne(instructions)
	fmt.Println(accumulator)
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	instructions := parseInstructions(input)
	accumulator, found := 0, false
	for i, instr := range instructions {
		instructionsCopy := make([]instruction, len(instructions))
		copy(instructionsCopy, instructions)

		// Replace jump with nop or vice versa.
		if instr.operation == operationJmp {
			instructionsCopy[i].operation = operationNop
		} else if instr.operation == operationNop {
			instructionsCopy[i].operation = operationJmp
		} else {
			continue
		}

		accumulator, found = runInstructionsPartTwo(instructionsCopy)
		if found {
			fmt.Printf("Found working order w/ accumulator value %d by changing instruction at index %d.\n", accumulator, i)
		}
	}
}
