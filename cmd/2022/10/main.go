package main

import (
	"aoc2022/internal/common"
	"aoc2022/internal/day"
	"bytes"
	"fmt"
	"strconv"
)

var (
	noop = []byte("noop")
	addx = []byte("addx")
)

func main() {
	day := day.New("Day 10: Cathode-Ray Tube", cmd, cmd2)
	day.Run()
}

func cmd(input []byte) (string, error) {
	var (
		cycle   = 0
		x       = 1
		signals = make([]int, 0)
	)

	checkSig := func() {
		if cycle != 0 && (cycle == 20 || (cycle-20)%40 == 0) {
			signals = append(signals, cycle*x)
		}
	}

	for _, row := range bytes.Split(input, common.NL) {
		w := bytes.Split(row, common.SP)
		switch {
		case bytes.Equal(w[0], noop):
			cycle++
			checkSig()
		case bytes.Equal(w[0], addx):
			for i := 0; i < 2; i++ {
				cycle++
				checkSig()
			}
			add, _ := strconv.Atoi(string(w[1]))
			x += add
		}
	}

	sum := 0
	for _, sig := range signals {
		sum += sig
	}

	return fmt.Sprint(sum), nil
}

func cmd2(input []byte) (string, error) {
	var (
		cycle = 0
		x     = 1
		line  = make([]string, 0, 40)
		lines = make([][]string, 0)
	)

	runCycle := func() {
		cycle++
		linePos := cycle - 40*len(lines)
		diff := linePos - x
		if diff > -1 && diff < 3 {
			line = append(line, "#")
		} else {
			line = append(line, ".")
		}

		if cycle != 0 && cycle%40 == 0 {
			lines = append(lines, line)
			line = make([]string, 0, 40)
		}
	}

	for _, row := range bytes.Split(input, common.NL) {
		w := bytes.Split(row, common.SP)
		switch {
		case bytes.Equal(w[0], noop):
			runCycle()
		case bytes.Equal(w[0], addx):
			for i := 0; i < 2; i++ {
				runCycle()
			}
			add, _ := strconv.Atoi(string(w[1]))
			x += add
		}
	}

	sum := 0
	for _, line := range lines {
		for _, s := range line {
			fmt.Print(s)
		}
		fmt.Println()
	}

	return fmt.Sprint(sum), nil
}
