package main

import (
	"aoc2022/internal/common"
	"aoc2022/internal/day"
	"aoc2022/pkg/data_structures/set"
	"bytes"
	"fmt"
	"strconv"
)

type Step struct {
	Direction rune
	Length    int
}

var (
	deltas = map[rune]common.Point2D{
		'U': {X: 0, Y: 1},
		'D': {X: 0, Y: -1},
		'R': {X: 1, Y: 0},
		'L': {X: -1, Y: 0},
	}
)

func main() {
	day := day.New("Day 9: Rope Bridge", cmd, cmd2)
	day.Run()
}

func parseInput(input []byte) []Step {
	steps := make([]Step, 0)
	rows := bytes.Split(input, common.NL)
	for _, row := range rows {
		step := bytes.Split(row, []byte(" "))
		l, _ := strconv.Atoi(string(step[1]))
		steps = append(steps, Step{
			Direction: bytes.Runes(step[0])[0],
			Length:    l,
		})
	}
	return steps
}

func isAdjacent(head, tail common.Point2D) bool {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if tail.X+dx == head.X && tail.Y+dy == head.Y {
				return true
			}
		}
	}
	return false
}

func setNewPos(head, tail *common.Point2D) {
	diffX := head.X - tail.X
	diffY := head.Y - tail.Y
	switch {
	case diffX > 0:
		tail.X += 1
	case diffX < 0:
		tail.X -= 1
	}
	switch {
	case diffY > 0:
		tail.Y += 1
	case diffY < 0:
		tail.Y -= 1
	}
}

func cmd(input []byte) (string, error) {
	steps := parseInput(input)
	coverage := set.New[common.Point2D]()

	head := common.Point2D{X: 0, Y: 0}
	tail := common.Point2D{X: 0, Y: 0}
	coverage.Add(tail)

	for _, step := range steps {
		delta := deltas[step.Direction]
		for i := 0; i < step.Length; i++ {
			head.X += delta.X
			head.Y += delta.Y

			if isAdjacent(head, tail) {
				continue
			}

			setNewPos(&head, &tail)
			coverage.Add(tail)
		}
	}

	return fmt.Sprint(coverage.Len()), nil
}

func cmd2(input []byte) (string, error) {
	steps := parseInput(input)
	coverage := set.New[common.Point2D]()

	points := make([]common.Point2D, 10)
	coverage.Add(points[9])

	for _, step := range steps {
		toMove := deltas[step.Direction]
		for i := 0; i < step.Length; i++ {
			points[0].X += toMove.X
			points[0].Y += toMove.Y

			for i := 0; i < 9; i++ {
				lead, follower := points[i], &points[i+1]
				if isAdjacent(lead, *follower) {
					continue
				}
				setNewPos(&lead, follower)
			}
			coverage.Add(points[9])
		}
	}

	return fmt.Sprint(coverage.Len()), nil
}
