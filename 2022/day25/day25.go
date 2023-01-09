package day25

import (
	"fmt"
	"log"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

type bound struct {
	lower int
	upper int
}

const (
	Base = 5
)

func snafuByteToInt(snafuByte byte) int {
	var val int
	switch snafuByte {
	case '2':
		val = 2
	case '1':
		val = 1
	case '0':
		val = 0
	case '-':
		val = -1
	case '=':
		val = -2
	default:
		log.Fatal("snafuByteToInt: invalid byte")
	}

	return val
}

func snafuByteToDecimal(snafuByte byte, power int) int {
	return snafuByteToInt(snafuByte) * util.PowInt(Base, power)
}

func snafuToDecimal(snafu string) int {
	val := 0
	for i := 0; i < len(snafu); i++ {
		val += snafuByteToDecimal(snafu[len(snafu)-1-i], i)
	}

	return val
}

func getLowerBound(digits int, current string) int {
	if len(current) > digits {
		log.Fatal("getLowerBound: current cannot be longer than digits")
	}

	currentDigit := digits - 1
	lowerBound := 0
	for _, char := range current {
		lowerBound += snafuByteToInt(byte(char)) * util.PowInt(Base, currentDigit)
		currentDigit--
	}

	for i := currentDigit; i >= 0; i-- {
		lowerBound -= 2 * util.PowInt(Base, i)
	}

	return lowerBound
}

func getUpperBound(digits int, current string) int {
	if len(current) > digits {
		log.Fatal("getUpperBound: current cannot be longer than digits")
	}

	currentDigit := digits - 1
	upperBound := 0
	for _, char := range current {
		upperBound += snafuByteToInt(byte(char)) * util.PowInt(Base, currentDigit)
		currentDigit--
	}

	for i := currentDigit; i >= 0; i-- {
		upperBound += 2 * util.PowInt(Base, i)
	}

	return upperBound
}

func getNumberOfSnafuDigits(decimal int) int {
	digits := 1
	lowerBound := getLowerBound(digits, "1")
	upperBound := getUpperBound(digits, "2")

	for {
		if decimal >= lowerBound && decimal <= upperBound {
			break
		}

		digits++
		lowerBound = getLowerBound(digits, "1")
		upperBound = getUpperBound(digits, "2")
	}

	return digits
}

func updateSnafuGuess(decimal int, digits int, current string) string {
	guesses := make([]string, 5)
	guesses[0] = current + "2"
	guesses[1] = current + "1"
	guesses[2] = current + "0"
	guesses[3] = current + "-"
	guesses[4] = current + "="

	bounds := map[string]bound{}
	for _, guess := range guesses {
		bounds[guess] = bound{getLowerBound(digits, guess), getUpperBound(digits, guess)}
	}

	var bestGuess string
	for g, b := range bounds {
		if decimal <= b.upper && decimal >= b.lower {
			bestGuess = g
			break
		}
	}

	return bestGuess
}

func decimalToSnafu(decimal int) string {
	digits := getNumberOfSnafuDigits(decimal)
	guess := ""
	for {
		if len(guess) >= digits {
			break
		}

		guess = updateSnafuGuess(decimal, digits, guess)
	}

	return guess
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)

	// get sum in decimal
	sum := 0
	for _, line := range input {
		sum += snafuToDecimal(line)
	}
	fmt.Printf("sum of SNAFU numbers in decimal: %d\n", sum)

	// convert decimal sum to snafu
	converted := decimalToSnafu(sum)
	fmt.Printf("sum of SNAFU numbers in SNAFU:   %s\n", converted)
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
