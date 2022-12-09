#!/bin/bash
YEAR=$1
DAY=$2

mkdir -p $YEAR'/day'$DAY

URL='https://adventofcode.com/'$YEAR'/day/'$DAY
INPUT_FILE=$YEAR/day$DAY/input.txt
SOLUTION_FILE=$YEAR/day$DAY/day$DAY.go
VIZ_FILE=$YEAR/day$DAY/viz.go

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
cp -i template/viz.go $VIZ_FILE
# Update the year and day in the templates
sed -i -E "s/day0/day$DAY/" $SOLUTION_FILE
sed -i -E "s/day0/day$DAY/" $VIZ_FILE
sed -i -E "s/2022/$YEAR/" $SOLUTION_FILE
sed -i -E "s/2022/$YEAR/" $VIZ_FILE
# Update main.go
echo "package main

import (
	day$DAY \"github.com/blaine-t-bush/advent-of-code/$YEAR/day$DAY\"
)

func main() {
	day$DAY.SolvePartOne()
	day$DAY.SolvePartTwo()
    // dayDAY.Viz()
}
" > "./main.go"
# Open problem page
"C:/Program Files/Google/Chrome/Application/chrome.exe" $URL
code -r $SOLUTION_FILE $INPUT_FILE