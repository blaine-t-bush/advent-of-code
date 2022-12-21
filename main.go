package main

import (
	"flag"

	day21 "github.com/blaine-t-bush/advent-of-code/2022/day21"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day21/example_input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day21/example_input.txt"
	} else {
		inputFile = "./2022/day21/input.txt"
	}

	day21.SolvePartOne(inputFile)
	day21.SolvePartTwo(inputFile)
}
