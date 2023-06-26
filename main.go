package main

import (
	"flag"

	day10 "github.com/blaine-t-bush/advent-of-code/2020/day10"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2020/day10/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2020/day10/example_input.txt"
	}

	day10.SolvePartOne(inputFile)
	day10.SolvePartTwo(inputFile)
	// day10.Viz()
}


