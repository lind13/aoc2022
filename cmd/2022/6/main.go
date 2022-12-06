package main

import (
	"aoc2022/internal/day"
	"bytes"
	"fmt"
)

func main() {
	day := day.New("Day 6: Tuning Trouble", cmd, cmd2)
	day.Run()
}

func getIndex(input string, n int) int {
	buf := bytes.NewBufferString(input)
	sequence := make([]byte, 0, n)
	for i := 1; true; i++ {
		b := buf.Next(1)
		if len(b) == 0 {
			return -1
		}

		if !bytes.Contains(sequence, b) {
			sequence = append(sequence, b...)
			if len(sequence) == n {
				return i
			}
		} else {
			s := bytes.IndexByte(sequence, b[0])
			sequence = append(sequence[s+1:], b...)
		}
	}
	return -1
}

func cmd(input string) (string, error) {
	return fmt.Sprint(getIndex(input, 4)), nil
}

func cmd2(input string) (string, error) {
	return fmt.Sprint(getIndex(input, 14)), nil
}
