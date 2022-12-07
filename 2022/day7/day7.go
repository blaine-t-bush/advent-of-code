package day7

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

type File struct {
	DirId int
	Name  string
	Size  int
}

const (
	inputFile = "./2022/day7/input.txt"
)

func parseCd(line string) string {
	r, err := regexp.Compile(`\$ cd (.+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if m == nil || len(m) != 2 {
		log.Fatal("could not parse target dir")
	}

	return m[1]
}

func parseFile(dirId int, line string) File {
	r, err := regexp.Compile(`(\d+) (.+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if m == nil || len(m) != 3 {
		log.Fatal("could not parse file")
	}

	size, err := strconv.Atoi(m[1])
	util.CheckErr(err)

	return File{
		DirId: dirId,
		Name:  m[2],
		Size:  size,
	}
}

func parseCommands(lines []string) (map[int]int, []File) {
	dirTree := map[int]int{0: 0}
	dirCount := 0
	presentDirId := 0
	parentDirId := 0
	files := []File{}

	for _, line := range lines {
		switch line[0:4] {
		case "$ ls":
			continue
		case "dir ":
			continue
		case "$ cd":
			targetDir := parseCd(line)
			switch targetDir {
			case "/":
				presentDirId = 0
				parentDirId = 0
			case "..":
				presentDirId = dirTree[presentDirId]
				parentDirId = dirTree[parentDirId]
			default:
				dirCount++
				parentDirId = presentDirId
				presentDirId = dirCount
				dirTree[presentDirId] = parentDirId
			}
		default:
			files = append(files, parseFile(presentDirId, line))
		}
	}

	return dirTree, files
}

func calcDirSize(dirId int, dirTree map[int]int, files []File) int {
	size := 0

	// get size of files contained directly
	for _, file := range files {
		if file.DirId == dirId {
			size += file.Size
		}
	}

	// repeat for each child directory
	for newDirId, parentDirId := range dirTree {
		if parentDirId == dirId {
			size += calcDirSize(newDirId, dirTree, files)
		}
	}

	return size
}

func calcDirSizes(dirTree map[int]int, files []File) map[int]int {
	dirSizes := map[int]int{}

	for dirId := range dirTree {
		if dirId != 0 { // skip root directory for now. we can simply sum all files to get it.
			dirSizes[dirId] = calcDirSize(dirId, dirTree, files)
		}
	}

	totalSize := 0
	for _, file := range files {
		totalSize += file.Size
	}
	dirSizes[0] = totalSize

	return dirSizes
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	dirTree, files := parseCommands(input)
	dirSizes := calcDirSizes(dirTree, files)

	solution := 0
	for _, size := range dirSizes {
		if size <= 100000 {
			solution += size
		}
	}

	fmt.Println(solution)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	dirTree, files := parseCommands(input)
	dirSizes := calcDirSizes(dirTree, files)

	totalSpace := 70000000
	desiredUnusedSpace := 30000000
	totalUsedSpace := dirSizes[0]
	totalUnusedSpace := totalSpace - totalUsedSpace
	spaceToDelete := desiredUnusedSpace - totalUnusedSpace

	minimumViableDirSize := totalUsedSpace
	for _, size := range dirSizes {
		if size >= spaceToDelete && size < minimumViableDirSize {
			minimumViableDirSize = size
		}
	}

	fmt.Println(minimumViableDirSize)
}
