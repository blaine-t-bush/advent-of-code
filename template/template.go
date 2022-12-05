package day0

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInput() {
	f, err := os.Open("./2022/day0/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func SolvePartOne() {
	getInput()
}

func SolvePartTwo() {
	getInput()
}
