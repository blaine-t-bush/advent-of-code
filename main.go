package main

import (
	"flag"

	day24 "github.com/blaine-t-bush/advent-of-code/2022/day24"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day24/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day24/example_input.txt"
	}

	day24.SolvePartOne(inputFile)
	day24.SolvePartTwo(inputFile)
	// day24.Viz()
}


