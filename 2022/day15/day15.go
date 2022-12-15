package day15

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2022/day15/input.txt"
)

type coord struct {
	x int
	y int
}

type sensor struct {
	pos       coord
	beaconPos coord
}

type xRange struct {
	x1 int
	x2 int
}

func (c1 coord) distanceTo(c2 coord) int {
	return util.AbsInt(c2.x-c1.x) + util.AbsInt(c2.y-c1.y)
}

func (s sensor) beaconDistance() int {
	return s.pos.distanceTo(s.beaconPos)
}

func (s sensor) exclusionRangeIncludesY(y int) bool {
	distance := s.beaconDistance()
	return y >= s.pos.y-distance && y <= s.pos.y+distance
}

func getBeaconCoords(sensors []sensor) []coord {
	coords := []coord{}
	for _, s := range sensors {
		coords = append(coords, s.beaconPos)
	}

	return coords
}

func getExclusionRangesAtY(sensors []sensor, y int) []coord {
	coords := []coord{}
	beaconCoords := getBeaconCoords(sensors)
	fmt.Println("fetched beacon coordinates")
	for i, s := range sensors {
		fmt.Printf("analyzing sensor %d...\n", i)
		if s.exclusionRangeIncludesY(y) {
			// get min x and max x at this y value
			// if y = s.y, min x is s.x - distance, max x is s.x + distance
			// if y = s.y +- 1, min x is s.x - distance + 1, max x is s.x + distance - 1
			distance := s.beaconDistance()
			offset := util.AbsInt(s.pos.y - y)
			minX, maxX := s.pos.x-distance+offset, s.pos.x+distance-offset
			fmt.Printf("  distance: %d\n", distance)
			fmt.Printf("  offset:   %d\n", offset)
			fmt.Printf("  x range:  %d, %d\n", minX, maxX)
			for x := minX; x <= maxX; x++ {
				if x%10000 == 0 {
					fmt.Printf("  analyzing x = %d\n", x)
				}
				possibleCoord := coord{x, y}
				if !util.InSlice[coord](possibleCoord, beaconCoords) {
					coords = append(coords, possibleCoord)
				}
			}
		}
	}

	return coords
}

func countUniqueCoords(coords []coord) int {
	uniques := map[coord]bool{}
	fmt.Println("counting unique coords...")
	fmt.Printf("  total coord count: %d\n", len(coords))
	for i, c := range coords {
		if i%10000 == 0 {
			fmt.Printf("  analyzing i = %d\n", i)
		}

		uniques[c] = true
	}

	return len(uniques)
}

func parseLine(line string) sensor {
	r, err := regexp.Compile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 5 {
		log.Fatal("could not parse line")
	}

	x1, err := strconv.Atoi(m[1])
	util.CheckErr(err)

	y1, err := strconv.Atoi(m[2])
	util.CheckErr(err)

	x2, err := strconv.Atoi(m[3])
	util.CheckErr(err)

	y2, err := strconv.Atoi(m[4])
	util.CheckErr(err)

	return sensor{
		pos: coord{
			x: x1,
			y: y1,
		},
		beaconPos: coord{
			x: x2,
			y: y2,
		},
	}
}

func parseLines(lines []string) []sensor {
	sensors := make([]sensor, len(lines))
	for i, line := range lines {
		sensors[i] = parseLine(line)
	}

	return sensors
}

func (c coord) calcTuningFrequency() int {
	return c.x*4000000 + c.y
}

func getExcludedMap(sensors []sensor) map[coord]bool {
	excluded := map[coord]bool{}
	fmt.Println("creating excluded map...")
	for i, s := range sensors {
		fmt.Printf("  analyzing sensor %d...\n", i)
		d := s.beaconDistance()
		for y := s.pos.y - d; y <= s.pos.y+d; y++ {
			if y%1000 == 0 {
				fmt.Printf("  y = %d of %d\n", y, s.pos.y+d)
			}
			offset := util.AbsInt(y - s.pos.y)
			for x := s.pos.x - d + offset; x <= s.pos.x+d-offset; x++ {
				excluded[coord{x, y}] = true
			}
		}
	}

	return excluded
}

func getExcludedBoundsX(sensors []sensor) map[int][]xRange {
	fmt.Println("getting excluded bounds...")
	excludedBounds := map[int][]xRange{}
	for i, s := range sensors {
		fmt.Printf("   analyzing sensor %d...\n", i)
		d := s.beaconDistance()
		// loop through y values
		for y := s.pos.y - d; y <= s.pos.y+d; y++ {
			// if map value doesn't exist, create new slice of xRanges
			// if map value does exist, append to slice of xRanges
			offset := util.AbsInt(y - s.pos.y)
			span := xRange{s.pos.x - d + offset, s.pos.x + d - offset}
			if _, exists := excludedBounds[y]; exists {
				copy := excludedBounds[y]
				excludedBounds[y] = append(copy, span)
			} else {
				excludedBounds[y] = []xRange{span}
			}
		}
	}

	return excludedBounds
}

func findDistressBeacon(excluded map[coord]bool, minX, maxX, minY, maxY int) coord {
	found := coord{}
	doBreak := false
	fmt.Println("finding distress beacon...")
	for y := minY; y <= maxY; y++ {
		if y%10000 == 0 {
			fmt.Printf("  y = %d\n", y)
		}
		for x := minX; x <= maxX; x++ {
			if _, exists := excluded[coord{x, y}]; !exists {
				found = coord{x, y}
				doBreak = true
				break
			}
		}

		if doBreak {
			break
		}
	}

	return found
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	sensors := parseLines(input)
	exclusionRanges := getExclusionRangesAtY(sensors, 2000000)
	count := countUniqueCoords(exclusionRanges)
	fmt.Println(count)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	sensors := parseLines(input)
	excludedBoundsX := getExcludedBoundsX(sensors)
	fmt.Println(len(excludedBoundsX))

	MAX := 4000000
	possibleXs := map[int][]int{}
	for y := 0; y <= MAX; y++ {
		bounds := excludedBoundsX[y]
		for _, bound := range bounds {
			if bound.x2+1 <= MAX {
				possibleXs[y] = append(possibleXs[y], bound.x2+1) // add the x-coord past the max
			}

			if bound.x1-1 >= 0 {
				possibleXs[y] = append(possibleXs[y], bound.x1-1) // add the x-coord before the min
			}
		}
	}

	beaconCoord := coord{}
	for y, possXs := range possibleXs {
		if y%10000 == 0 {
			fmt.Printf("analyzing possible Xs for y = %d\n", y)
		}

		for _, possX := range possXs {
			found := true
			for _, bounds := range excludedBoundsX[y] {
				if possX >= bounds.x1 && possX <= bounds.x2 {
					found = false
					break
				}
			}

			if found {
				beaconCoord = coord{possX, y}
			}
		}
	}

	fmt.Println(beaconCoord.calcTuningFrequency())
	// min, max := 0, 20
	// excluded := getExcludedMap(sensors)
	// beacon := findDistressBeacon(excluded, min, max, min, max)
	// fmt.Println(beacon)
	// fmt.Println(beacon.calcTuningFrequency())
}
