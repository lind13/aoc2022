package day

import (
	"aoc2022/internal/file"
	"fmt"
	"time"
)

type Cmd func(input []byte) (string, error)

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
		start := time.Now()
		result, err := cmd(input)
		duration := time.Since(start)
		if err != nil {
			fmt.Printf("%d: %v\n", i+1, err)
			continue
		}
		fmt.Printf("%d: [time: %s, result: %s]\n", i+1, duration, result)
	}
}
