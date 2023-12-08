package main

import (
	"flag"
	"fmt"

	day5 "github.com/blaine-t-bush/advent-of-code/2023/day5"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = "./2023/day5/input.txt"
	examplePtr := flag.Bool("ex", false, "use example input instead of real input?")
	flag.Parse()

	if *examplePtr {
		inputFile = "./2023/day5/input.example.txt"
	}

	res1 := day5.SolvePartOne(inputFile)
	fmt.Printf("Part I: lowest value %d\n", res1)

	res2 := day5.SolvePartTwo(inputFile)
	fmt.Printf("Part II: total %d\n", res2)
}


