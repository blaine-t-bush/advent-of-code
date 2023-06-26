package util

import (
	"golang.org/x/exp/constraints"
)

func CountInSlice[V comparable](needle V, elements []V) int {
	count := 0
	for _, element := range elements {
		if element == needle {
			count++
		}
	}

	return count
}

func MaxInSlice[V constraints.Ordered](elements []V) V {
	max := elements[0]
	for _, element := range elements {
		if element > max {
			max = element
		}
	}

	return max
}

func MinInSlice[V constraints.Ordered](elements []V) V {
	min := elements[0]
	for _, element := range elements {
		if element < min {
			min = element
		}
	}

	return min
}

func AllEqual[V comparable](elements []V) bool {
	for _, element := range elements {
		if element != elements[0] {
			return false
		}
	}

	return true
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
