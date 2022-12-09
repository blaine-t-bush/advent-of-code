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

func MaxInts(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func MinIntsSlice(nums []int) int {
	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func MaxIntsSlice(nums []int) int {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
