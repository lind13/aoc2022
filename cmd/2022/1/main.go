package main

import (
	"aoc2022/internal/day"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	day := day.New("Day 1: Calorie Counting", cmd, cmd2)
	day.Run()
}

func cmd(input []byte) (string, error) {
	inputArr := strings.Split(string(input), "\n")
	calorieElfMap := make(map[int]int)

	elf := 1
	for _, calStr := range inputArr {
		if calStr == "" {
			elf++
			continue
		}

		cal, err := strconv.Atoi(calStr)
		if err != nil {
			return "", err
		}

		calorieElfMap[elf] += cal
	}

	count := 0
	for _, v := range calorieElfMap {
		if count == 0 {
			count = v
		}

		if v > count {
			count = v
		}
	}

	return fmt.Sprint(count), nil
}

func cmd2(input []byte) (string, error) {
	inputArr := strings.Split(string(input), "\n")
	calorieElfMap := make(map[int]int)

	elf := 1
	for _, calStr := range inputArr {
		if calStr == "" {
			elf++
			continue
		}

		cal, err := strconv.Atoi(calStr)
		if err != nil {
			return "", err
		}

		calorieElfMap[elf] += cal
	}

	foods := make([]int, len(calorieElfMap))
	for i := 0; i < len(calorieElfMap); i++ {
		foods[i] = calorieElfMap[i+1]
	}

	sort.Slice(foods, func(i, j int) bool {
		return foods[i] < foods[j]
	})

	top3 := foods[len(foods)-3:]

	sum := 0
	for _, v := range top3 {
		sum += v
	}

	return fmt.Sprint(sum), nil
}
