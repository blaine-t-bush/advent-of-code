package main

import (
	"flag"

	day12 "github.com/blaine-t-bush/advent-of-code/2022/day12"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day12/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day12/example_input.txt"
	}

	day12.SolvePartOne(inputFile)
	day12.SolvePartTwo(inputFile)
	// day12.Viz()
}
