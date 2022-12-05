package main

import (
	"aoc2022/internal/day"
	"aoc2022/pkg/chunks"
	"aoc2022/pkg/data_structures/stack"
	"strconv"
	"strings"
)

func main() {
	day := day.New("Day 5: Supply Stacks", cmd, cmd2)
	day.Run()
}

func getStacks() []*stack.Stack[string] {
	//Went faster to hardcode it... 😪
	return []*stack.Stack[string]{
		stack.NewFromArr([]string{"R", "N", "P", "G"}),
		stack.NewFromArr([]string{"T", "J", "B", "L", "C", "S", "V", "H"}),
		stack.NewFromArr([]string{"T", "D", "B", "M", "N", "L"}),
		stack.NewFromArr([]string{"R", "V", "P", "S", "B"}),
		stack.NewFromArr([]string{"G", "C", "Q", "S", "W", "M", "V", "H"}),
		stack.NewFromArr([]string{"W", "Q", "S", "C", "D", "B", "J"}),
		stack.NewFromArr([]string{"F", "Q", "L"}),
		stack.NewFromArr([]string{"W", "M", "H", "T", "D", "L", "F", "V"}),
		stack.NewFromArr([]string{"L", "P", "B", "V", "M", "J", "F"}),
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
		arr := make([]int, instruction[0])
		chunks := chunks.Create(arr, instruction[0])
		for i := 0; i < len(chunks); i++ {
			toMove := make([]string, len(chunks[i]))
			for j := 0; j < len(chunks[i]); j++ {
				toMove[j] = stacks[instruction[1]-1].Pop()
			}
			internalStack := stack.NewFromArr(toMove)
			len := internalStack.Len()
			for k := 0; k < len; k++ {
				stacks[instruction[2]-1].Push(internalStack.Pop())
			}
		}
	}

	sb := strings.Builder{}
	for _, stack := range stacks {
		sb.WriteString(stack.Pop())
	}

	return sb.String(), nil
}
