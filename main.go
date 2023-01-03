package main

import (
	"flag"

	day23 "github.com/blaine-t-bush/advent-of-code/2022/day23"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day23/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day23/example_input.txt"
	}

	day23.SolvePartOne(inputFile)
	day23.SolvePartTwo(inputFile)
	// day23.Viz()
}
