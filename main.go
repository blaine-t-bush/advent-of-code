package main

import (
	"flag"
	"fmt"

	day2 "github.com/blaine-t-bush/advent-of-code/2023/day2"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2023/day2/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2023/day2/example_input.txt"
	}

	res1 := day2.SolvePartOne(inputFile)
	fmt.Printf("Part I: total %d\n", res1)

	res2 := day2.SolvePartTwo(inputFile)
	fmt.Printf("Part II: total %d\n", res2)
}


