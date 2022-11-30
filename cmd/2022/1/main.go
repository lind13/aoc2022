package main

import (
	"aoc2022/internal/day"
	"errors"
)

func main() {
	day := day.New("Day 1", cmd, cmd2)
	day.Run()
}

func cmd(input string) (string, error) {
	return "1", nil
}

func cmd2(input string) (string, error) {
	return "", errors.New("NOT IMPLEMENTED")
}
