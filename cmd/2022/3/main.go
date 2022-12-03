package main

import (
	"aoc2022/internal/day"
	"aoc2022/pkg/chunks"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	day := day.New("Day 3: Rucksack Reorganization", cmd, cmd2)
	day.Run()
}

func cmd(input string) (string, error) {
	rucksacks := strings.Split(input, "\n")
	points := make([]int, 0, len(rucksacks))

mainLoop:
	for _, rucksack := range rucksacks {
		c1, c2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
		for _, x := range c1 {
			for _, y := range c2 {
				if x == y {
					points = append(points, toPrio(x))
					continue mainLoop
				}
			}
		}
	}

	sum := 0
	for _, x := range points {
		sum += x
	}

	return fmt.Sprint(sum), nil
}

func toPrio(r rune) int {
	if unicode.IsLower(r) {
		return int(r - 96)
	}
	return int(r - 65 + 27)
}

func cmd2(input string) (string, error) {
	groups := chunks.Create(strings.Split(input, "\n"), 3)
	badgePoints := make([]int, 0, len(groups))

	for _, group := range groups {
		for _, r := range group[0] {
			if strings.ContainsRune(group[1], r) && strings.ContainsRune(group[2], r) {
				badgePoints = append(badgePoints, toPrio(r))
				break
			}
		}
	}

	sum := 0
	for _, x := range badgePoints {
		sum += x
	}

	return fmt.Sprint(sum), nil
}
