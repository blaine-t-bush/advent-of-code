package util

import "strconv"

func StringsToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, str := range strings {
		num, err := strconv.Atoi(str)
		CheckErr(err)
		ints[i] = num
	}

	return ints
}
