#!/bin/bash
YEAR=$1
DAY=$2

mkdir -p $YEAR'/day'$DAY

URL='https://adventofcode.com/'$YEAR'/day/'$DAY
INPUT_FILE=$YEAR/day$DAY/input.txt
SOLUTION_FILE=$YEAR/day$DAY/day$DAY.go

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
cp -i template/template.go $SOLUTION_FILE
# Updates the year and day in the template
sed -i -E "s/day0/day$DAY/" $SOLUTION_FILE
sed -i -E "s/2022/$YEAR/" $SOLUTION_FILE
C:/Program/Google/Chrome/Application/chrome.exe $URL
code -r $SOLUTION_FILE $INPUT_FILE