package file

import (
	"os"
)

const (
	inputFile = "./data/input.txt"
)

func ReadInput() ([]byte, error) {
	return os.ReadFile(inputFile)
}
