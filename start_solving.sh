#!/bin/bash
YEAR=$1
DAY=$2

mkdir -p $YEAR'/day'$DAY

URL='https://adventofcode.com/'$YEAR'/day/'$DAY
INPUT_FILE=$YEAR/day$DAY/input.txt
EXAMPLE_FILE=$YEAR/day$DAY/example_input.txt
SOLUTION_FILE=$YEAR/day$DAY/day$DAY.go

# Fetch problem input
max_fails=10
cur_fails=0
until $(curl $URL'/input' --config aoc_session --output $INPUT_FILE --silent --fail --retry 10 --retry-delay 5)
do
    ((cur_fails++))
    echo $cur_fails': Not quite yet...'
    if [ $cur_fails -ge $max_fails ]
    then
       echo 'Puzzle not yet released, please be patient.'
       exit 1
    fi
    sleep 5
done

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

	day$DAY \"github.com/blaine-t-bush/advent-of-code/2022/day$DAY\"
)

func main() {
	// parse command line flags to determine appropriate input file
	var inputFile string = \"./2022/day$DAY/input.txt\"
	examplePtr := flag.Bool(\"ex\", false, \"use example input instead of real input?\")
	flag.Parse()

	if *examplePtr {
		inputFile = \"./2022/day$DAY/example_input.txt\"
	}

	day$DAY.SolvePartOne(inputFile)
	day$DAY.SolvePartTwo(inputFile)
	// day$DAY.Viz()
}

" > "./main.go"
# Open problem page
"C:/Program Files/Google/Chrome/Application/chrome.exe" $URL
code -r $SOLUTION_FILE $INPUT_FILE $EXAMPLE_FILE