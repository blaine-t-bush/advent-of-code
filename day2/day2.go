package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	MoveRock Move = iota
	MovePaper
	MoveScissors
)

type Move int

type Round struct {
	Count int
	Byte1 byte
	Byte2 byte
	Opponent Move
	Player Move
}

func (m Move) getWinningMove() Move {
	switch m {
	case MoveRock:
		return MovePaper
	case MovePaper:
		return MoveScissors
	case MoveScissors:
		return MoveRock
	default:
		log.Fatal("failed to determine winning move")
		return 0
	}
}

func (m Move) getLosingMove() Move {
	switch m {
	case MoveRock:
		return MoveScissors
	case MovePaper:
		return MoveRock
	case MoveScissors:
		return MovePaper
	default:
		log.Fatal("failed to determine losing move")
		return 0
	}
}

func (m Move) getDrawingMove() Move {
	return m
}

func (r *Round) getScore() int {
	score := 0

	if r.Opponent == MoveRock && r.Player == MovePaper || r.Opponent == MovePaper && r.Player == MoveScissors || r.Opponent == MoveScissors && r.Player == MoveRock {
		score += 6
	} else if r.Player == r.Opponent {
		score += 3
	}

	switch r.Player {
	case MoveRock:
		score += 1
	case MovePaper:
		score += 2
	case MoveScissors:
		score += 3
	}

	return score
}

func stringToMove(char byte) Move {
	switch char {
	case 'A':
		return MoveRock
	case 'B':
		return MovePaper
	case 'C':
		return MoveScissors
	case 'X':
		return MoveRock
	case 'Y':
		return MovePaper
	case 'Z':
		return MoveScissors
	default:
		log.Fatal("failed to convert string to move")
		return 0
	}
}

func stringToStrategy(opponent Move, char byte) Move {
	switch char {
	case 'X':
		return opponent.getLosingMove()
	case 'Y':
		return opponent.getDrawingMove()
	case 'Z':
		return opponent.getWinningMove()
	default:
		log.Fatal("failed to convert string to strategy")
		return 0
	}
}

func getInput() []Round {
	f, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	roundCount := 0
	rounds := []Round{}
	for scanner.Scan() {
		text := scanner.Text()
		roundCount++
		rounds = append(rounds, Round{Count: roundCount, Byte1: text[0], Byte2: text[2]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rounds
}

func SolvePartOne() int {
	score := 0
	for _, round := range getInput() {
		round.Opponent = stringToMove(round.Byte1)
		round.Player = stringToMove(round.Byte2)
		score += round.getScore()
	}

	fmt.Println(score)
	return score
}

func SolvePartTwo() int {
	score := 0
	for _, round := range getInput() {
		round.Opponent = stringToMove(round.Byte1)
		round.Player = stringToStrategy(round.Opponent, round.Byte2)
		score += round.getScore()
	}

	fmt.Println(score)
	return score
}