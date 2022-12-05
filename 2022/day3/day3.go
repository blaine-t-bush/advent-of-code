package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Item rune

type Rucksack struct {
	Count    int
	Contents string
}

func (i Item) getPriority() int {
	if i <= 90 { // uppercase ASCII
		return int(i) - 38
	} else if i <= 122 { // lowercase ASCII
		return int(i) - 96
	}

	log.Fatal("could not get priority of rune")
	return 0
}

func (r Rucksack) getCompartmentOne() string {
	return r.Contents[:len(r.Contents)/2]
}

func (r Rucksack) getCompartmentTwo() string {
	return r.Contents[len(r.Contents)/2:]
}

func (r Rucksack) getSharedItem() Item {
	compartment1 := r.getCompartmentOne()
	compartment2 := r.getCompartmentTwo()

	for _, char1 := range compartment1 {
		for _, char2 := range compartment2 {
			if char1 == char2 {
				return Item(char1)
			}
		}
	}

	log.Fatal("could not find shared item")
	return Item(0)
}

func getBadge(r1, r2, r3 Rucksack) Item {
	for _, char1 := range r1.Contents {
		for _, char2 := range r2.Contents {
			for _, char3 := range r3.Contents {
				if char1 == char2 && char2 == char3 {
					return Item(char1)
				}
			}
		}
	}

	log.Fatal("could not find common badge")
	return Item(0)
}

func getInput() []Rucksack {
	f, err := os.Open("./2022/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	rucksackCount := 0
	rucksacks := []Rucksack{}
	for scanner.Scan() {
		text := scanner.Text()
		rucksackCount++
		rucksacks = append(rucksacks, Rucksack{Count: rucksackCount, Contents: text})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rucksacks
}

func SolvePartOne() int {
	prioritySum := 0
	rucksacks := getInput()
	for _, rucksack := range rucksacks {
		prioritySum += rucksack.getSharedItem().getPriority()
	}

	fmt.Println(prioritySum)
	return prioritySum
}

func SolvePartTwo() int {
	prioritySum := 0
	rucksacks := getInput()
	for i := 0; i < len(rucksacks); i += 3 {
		prioritySum += getBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2]).getPriority()
	}

	fmt.Println(prioritySum)
	return prioritySum
}
