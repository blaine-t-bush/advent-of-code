package main

import (
	"flag"
	"fmt"

	day4 "github.com/blaine-t-bush/advent-of-code/2023/day4"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2023/day4/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2023/day4/input.example.txt"
	}

	res1 := day4.SolvePartOne(inputFile)
	fmt.Printf("Part I: total %d\n", res1)

	res2 := day4.SolvePartTwo(inputFile)
	fmt.Printf("Part II: total %d\n", res2)
}


