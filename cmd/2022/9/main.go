package main

import (
	"aoc2022/internal/common"
	"aoc2022/internal/day"
	"aoc2022/pkg/data_structures/set"
	"aoc2022/pkg/point"
	"bytes"
	"fmt"
	"strconv"
)

type Step struct {
	Direction rune
	Length    int
}

var (
	deltas = map[rune]point.Point2D{
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

func setNewPos(a, b *point.Point2D) {
	diffX, diffY := a.X-b.X, a.Y-b.Y
	switch {
	case diffX > 0:
		b.X += 1
	case diffX < 0:
		b.X -= 1
	}
	switch {
	case diffY > 0:
		b.Y += 1
	case diffY < 0:
		b.Y -= 1
	}
}

func cmd(input []byte) (string, error) {
	steps := parseInput(input)
	coverage := set.New[point.Point2D]()

	head := point.Point2D{X: 0, Y: 0}
	tail := point.Point2D{X: 0, Y: 0}
	coverage.Add(tail)

	for _, step := range steps {
		delta := deltas[step.Direction]
		for i := 0; i < step.Length; i++ {
			head.X += delta.X
			head.Y += delta.Y

			if tail.InReach(&head, 1) {
				continue
			}

			setNewPos(&head, &tail)
			coverage.Add(tail)
		}
	}

	return fmt.Sprint(coverage.Len()), nil
}

func cmd2(input []byte) (string, error) {
	knots := 10
	steps := parseInput(input)
	coverage := set.New[point.Point2D]()

	points := make([]point.Point2D, 10)
	coverage.Add(points[knots-1])

	for _, step := range steps {
		toMove := deltas[step.Direction]
		for i := 0; i < step.Length; i++ {
			points[0].X += toMove.X
			points[0].Y += toMove.Y

			for i := 0; i < 9; i++ {
				lead, follower := points[i], &points[i+1]
				if follower.InReach(&lead, 1) {
					continue
				}
				setNewPos(&lead, follower)
			}
			coverage.Add(points[knots-1])
		}
	}

	return fmt.Sprint(coverage.Len()), nil
}
