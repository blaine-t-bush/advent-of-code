package util

import (
	"bufio"
	"log"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadInput(filename string) []string {
	f, err := os.Open(filename)
	CheckErr(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var contents []string
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	CheckErr(scanner.Err())

	return contents
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func MinInts(n1, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
