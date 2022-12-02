package main

import (
	"aoc2022/internal/day"
	"fmt"
	"strings"
)

const (
	X = "X"
	Y = "Y"
	Z = "Z"

	A = "A"
	B = "B"
	C = "C"

	LOSS = 0
	DRAW = 3
	WIN  = 6
)

var (
	norm = map[string]int{
		X: 1, //Rock
		Y: 2, //Paper
		Z: 3, //Scissors
		A: 1,
		B: 2,
		C: 3,
	}
)

func main() {
	day := day.New("Day 2: Rock Paper Scissors", cmd, cmd2)
	day.Run()
}

func cmd(input string) (string, error) {
	rounds := strings.Split(input, "\n")
	score := 0
	for _, round := range rounds {
		picks := strings.Split(round, " ")
		opp := norm[picks[0]]
		you := norm[picks[1]]
		switch you {
		case 1:
			if opp == 1 {
				score += DRAW + you
			}
			if opp == 2 {
				score += LOSS + you
			}
			if opp == 3 {
				score += WIN + you
			}
		case 2:
			if opp == 2 {
				score += DRAW + you
			}
			if opp == 1 {
				score += WIN + you
			}
			if opp == 3 {
				score += LOSS + you
			}
		case 3:
			if opp == 3 {
				score += DRAW + you
			}
			if opp == 2 {
				score += WIN + you
			}
			if opp == 1 {
				score += LOSS + you
			}
		}
	}

	return fmt.Sprint(score), nil
}

func cmd2(input string) (string, error) {
	rounds := strings.Split(input, "\n")
	score := 0
	for _, round := range rounds {
		picks := strings.Split(round, " ")
		opp := norm[picks[0]]
		predictedResult := norm[picks[1]]

		switch predictedResult {
		case 1:
			if opp == 1 {
				score += LOSS + 3
			}
			if opp == 2 {
				score += LOSS + 1
			}
			if opp == 3 {
				score += LOSS + 2
			}
		case 2:
			if opp == 2 {
				score += DRAW + 2
			}
			if opp == 1 {
				score += DRAW + 1
			}
			if opp == 3 {
				score += DRAW + 3
			}
		case 3:
			if opp == 3 {
				score += WIN + 1
			}
			if opp == 2 {
				score += WIN + 3
			}
			if opp == 1 {
				score += WIN + 2
			}
		}
	}

	return fmt.Sprint(score), nil
}
