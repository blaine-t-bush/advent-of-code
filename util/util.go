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