package main

import (
	"flag"
	"fmt"

	day3 "github.com/blaine-t-bush/advent-of-code/2023/day3"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2023/day3/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2023/day3/example_input.txt"
	}

	res1 := day3.SolvePartOne(inputFile)
	fmt.Printf("Part I: total %d\n", res1)

	res2 := day3.SolvePartTwo(inputFile)
	fmt.Printf("Part II: total %d\n", res2)
}


