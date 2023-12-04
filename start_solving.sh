#!/bin/bash
YEAR=$1
DAY=$2

mkdir -p "$YEAR"/day"$DAY"

URL="https://adventofcode.com/"$YEAR"/day/"$DAY"/input"
INPUT_FILE="$YEAR"/day"$DAY"/input.txt"
EXAMPLE_FILE="$YEAR"/day"$DAY"/example_input.txt"
SOLUTION_FILE="$YEAR"/day"$DAY"/day"$DAY.go"

# Fetch problem input
curl --config aoc_session --output $INPUT_FILE --retry 10 "$URL"/input"
echo "It's ready, start solving!"

# Copy the template solution file and visualization file
cp -i template/template.go $SOLUTION_FILE
# Update the year and day in the templates
sed -i -E "s/day0/day$DAY/" $SOLUTION_FILE
sed -i -E "s/2022/$YEAR/" $SOLUTION_FILE
# Create example input file
touch $EXAMPLE_FILE
# Update main.go
echo "package main

import (
	\"flag\"

	day$DAY \"github.com/blaine-t-bush/advent-of-code/$YEAR/day$DAY\"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = \"./$YEAR/day$DAY/input.txt\"
	examplePtr := flag.Bool(\"ex\", false, \"use example input instead of real input?\")
	flag.Parse()

	if *examplePtr {
		inputFile = \"./$YEAR/day$DAY/example_input.txt\"
	}

	day$DAY.SolvePartOne(inputFile)
	day$DAY.SolvePartTwo(inputFile)
	// day$DAY.Viz()
}

" > "./main.go"
