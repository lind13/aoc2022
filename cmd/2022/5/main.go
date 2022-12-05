package main

import (
	"aoc2022/internal/day"
	"aoc2022/pkg/data_structures/stack"
	"strconv"
	"strings"
)

func main() {
	day := day.New("Day 5: Supply Stacks", cmd, cmd2)
	day.Run()
}

func getStacks() []stack.Stack[string] {
	//Went faster to hardcode it... 😪
	return []stack.Stack[string]{
		{"R", "N", "P", "G"},
		{"T", "J", "B", "L", "C", "S", "V", "H"},
		{"T", "D", "B", "M", "N", "L"},
		{"R", "V", "P", "S", "B"},
		{"G", "C", "Q", "S", "W", "M", "V", "H"},
		{"W", "Q", "S", "C", "D", "B", "J"},
		{"F", "Q", "L"},
		{"W", "M", "H", "T", "D", "L", "F", "V"},
		{"L", "P", "B", "V", "M", "J", "F"},
	}
}

func processInput(input string) [][3]int {
	moves := make([][3]int, 0)
	rows := strings.Split(strings.Split(input, "\n\n")[1], "\n")
	for _, row := range rows {
		parts := strings.Split(row, " ")
		c, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		moves = append(moves, [3]int{c, from, to})
	}
	return moves
}

func cmd(input string) (string, error) {
	stacks := getStacks()
	instructions := processInput(input)

	for _, instruction := range instructions {
		for i := 0; i < instruction[0]; i++ {
			stacks[instruction[2]-1].Push(stacks[instruction[1]-1].Pop())
		}
	}

	sb := strings.Builder{}
	for _, stack := range stacks {
		sb.WriteString(stack.Pop())
	}

	return sb.String(), nil
}

func cmd2(input string) (string, error) {
	stacks := getStacks()
	instructions := processInput(input)

	for _, instruction := range instructions {
		stacks[instruction[2]-1].PushN(stacks[instruction[1]-1].PopN(instruction[0]))
	}

	sb := strings.Builder{}
	for _, stack := range stacks {
		sb.WriteString(stack.Pop())
	}

	return sb.String(), nil
}
