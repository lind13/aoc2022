package main

import (
	"aoc2022/internal/day"
	"bytes"
	"fmt"
)

func main() {
	day := day.New("Day 8: Treetop Tree House", cmd, cmd2)
	day.Run()
}

func parseInput(input []byte) [][][]byte {
	grid := make([][][]byte, 0)
	rows := bytes.Split(input, []byte("\n"))
	for _, row := range rows {
		grid = append(grid, bytes.Split(row, []byte("")))
	}
	return grid
}

func isBlocked(y, x int, grid [][][]byte) bool {
	tree := grid[y][x]
	blocked := 0
	for t := y - 1; t >= 0; t-- {
		if bytes.Compare(grid[t][x], tree) > -1 {
			blocked++
			break
		}
	}

	for b := y + 1; b < len(grid); b++ {
		if bytes.Compare(grid[b][x], tree) > -1 {
			blocked++
			break
		}
	}

	for r := x + 1; r < len(grid[y]); r++ {
		if bytes.Compare(grid[y][r], tree) > -1 {
			blocked++
			break
		}
	}

	for l := x - 1; l >= 0; l-- {
		if bytes.Compare(grid[y][l], tree) > -1 {
			blocked++
			break
		}
	}

	return blocked == 4
}

func cmd(input []byte) (string, error) {
	grid := parseInput(input)
	visible := make([][]byte, 0, len(input))
	for y, row := range grid {
		for x, tree := range row {
			if y == 0 || x == 0 || y == len(grid)-1 || x == len(row)-1 {
				//edge
				visible = append(visible, tree)
				continue
			}

			if !isBlocked(y, x, grid) {
				visible = append(visible, tree)
			}
		}
	}
	return fmt.Sprint(len(visible)), nil
}

func sumVisibleTrees(y, x int, grid [][][]byte) int {
	tree := grid[y][x]
	b1, b2, b3, b4 := 0, 0, 0, 0
	for t := y - 1; t >= 0; t-- {
		if bytes.Compare(grid[t][x], tree) > -1 {
			b1++
			break
		}
		b1++
	}

	for b := y + 1; b < len(grid); b++ {
		if bytes.Compare(grid[b][x], tree) > -1 {
			b2++
			break
		}
		b2++
	}

	for r := x + 1; r < len(grid[y]); r++ {
		if bytes.Compare(grid[y][r], tree) > -1 {
			b3++
			break
		}
		b3++
	}

	for l := x - 1; l >= 0; l-- {
		if bytes.Compare(grid[y][l], tree) > -1 {
			b4++
			break
		}
		b4++
	}

	return b1 * b2 * b3 * b4
}

func cmd2(input []byte) (string, error) {
	grid := parseInput(input)
	max := 0
	for y, row := range grid {
		for x := range row {
			if y == 0 || x == 0 || y == len(grid)-1 || x == len(row)-1 {
				//edge
				continue
			}

			sum := sumVisibleTrees(y, x, grid)
			if sum > max {
				max = sum
			}
		}
	}
	return fmt.Sprint(max), nil
}
