package main

import (
	"flag"

	day22 "github.com/blaine-t-bush/advent-of-code/2022/day22"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2022/day22/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2022/day22/example_input.txt"
	}

	day22.SolvePartOne(inputFile)
	day22.SolvePartTwo(inputFile)
	// day22.Viz()
}


