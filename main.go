package main

import (
	"flag"

	day11 "github.com/blaine-t-bush/advent-of-code/2022/day11"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day11/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day11/example_input.txt"
	}

	day11.SolvePartOne(inputFile)
	day11.SolvePartTwo(inputFile)
	// day11.Viz()
}
