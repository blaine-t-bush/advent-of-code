package main

import (
	"flag"

	day17 "github.com/blaine-t-bush/advent-of-code/2022/day17"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day17/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day17/example_input.txt"
	}

	day17.SolvePartOne(inputFile)
	day17.SolvePartTwo(inputFile)
	// day17.Viz()
}
