package day

import (
	"aoc2022/internal/file"
	"fmt"
	"time"
)

type Cmd func(input string) (string, error)

type Day struct {
	id   string
	cmds []Cmd
}

func New(id string, cmd ...Cmd) *Day {
	return &Day{
		id:   id,
		cmds: cmd,
	}
}

func (d *Day) Run() {
	fmt.Printf("\u001b[36;1m-=< ❄️ >=- %s -=< ❄️ >=- 🎅\n", d.id)

	input, err := file.ReadInput()
	if err != nil {
		panic(err)
	}

	for i, cmd := range d.cmds {
		stringInput := string(input)
		start := time.Now()
		result, err := cmd(stringInput)
		duration := time.Since(start)
		if err != nil {
			fmt.Printf("%d: %v\n", i+1, err)
			continue
		}
		fmt.Printf("%d: [time: %s, result: %s]\n", i+1, duration, result)
	}
}
