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

func SumInts(vals []int) int {
	sum := 0
	for _, val := range vals {
		sum += val
	}

	return sum
}

// n to the power of m
func PowInt(n, m int) int {
	if m == 0 {
		return 1
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}

	return result
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

func IntInSlice(needle int, haystack []int) bool {
	for _, num := range haystack {
		if num == needle {
			return true
		}
	}

	return false
}

func RemoveIntFromSlice(remove int, nums []int) []int {
	new := []int{}
	for _, num := range nums {
		if num != remove {
			new = append(new, num)
		}
	}

	return new
}

func RemoveFromSlice[V comparable](remove V, elements []V) []V {
	new := []V{}
	for _, element := range elements {
		if remove != element {
			new = append(new, element)
		}
	}

	return new
}

func InSlice[V comparable](needle V, elements []V) bool {
	for _, element := range elements {
		if needle == element {
			return true
		}
	}

	return false
}

func UniqueSlice[V comparable](elements []V) []V {
	new := []V{}
	for _, element := range elements {
		if !InSlice(element, new) {
			new = append(new, element)
		}
	}

	return new
}

func CommonElements[V comparable](groups [][]V) []V {
	common := []V{}
	for _, group := range groups {
		for _, element := range group {
			inAllGroups := true
			for _, searchGroup := range groups {
				if !InSlice(element, searchGroup) {
					inAllGroups = false
					break
				}
			}

			if inAllGroups {
				common = append(common, element)
			}
		}
	}

	return UniqueSlice(common)
}
