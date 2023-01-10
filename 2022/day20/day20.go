package day20

import (
	"fmt"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	decryptionKey = 811589153
)

func parse(lines []string) (map[int]*util.Node, util.List, []int) {
	indicesToNodes := map[int]*util.Node{}
	list := util.List{}
	slice := []int{}
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		util.CheckErr(err)
		node := list.Insert(num)
		slice = append(slice, num)
		indicesToNodes[i] = node
	}

	return indicesToNodes, list, slice
}

func parseAndDecrypt(lines []string) (map[int]*util.Node, util.List, []int) {
	indicesToNodes := map[int]*util.Node{}
	list := util.List{}
	slice := []int{}
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		util.CheckErr(err)
		num *= decryptionKey
		node := list.Insert(num)
		slice = append(slice, num)
		indicesToNodes[i] = node
	}

	return indicesToNodes, list, slice
}

func SolvePartOne(inputFile string) {
	fmt.Println("solving part one")
	// parse inputs
	input := util.ReadInput(inputFile)
	indicesToNodes, list, slice := parse(input)

	// mix list according to rules
	for i, num := range slice {
		list.MoveX(indicesToNodes[i], num)
	}

	// find node where 0 is and get the 1000th, 2000th, and 3000th numbers after it,
	// wrapping cyclically
	var zeroIndex int
	for i, num := range slice {
		if num == 0 {
			zeroIndex = i
		}
	}

	current := indicesToNodes[zeroIndex]
	var coord1, coord2, coord3 int
	for i := 1; i <= 3000; i++ {
		current = list.NextCyclic(current)
		switch i {
		case 1000:
			coord1 = current.Key().(int)
		case 2000:
			coord2 = current.Key().(int)
		case 3000:
			coord3 = current.Key().(int)
		default:
			continue
		}
	}

	sum := coord1 + coord2 + coord3
	fmt.Printf("  found grove coordinates %d, %d, %d with sum %d\n", coord1, coord2, coord3, sum)
}

func SolvePartTwo(inputFile string) {
	fmt.Println("solving part one")
	// parse inputs
	input := util.ReadInput(inputFile)
	indicesToNodes, list, slice := parseAndDecrypt(input)

	// mix list according to rules
	length := len(slice)
	for j := 0; j < 10; j++ {
		fmt.Printf("  performing mix %d\n", j)
		for i, num := range slice {
			// only need to perform the non-cyclic part of moves
			// that is, subtract all full cycles from the number of steps, then
			// perform the move with the remainder
			list.MoveX(indicesToNodes[i], num%(length-1))
		}
	}

	// find node where 0 is and get the 1000th, 2000th, and 3000th numbers after it,
	// wrapping cyclically
	var zeroIndex int
	for i, num := range slice {
		if num == 0 {
			zeroIndex = i
		}
	}

	current := indicesToNodes[zeroIndex]
	var coord1, coord2, coord3 int
	for i := 1; i <= 3000; i++ {
		current = list.NextCyclic(current)
		switch i {
		case 1000:
			coord1 = current.Key().(int)
		case 2000:
			coord2 = current.Key().(int)
		case 3000:
			coord3 = current.Key().(int)
		default:
			continue
		}
	}

	sum := coord1 + coord2 + coord3
	fmt.Printf("  found grove coordinates %d, %d, %d with sum %d\n", coord1, coord2, coord3, sum)
}
