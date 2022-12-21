package main

import (
	"flag"

	day19 "github.com/blaine-t-bush/advent-of-code/2022/day19"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day19/example_input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day19/example_input.txt"
	} else {
		inputFile = "./2022/day19/input.txt"
	}

	day19.SolvePartOne(inputFile)
	day19.SolvePartTwo(inputFile)
}
