package main

import (
	"flag"

	day25 "github.com/blaine-t-bush/advent-of-code/2022/day25"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day25/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day25/example_input.txt"
	}

	day25.SolvePartOne(inputFile)
	day25.SolvePartTwo(inputFile)
	// day25.Viz()
}


