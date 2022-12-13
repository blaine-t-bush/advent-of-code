package day4

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2020/day4/input.txt"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p passport) isValid() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func (p passport) isValidByr() bool {
	byrNum, err := strconv.Atoi(p.byr)
	return err == nil && len(p.byr) == 4 && byrNum >= 1920 && byrNum <= 2002
}

func (p passport) isValidIyr() bool {
	iyrNum, err := strconv.Atoi(p.iyr)
	return err == nil && len(p.iyr) == 4 && iyrNum >= 2010 && iyrNum <= 2020
}

func (p passport) isValidEyr() bool {
	eyrNum, err := strconv.Atoi(p.eyr)
	return err == nil && len(p.eyr) == 4 && eyrNum >= 2020 && eyrNum <= 2030
}

func (p passport) isValidHgt() bool {
	rHgt, err := regexp.Compile(`^(\d+)(cm|in)$`)
	util.CheckErr(err)

	mHgt := rHgt.FindStringSubmatch(p.hgt)
	if len(mHgt) != 3 {
		return false
	}

	hgtNum, err := strconv.Atoi(mHgt[1])
	if mHgt[2] == "cm" {
		return err == nil && len(mHgt[1]) >= 2 && len(mHgt[1]) <= 3 && hgtNum >= 150 && hgtNum <= 193
	} else {
		return err == nil && len(mHgt[1]) == 2 && hgtNum >= 59 && hgtNum <= 76
	}
}

func (p passport) isValidHcl() bool {
	rHcl, err := regexp.Compile(`^#[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]$`)
	util.CheckErr(err)
	return rHcl.MatchString(p.hcl)
}

func (p passport) isValidEcl() bool {
	rEcl, err := regexp.Compile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	util.CheckErr(err)
	return rEcl.MatchString(p.ecl)
}

func (p passport) isValidPid() bool {
	rPid, err := regexp.Compile(`^\d\d\d\d\d\d\d\d\d$`)
	util.CheckErr(err)
	return rPid.MatchString(p.pid)
}

func (p passport) isValidStrict() bool {
	return p.isValid() && p.isValidByr() && p.isValidIyr() && p.isValidEyr() && p.isValidHgt() && p.isValidHcl() && p.isValidEcl() && p.isValidPid()
}

func parsePassport(lines []string) passport {
	// concatenate all lines together with spaces
	// split concatenated line into slice by spaces
	// for each string:
	//   split by colon
	//   parse first section to get key and second section to get value

	parsed := passport{}
	concatenated := strings.Split(strings.Trim(strings.Join(lines, " "), " \r\n\t"), " ")

	for _, data := range concatenated {
		split := strings.Split(data, ":")
		switch split[0] {
		case "byr":
			parsed.byr = split[1]
		case "iyr":
			parsed.iyr = split[1]
		case "eyr":
			parsed.eyr = split[1]
		case "hgt":
			parsed.hgt = split[1]
		case "hcl":
			parsed.hcl = split[1]
		case "ecl":
			parsed.ecl = split[1]
		case "pid":
			parsed.pid = split[1]
		case "cid":
			parsed.cid = split[1]
		default:
			log.Fatal("unknown key detected")
		}
	}

	return parsed
}

func splitPassports(lines []string) [][]string {
	// split into separate slices of lines based on empty lines
	entries := [][]string{}
	temp := []string{}
	for _, line := range lines {
		if line == "\n" || line == "\r" || line == "\r\n" || line == "\n\r" || line == "" {
			entries = append(entries, temp)
			temp = []string{}
		} else {
			temp = append(temp, line)
		}
	}

	// and append last entry which was not terminated with an empty line
	entries = append(entries, temp)

	return entries
}

func SolvePartOne() {
	// get input
	input := util.ReadInput(inputFile)

	// get passports
	entries := splitPassports(input)

	// check if each passport is valid
	validCount := 0
	for _, entry := range entries {
		p := parsePassport(entry)
		if p.isValid() {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func SolvePartTwo() {
	// get input
	input := util.ReadInput(inputFile)

	// get passports
	entries := splitPassports(input)

	// check if each passport is valid
	validCount := 0
	for _, entry := range entries {
		p := parsePassport(entry)
		if p.isValidStrict() {
			validCount++
		}
	}

	fmt.Println(validCount)
}
