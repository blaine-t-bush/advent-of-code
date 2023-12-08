package day5

import (
	"log"
	"strconv"
	"strings"

	"github.com/blaine-t-bush/advent-of-code/util"
)

type almanacEntry struct {
	sourceType string
	destType string
	maps []almanacEntryMap
}

type almanacEntryMap struct {
	sourceStart int
	destStart int
	span int
}

func parseSeeds(line string) []int {
	seedsAsStrings := strings.Fields(line[7:])
	seeds := []int{}
	for _, seedAsString := range seedsAsStrings {
		num, err := strconv.Atoi(seedAsString)
		if err != nil {
			log.Fatal("Could not parse seed number")
		}
		seeds = append(seeds, num)
	}
	return seeds
}

func parseAlmanacEntryTypes(line string) almanacEntry {
	// Parse types
	toSplit := line[0:len(line)-5]
	toSplitFields := strings.Split(toSplit, "-to-")
	sourceType, destType := toSplitFields[0], toSplitFields[1]
	return almanacEntry{sourceType: sourceType, destType: destType, maps: []almanacEntryMap{}}
}

func parseAlmanacEntryMap(line string) almanacEntryMap {
	fields := strings.Fields(line)
	sourceStartNum, err := strconv.Atoi(fields[0])
	if err != nil {
		log.Fatal("Could not parse source start")
	}
	destStartNum, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatal("Could not parse destination start")
	}
	span, err := strconv.Atoi(fields[2])
	if err != nil {
		log.Fatal("Could not parse span")
	}
	return almanacEntryMap{sourceStart: sourceStartNum, destStart: destStartNum, span: span}
}

func parseAlmanacEntries(lines []string) []almanacEntry {
	currentEntry := almanacEntry{}
	allEntries := []almanacEntry{}
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			// Blank line. Continue to next and start a new almanac entry.
			allEntries = append(allEntries, currentEntry)
		} else if line[len(line)-1] == ':' {
			// Map description line. Parse the types.
			currentEntry = parseAlmanacEntryTypes(line)
		} else {
			// Map ranges line. Parse the values.
			currentEntry.maps = append(currentEntry.maps, parseAlmanacEntryMap(line))
		}
	}
	// Input doesn't end with a newline, so we append the latest entry as well.
	allEntries = append(allEntries, currentEntry)

	return allEntries
}

func applyMap(almanacMap almanacEntryMap, num int) int {
	if num < almanacMap.destStart || num >= almanacMap.destStart + almanacMap.span {
		return num
	}
	offset := num - almanacMap.destStart
	mappedNum := almanacMap.sourceStart + offset
	return mappedNum
}

func applyMaps(almanacMaps []almanacEntryMap, num int) int {
	var mappedNum int
	for _, almanacMap := range almanacMaps {
		mappedNum = applyMap(almanacMap, num)
		if mappedNum != num {
			break
		}
	}
	return mappedNum
}

func selectEntry(almanacEntries []almanacEntry, sourceType string) almanacEntry {
	entry := almanacEntry{}
	entryFound := false
	for _, e := range almanacEntries {
		if e.sourceType == sourceType {
			entry = e
			entryFound = true
			break
		}
	}
	
	if !entryFound {
		log.Fatal("Could not find map with matching sourceType")
	}

	return entry
}

func applyMapsSequentially(almanacEntries []almanacEntry, sourceType string, destType string, num int) int {
	reachedDestType := false
	mappedNum := num
	for {
		currentEntry := selectEntry(almanacEntries, sourceType)
		newMappedNum := applyMaps(currentEntry.maps, mappedNum)
		// fmt.Printf("Mapped %d-%s to %d-%s\n", mappedNum, sourceType, newMappedNum, currentEntry.destType)
		mappedNum = newMappedNum
		sourceType = currentEntry.destType
		if sourceType == destType {
			reachedDestType = true
			break
		}
	}

	if !reachedDestType {
		log.Fatal("Could not follow maps to destType")
	}

	return mappedNum
}

func SolvePartOne(inputFile string) int {
	lines := util.ReadInput(inputFile)
	seeds := parseSeeds(lines[0])
	almanacEntries := parseAlmanacEntries(lines)
	locations := []int{}
	for _, seed := range seeds {
		locations = append(locations, applyMapsSequentially(almanacEntries, "seed", "location", seed))
	}
	return util.MinInSlice[int](locations)
}

func SolvePartTwo(inputFile string) int {
	lines := util.ReadInput(inputFile)
	seeds := parseSeeds(lines[0])
	almanacEntries := parseAlmanacEntries(lines)
	locations := []int{}
	for _, seed := range seeds {
		locations = append(locations, applyMapsSequentially(almanacEntries, "seed", "location", seed))
	}
	return util.MinInSlice[int](locations)
}