package main

import "aoc2022/internal/day"

func main() {
	day := day.New("Day 9", cmd)
	day.Run()
}

func cmd(input string) (string, error) {
	return input, nil
}
