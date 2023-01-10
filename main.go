package main

import (
	"flag"

	day20 "github.com/blaine-t-bush/advent-of-code/2022/day20"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day20/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day20/example_input.txt"
	}

	day20.SolvePartOne(inputFile)
	day20.SolvePartTwo(inputFile)
	// day20.Viz()
}
