package day6

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/blaine-t-bush/advent-of-code/util"
)

type race struct {
	duration int
	record int
}

func parseRaces(lines []string) []race {
	races := []race{}
	timeFields := strings.Fields(lines[0])[1:]
	distanceFields := strings.Fields(lines[1])[1:]
	for i := 0; i < len(timeFields); i++ {
		timeNum, err := strconv.Atoi(timeFields[i])
		if err != nil {
			log.Fatal("Could not parse race duration")
		}
		distanceNum, err := strconv.Atoi(distanceFields[i])
		if err != nil {
			log.Fatal("Could not parse record distance")
		}
		races = append(races, race{duration: timeNum, record: distanceNum})
	}
	return races
}

func parseRacesWithKerning(lines []string) race {
	timeFields := strings.Fields(lines[0])[1:]
	timeString := strings.Join(timeFields, "")
	distanceFields := strings.Fields(lines[1])[1:]
	distanceString := strings.Join(distanceFields, "")
	timeNum, err := strconv.Atoi(timeString)
	if err != nil {
		log.Fatal("Could not parse race duration")
	}
	distanceNum, err := strconv.Atoi(distanceString)
	if err != nil {
		log.Fatal("Could not parse record distance")
	}
	return race{duration: timeNum, record: distanceNum}
}

func (r *race) lowerBound() int {
	tF := float64(r.duration)
	dR := float64(r.record)
	return int(math.Ceil(tF/2 - math.Sqrt(tF*tF/4 - dR) + 0.0001))
}

func (r *race) upperBound() int {
	tF := float64(r.duration)
	dR := float64(r.record)
	return int(math.Floor(tF/2 + math.Sqrt(tF*tF/4 - dR) - 0.0001))
}

func (r *race) winCount() int {
	return r.upperBound() - r.lowerBound() + 1
}

func SolvePartOne(inputFile string) int {
	lines := util.ReadInput(inputFile)
	races := parseRaces(lines)
	product := 1
	for _, r := range races {
		fmt.Printf("Found bounds (%d, %d) for tF=%d, dR=%d\n", r.lowerBound(), r.upperBound(), r.duration, r.record)
		fmt.Printf("Number of winning charge durations: %d\n", r.winCount())
		product *= r.winCount()
	}
	return product
}

func SolvePartTwo(inputFile string) int {
	lines := util.ReadInput(inputFile)
	r := parseRacesWithKerning(lines)
	fmt.Printf("Found bounds (%d, %d) for tF=%d, dR=%d\n", r.lowerBound(), r.upperBound(), r.duration, r.record)
	product := r.winCount()
	fmt.Printf("Number of winning charge durations: %d\n", product)
	return product
}