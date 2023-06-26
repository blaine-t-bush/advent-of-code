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

func LeastMultiple(nums []int) int {
	max := 1
	for _, num := range nums {
		max *= num
	}

	for m := 2; m <= max; m++ {
		isMultiple := true
		for _, num := range nums {
			if num%m != 0 {
				isMultiple = false
				break
			}
		}

		if isMultiple {
			return m
		}
	}

	return max
}
