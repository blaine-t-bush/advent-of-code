package day18

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile     = "./2022/day18/input.txt"
	maxIterations = 1000
)

type cube struct {
	x int
	y int
	z int
}

func (c cube) left() cube {
	return cube{c.x - 1, c.y, c.z}
}

func (c cube) right() cube {
	return cube{c.x + 1, c.y, c.z}
}

func (c cube) above() cube {
	return cube{c.x, c.y, c.z + 1}
}

func (c cube) below() cube {
	return cube{c.x, c.y, c.z - 1}
}

func (c cube) forward() cube {
	return cube{c.x, c.y + 1, c.z}
}

func (c cube) backward() cube {
	return cube{c.x, c.y - 1, c.z}
}

func parseCubes(lines []string) []cube {
	cubes := []cube{}
	for _, line := range lines {
		cubesStrings := strings.Split(line, ",")
		if len(cubesStrings) != 3 {
			log.Fatal("parseCubes: did not parse 3 values")
		}

		x, err := strconv.Atoi(cubesStrings[0])
		util.CheckErr(err)

		y, err := strconv.Atoi(cubesStrings[1])
		util.CheckErr(err)

		z, err := strconv.Atoi(cubesStrings[2])
		util.CheckErr(err)

		cubes = append(cubes, cube{x, y, z})
	}

	return cubes
}

func (c1 cube) countConnections(cubes []cube) int {
	count := 0
	for _, c2 := range cubes {
		if c1.x == c2.x && c1.y == c2.y && util.AbsInt(c1.z-c2.z) == 1 {
			count++
		} else if c1.x == c2.x && util.AbsInt(c1.y-c2.y) == 1 && c1.z == c2.z {
			count++
		} else if util.AbsInt(c1.x-c2.x) == 1 && c1.y == c2.y && c1.z == c2.z {
			count++
		}
	}

	return count
}

func createMap(cubes []cube) map[cube]bool {
	cubeMap := map[cube]bool{}
	for _, c := range cubes {
		cubeMap[c] = true
	}

	return cubeMap
}

func (c cube) isInPocket(cubeMap map[cube]bool) bool {
	isSolidAbove, isSolidBelow, isSolidRight, isSolidLeft, isSolidForward, isSolidBackward := false, false, false, false, false, false

	if cubeMap[c] {
		return false
	}

	newC := c
	for i := 0; i < maxIterations; i++ {
		newC = newC.above()
		if cubeMap[newC] {
			isSolidAbove = true
			break
		}
	}

	newC = c
	for i := 0; i < maxIterations; i++ {
		newC = newC.below()
		if cubeMap[newC] {
			isSolidBelow = true
			break
		}
	}

	newC = c
	for i := 0; i < maxIterations; i++ {
		newC = newC.right()
		if cubeMap[newC] {
			isSolidRight = true
			break
		}
	}

	newC = c
	for i := 0; i < maxIterations; i++ {
		newC = newC.left()
		if cubeMap[newC] {
			isSolidLeft = true
			break
		}
	}

	newC = c
	for i := 0; i < maxIterations; i++ {
		newC = newC.forward()
		if cubeMap[newC] {
			isSolidForward = true
			break
		}
	}

	newC = c
	for i := 0; i < maxIterations; i++ {
		newC = newC.backward()
		if cubeMap[newC] {
			isSolidBackward = true
			break
		}
	}

	return isSolidAbove && isSolidBelow && isSolidLeft && isSolidRight && isSolidForward && isSolidBackward
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	cubes := parseCubes(input)
	openFaces := len(cubes) * 6
	for _, c := range cubes {
		openFaces -= c.countConnections(cubes)
	}

	fmt.Println(openFaces)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	cubes := parseCubes(input)
	openFaces := len(cubes) * 6
	for _, c := range cubes {
		openFaces -= c.countConnections(cubes)
	}

	cubeMap := createMap(cubes)
	minX, minY, minZ := 1000, 1000, 1000
	maxX, maxY, maxZ := -1000, -1000, -1000
	for _, c := range cubes {
		if c.x < minX {
			minX = c.x
		} else if c.x > maxX {
			maxX = c.x
		}

		if c.y < minY {
			minY = c.y
		} else if c.y > maxY {
			maxY = c.y
		}

		if c.z < minZ {
			minZ = c.z
		} else if c.z > maxZ {
			maxZ = c.z
		}
	}

	trappedCubes := []cube{}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				c := cube{x, y, z}
				if c.isInPocket(cubeMap) {
					trappedCubes = append(trappedCubes, c)
				}
			}
		}
	}

	filled := append(cubes, trappedCubes...)
	filledOpenFaces := len(filled) * 6
	for _, c := range filled {
		filledOpenFaces -= c.countConnections(filled)
	}

	// answer is greater that 2510
	fmt.Println(filledOpenFaces)
}
