package main

import (
	"flag"
	"fmt"

	day6 "github.com/blaine-t-bush/advent-of-code/2023/day6"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2023/day6/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2023/day6/input.example.txt"
	}

	res1 := day6.SolvePartOne(inputFile)
	fmt.Printf("Part I: product %d\n", res1)

	res2 := day6.SolvePartTwo(inputFile)
	fmt.Printf("Part II: product %d\n", res2)
}


