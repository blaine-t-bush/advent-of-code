package day8

import (
	"fmt"
	"log"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2022/day8/input.txt"
)

type Coord struct {
	X int
	Y int
}

type Tree struct {
	X int
	Y int
	H int
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func reverseStrings(strings []string) []string {
	flipped := []string{}
	for _, str := range strings {
		flipped = append(flipped, reverseString(str))
	}

	return flipped
}

func rotateStrings(strings []string) []string {
	// rotates 90 degrees clockwise (so e.g. top-left becomes top-right)
	// assumes each string has same length as slice (i.e. square)
	rotated := make([][]string, len(strings))
	for i := range rotated {
		rotated[i] = make([]string, len(strings))
	}

	for y, str := range strings {
		for x, char := range str {
			rotated[x][len(strings)-y-1] = string(char)
		}
	}

	// concatenate each row
	output := make([]string, len(strings))
	for y, str := range rotated {
		for _, char := range str {
			output[y] += char
		}
	}

	return output
}

func containsTree(trees []Tree, test Tree) bool {
	for _, tree := range trees {
		if tree.X == test.X && tree.Y == test.Y {
			return true
		}
	}

	return false
}

func removeDuplicateTrees(trees []Tree) []Tree {
	unique := []Tree{}
	for _, tree := range trees {
		if !containsTree(unique, tree) {
			unique = append(unique, tree)
		}
	}

	return unique
}

func getVisibleTreesLTR(lines []string) []Tree {
	var currentMaxH int
	visibleTrees := []Tree{}
	for y, line := range lines {
		currentMaxH = -1
		for x, char := range line {
			h := int(char - '0')
			if h > currentMaxH {
				currentMaxH = h
				visibleTrees = append(visibleTrees, Tree{x, y, h})
			} else {
				continue
			}
		}
	}

	return visibleTrees
}

func getVisibleTrees(lines []string) []Tree {
	len := len(lines)
	rotated := rotateStrings(lines)
	visibleTreesLTR := getVisibleTreesLTR(lines)
	visibleTreesRTL := getVisibleTreesLTR(reverseStrings(lines))
	visibleTreesTTB := getVisibleTreesLTR(reverseStrings(rotated))
	visibleTreesBTT := getVisibleTreesLTR(rotated)

	visibleTrees := visibleTreesLTR
	for _, tree := range visibleTreesRTL {
		transformedTree := Tree{len - tree.X - 1, tree.Y, tree.H}
		visibleTrees = append(visibleTrees, transformedTree)
	}
	for _, tree := range visibleTreesTTB {
		transformedTree := Tree{len - tree.X - 1, tree.Y, tree.H}
		transformedTree = Tree{transformedTree.Y, len - transformedTree.X - 1, tree.H}
		visibleTrees = append(visibleTrees, transformedTree)
	}
	for _, tree := range visibleTreesBTT {
		transformedTree := Tree{tree.Y, len - tree.X - 1, tree.H}
		visibleTrees = append(visibleTrees, transformedTree)
	}

	return visibleTrees
}

func getAllTrees(lines []string) []Tree {
	trees := []Tree{}
	for y, line := range lines {
		for x, char := range line {
			trees = append(trees, Tree{x, y, int(char - '0')})
		}
	}

	return trees
}

func getTreeByCoord(trees []Tree, coord Coord) Tree {
	for _, tree := range trees {
		if tree.X == coord.X && tree.Y == coord.Y {
			return tree
		}
	}

	log.Fatal("could not find tree")
	return Tree{}
}

func visualizeTrees(trees []Tree, size int) {
	viz := make([][]string, size)
	for i := range viz {
		viz[i] = make([]string, size)
	}

	for y, r := range viz {
		for x := range r {
			viz[y][x] = "O"
		}
	}

	for _, tree := range trees {
		viz[tree.Y][tree.X] = "X"
	}

	for _, r := range viz {
		for _, c := range r {
			fmt.Print(c)
		}
		fmt.Print("\n")
	}
}

func getScenicScore(coord Coord, trees []Tree, size int) int {
	tree := getTreeByCoord(trees, coord)

	viewDistanceAbove := 0
	for y := tree.Y - 1; y >= 0; y-- {
		if getTreeByCoord(trees, Coord{tree.X, y}).H < tree.H {
			viewDistanceAbove++
		} else {
			viewDistanceAbove++
			break
		}
	}

	viewDistanceBelow := 0
	for y := tree.Y + 1; y < size; y++ {
		if getTreeByCoord(trees, Coord{tree.X, y}).H < tree.H {
			viewDistanceBelow++
		} else {
			viewDistanceBelow++
			break
		}
	}

	viewDistanceLeft := 0
	for x := tree.X - 1; x >= 0; x-- {
		if getTreeByCoord(trees, Coord{x, tree.Y}).H < tree.H {
			viewDistanceLeft++
		} else {
			viewDistanceLeft++
			break
		}
	}

	viewDistanceRight := 0
	for x := tree.X + 1; x < size; x++ {
		if getTreeByCoord(trees, Coord{x, tree.Y}).H < tree.H {
			viewDistanceRight++
		} else {
			viewDistanceRight++
			break
		}
	}

	return viewDistanceAbove * viewDistanceBelow * viewDistanceLeft * viewDistanceRight
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	visibleTrees := removeDuplicateTrees(getVisibleTrees(input))
	// visualizeTrees(visibleTrees, len(input))
	fmt.Printf("visible tree count: %d\n", len(visibleTrees))
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	allTrees := getAllTrees(input)
	size := len(input)

	maxScore := 0
	for _, tree := range allTrees {
		score := getScenicScore(Coord{tree.X, tree.Y}, allTrees, size)
		if score > maxScore {
			maxScore = score
		}
	}

	fmt.Println(maxScore)
}
