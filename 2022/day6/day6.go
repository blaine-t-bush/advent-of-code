package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput() string {
	f, err := os.Open("./2022/day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}

func sliceContainsCharacter(chars []string, check string) bool {
	for _, char := range chars {
		if char == check {
			return true
		}
	}

	return false
}

func charactersAreUnique(chars []string) bool {
	for i, char := range chars {
		filteredChars := []string{}
		filteredChars = append(filteredChars, chars[:i]...)
		if i < len(chars)-1 {
			filteredChars = append(filteredChars, chars[i+1:]...)
		}
		if sliceContainsCharacter(filteredChars, char) {
			return false
		}
	}

	return true
}

// get index of first character after start-of-packet marker,
// which is any sequence of four unique characters.
func getFirstMarker(text string, uniqueChars int) int {
	length := len(text)
	for i := 0; i <= length-uniqueChars+1; i++ {
		if charactersAreUnique(strings.Split(text[i:i+uniqueChars], "")) {
			return i + uniqueChars - 1
		}
	}

	log.Fatal("could not find start marker")
	return -1
}

func SolvePartOne() int {
	signal := getInput()
	firstStartMarkerIndex := getFirstMarker(signal, 4) + 1
	fmt.Println(firstStartMarkerIndex)
	return firstStartMarkerIndex
}

func SolvePartTwo() int {
	signal := getInput()
	firstStartMarkerIndex := getFirstMarker(signal, 14) + 1
	fmt.Println(firstStartMarkerIndex)
	return firstStartMarkerIndex
}
