package day9

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	inputFile = "./2022/day9/input.txt"
	dirUp     = iota
	dirRight
	dirDown
	dirLeft
)

type coord struct {
	x int
	y int
}

type rope struct {
	knots             map[int]coord
	tailVisitedCoords []coord
}

type command struct {
	dir   int
	steps int
}

func (r rope) len() int {
	return len(r.knots)
}

func (r rope) knotToFollowerDistance(knotIndex int) int {
	head := r.knots[knotIndex]
	tail := r.knots[knotIndex+1]
	if head.x == tail.x {
		return util.AbsInt(head.y - tail.y)
	} else if head.y == tail.y {
		return util.AbsInt(head.x - tail.x)
	} else {
		// number of diagonal units is equal to the smaller of the x and y distances
		// the (abs) difference of the x and y distances is the number of straight units
		// the sum of those two values is the total distance
		xDist := util.AbsInt(head.x - tail.x)
		yDist := util.AbsInt(head.y - tail.y)
		if xDist > yDist {
			return util.MinInts(xDist, yDist) + util.AbsInt(xDist-yDist)
		} else {
			return util.MinInts(xDist, yDist) + util.AbsInt(xDist-yDist)
		}
	}
}

// func (r rope) isValid(knotIndex int) bool {
// 	return r.knotToFollowerDistance(knotIndex) <= 1
// }

func (r *rope) move(dir, steps int) {
	for i := 0; i < steps; i++ {
		// first move head according to direction
		r.moveHead(dir)

		// then move following knots
		for ii := 1; ii < r.len(); ii++ {
			r.moveFollower(ii)
		}

		// update list of visited coords
		r.addTailVisitedCoord()
	}
}

func (r *rope) moveHead(dir int) {
	head := r.knots[0]
	switch dir {
	case dirUp:
		head.y += 1
	case dirDown:
		head.y -= 1
	case dirRight:
		head.x += 1
	case dirLeft:
		head.x -= 1
	default:
		log.Fatal("invalid rope movement direction")
	}

	r.knots[0] = head
}

func (r *rope) moveFollower(followerIndex int) {
	moveDist := r.knotToFollowerDistance(followerIndex-1) - 1
	if moveDist == 0 {
		return
	}

	head := r.knots[followerIndex-1]
	tail := r.knots[followerIndex]
	xDist := head.x - tail.x
	yDist := head.y - tail.y

	// if in same column or row, move dist along that line
	if xDist == 0 && yDist < 0 {
		// move down
		tail.y -= moveDist
	} else if xDist == 0 && yDist > 0 {
		// move up
		tail.y += moveDist
	} else if xDist < 0 && yDist == 0 {
		// move left
		tail.x -= moveDist
	} else if xDist > 0 && yDist == 0 {
		// move right
		tail.x += moveDist
	}

	// if not in same column or row, move dist along diagonal
	if xDist > 0 && yDist > 0 {
		// move up right
		tail.y += moveDist
		tail.x += moveDist
	} else if xDist > 0 && yDist < 0 {
		// move down right
		tail.y -= moveDist
		tail.x += moveDist
	} else if xDist < 0 && yDist > 0 {
		// move up left
		tail.y += moveDist
		tail.x -= moveDist
	} else if xDist < 0 && yDist < 0 {
		// move down left
		tail.y -= moveDist
		tail.x -= moveDist
	}

	r.knots[followerIndex-1] = head
	r.knots[followerIndex] = tail
}

func (r *rope) addTailVisitedCoord() {
	tail := r.knots[r.len()-1]
	for _, coord := range r.tailVisitedCoords {
		if coord.x == tail.x && coord.y == tail.y {
			return
		}
	}

	r.tailVisitedCoords = append(r.tailVisitedCoords, coord{tail.x, tail.y})
}

func parseCommand(line string) command {
	r, err := regexp.Compile(`(U|D|L|R) (\d+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 3 {
		log.Fatal("could not parse command")
	}

	var dir int
	switch m[1] {
	case "U":
		dir = dirUp
	case "D":
		dir = dirDown
	case "L":
		dir = dirLeft
	case "R":
		dir = dirRight
	default:
		log.Fatal("could not parse command direction")
	}

	steps, err := strconv.Atoi(m[2])
	util.CheckErr(err)

	return command{
		dir:   dir,
		steps: steps,
	}
}

func parseCommands(lines []string) []command {
	commands := make([]command, len(lines))
	for i, line := range lines {
		commands[i] = parseCommand(line)
	}

	return commands
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	commands := parseCommands(input)
	r := rope{
		knots: map[int]coord{
			0: {0, 0},
			1: {0, 0},
			2: {0, 0},
			3: {0, 0},
			4: {0, 0},
			5: {0, 0},
			6: {0, 0},
			7: {0, 0},
			8: {0, 0},
			9: {0, 0},
		},
		tailVisitedCoords: []coord{
			{0, 0},
		},
	}

	for _, command := range commands {
		r.move(command.dir, command.steps)
	}

	fmt.Println(len(r.tailVisitedCoords))

}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
