package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	Id       int
	calories []int
}

func (e *Elf) getTotalCalories() int {
	totalCalories := 0
	for _, calorie := range e.calories {
		totalCalories += calorie
	}

	return totalCalories
}

func getInput() []Elf {
	f, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	elfCount := 0
	currentElf := Elf{Id: elfCount, calories: []int{}}
	elves := []Elf{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elfCount++
			elves = append(elves, currentElf)
			currentElf = Elf{Id: elfCount, calories: []int{}}
		} else {
			intValue, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			currentElf.calories = append(currentElf.calories, intValue)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elves
}

func getMostCalories(elves []Elf) int {
	highestCalories := 0
	for _, elf := range elves {
		currentCalories := elf.getTotalCalories()
		if currentCalories > highestCalories {
			highestCalories = currentCalories
		}
	}

	return highestCalories
}

// get the total calories of the X elves with the most calories
func getTopXMostCalories(elves []Elf, x int) int {
	totalCalories := 0

	for i := 0; i < x; i++ {
		// add the highest calories in current list of elves
		highestCalories := 0
		highestId := -1
		for _, elf := range elves {
			currentCalories := elf.getTotalCalories()
			if currentCalories > highestCalories {
				highestId = elf.Id
				highestCalories = currentCalories
			}
		}

		totalCalories += highestCalories

		// then remove that elf from list
		elves = removeElfById(elves, highestId)
	}

	return totalCalories
}

func removeElfById(elves []Elf, idToRemove int) []Elf {
	for i, elf := range elves {
		if elf.Id == idToRemove {
			elves = append(elves[:i], elves[i+1:]...)
			break
		}
	}

	return elves
}

func SolvePartOne() {
	fmt.Println(getMostCalories(getInput()))
}

func SolvePartTwo() {
	fmt.Println(getTopXMostCalories(getInput(), 3))
}
