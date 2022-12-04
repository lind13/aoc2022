package main

import (
	"aoc2022/internal/day"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	day := day.New("Day 4: Camp Cleanup", cmd, cmd2)
	day.Run()
}

func toPairs(input string) [][2][2]int {
	pairs := make([][2][2]int, 0)
	for _, row := range strings.Split(input, "\n") {
		pair := [2][2]int{{}, {}}
		elfs := strings.Split(row, ",")
		for i, elf := range elfs {
			workRange := strings.Split(elf, "-")
			start, _ := strconv.Atoi(workRange[0])
			stop, _ := strconv.Atoi(workRange[1])
			pair[i] = [2]int{start, stop}
		}

		pairs = append(pairs, pair)
	}
	return pairs
}

func isWithin(a, b [2]int) bool {
	if a[0] >= b[0] && a[1] <= b[1] {
		return true
	}
	if b[0] >= a[0] && b[1] <= a[1] {
		return true
	}
	return false
}

func cmd(input string) (string, error) {
	count := 0
	for _, pair := range toPairs(input) {
		if isWithin(pair[0], pair[1]) {
			count++
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func isOverlap(a, b [2]int) bool {
	if a[0] >= b[0] && a[0] <= b[1] {
		return true
	}
	if b[0] >= a[0] && b[0] <= a[1] {
		return true
	}
	return false
}

func cmd2(input string) (string, error) {
	count := 0
	for _, pair := range toPairs(input) {
		if isOverlap(pair[0], pair[1]) {
			count++
		}
	}

	return fmt.Sprintf("%d", count), nil
}
