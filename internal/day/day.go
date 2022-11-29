package day

import (
	"aoc2022/internal/file"
	"fmt"
	"time"

	tm "github.com/buger/goterm"
)

type Cmd func(input string) (string, error)

type Day struct {
	id   string
	cmds []Cmd
}

type state uint8

const (
	idle state = iota
	running
	done
	failure
)

func (s state) String() string {
	switch s {
	case idle:
		return "idle"
	case running:
		return "running..."
	case done:
		return "done"
	case failure:
		return "error"
	default:
		return "unknown state"
	}
}

type cmdResult struct {
	state  state
	result string
	time   time.Duration
}

func New(id string, cmd ...Cmd) *Day {
	return &Day{
		id:   id,
		cmds: cmd,
	}
}

func (d *Day) Run() {
	tm.Clear()
	input, err := file.ReadInput()
	if err != nil {
		panic(err)
	}

	res := make([]cmdResult, len(d.cmds))
	for i := 0; i < len(d.cmds); i++ {
		res[i] = cmdResult{
			state:  idle,
			result: "",
			time:   0,
		}
	}

	d.draw(res)

	for i, cmd := range d.cmds {
		res[i] = cmdResult{
			state: running,
		}
		d.draw(res)
		start := time.Now()
		result, err := cmd(string(input))
		time.Sleep(time.Second * 2)
		if err != nil {
			res[i] = cmdResult{
				state: failure,
			}
			d.draw(res)
			continue
		}
		duration := time.Since(start)

		res[i] = cmdResult{
			state:  done,
			result: result,
			time:   duration,
		}
		d.draw(res)
	}
}

func (d *Day) draw(results []cmdResult) {
	tm.MoveCursor(1, 1)
	tm.Println(tm.Bold(tm.Color(d.id, tm.CYAN)))
	totals := tm.NewTable(0, 10, 5, ' ', 1)
	fmt.Fprint(totals, "Command\tTime\tResult\n")
	for i := 0; i < len(results); i++ {
		result := results[i]
		switch result.state {
		case idle:
			fallthrough
		case running:
			fmt.Fprintf(totals, "%d\t%s\n", i+1, result.state)
		case done:
			fmt.Fprintf(totals, "%d\t%s\t%s\n", i+1, result.time, result.result)
		}
	}
	tm.Print(totals)
	tm.Flush()
}
