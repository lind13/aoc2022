package main

import (
	"aoc2022/internal/day"
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

const (
	TOTAL_SPACE   = 70000000
	NEEDED_UNUSED = 30000000
)

var (
	CMD  = []byte("$")
	CD   = []byte("cd")
	OUT  = []byte("..")
	LS   = []byte("ls")
	ROOT = []byte("/")
	DIR  = []byte("dir")
)

type FileInfo struct {
	Name []byte
	Size int
}

type Directory struct {
	Path     [][]byte
	Files    []*FileInfo
	Children []*Directory
	Parent   *Directory
}

func main() {
	day := day.New("Day 7: No Space Left On Device", cmd, cmd2)
	day.Run()
}

func preProcess(input []byte) [][]byte {
	return bytes.Split(input, []byte("\n"))
}

func createTree(rows [][]byte) *Directory {
	path := [][]byte{}
	var root *Directory = nil
	var currentNode *Directory = nil

	ls := false
	for _, row := range rows {
		w := bytes.Split(row, []byte(" "))
		if bytes.Equal(w[0], CMD) {
			ls = false
			switch {
			case bytes.Equal(w[1], CD):
				switch {
				case bytes.Equal(w[2], ROOT):
					path = [][]byte{[]byte("root")}
					root = &Directory{
						Path:   path,
						Parent: nil,
					}
					currentNode = root
					continue
				case bytes.Equal(w[2], OUT):
					currentNode = currentNode.Parent
					path = path[:len(path)-1]
					continue
				default:
					path = append(path, w[2])
					nextNode := &Directory{
						Path:   path,
						Parent: currentNode,
					}
					currentNode.Children = append(currentNode.Children, nextNode)
					currentNode = nextNode
					continue
				}
			case bytes.Equal(w[1], LS):
				ls = true
				continue
			}
		} else if ls {
			switch {
			case !bytes.Equal(w[0], DIR):
				size, _ := strconv.Atoi(string(w[0]))
				fi := &FileInfo{
					Name: w[1],
					Size: size,
				}
				currentNode.Files = append(currentNode.Files, fi)
			}
		}
	}
	return root
}

func getTotalSize(node *Directory, size *int) {
	for _, file := range node.Files {
		*size += file.Size
	}
	for _, child := range node.Children {
		getTotalSize(child, size)
	}
}

func walk(node *Directory, sizes *[]int) {
	size := 0
	getTotalSize(node, &size)

	*sizes = append(*sizes, size)
	for _, child := range node.Children {
		walk(child, sizes)
	}
}

func cmd(input []byte) (string, error) {
	rows := preProcess(input)
	root := createTree(rows)
	totalSizes := make([]int, 0, len(rows))
	walk(root, &totalSizes)

	sum := 0
	for _, v := range totalSizes {
		if v < 100001 {
			sum += v
		}
	}

	return fmt.Sprint(sum), nil
}

func cmd2(input []byte) (string, error) {
	rows := preProcess(input)
	root := createTree(rows)
	sizes := make([]int, 0, len(rows))
	walk(root, &sizes)

	usedSpace := 0
	getTotalSize(root, &usedSpace)
	unusedSpace := TOTAL_SPACE - usedSpace

	arr := make([]int, 0, len(sizes))
	for _, size := range sizes {
		if unusedSpace+size >= NEEDED_UNUSED {
			arr = append(arr, size)
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return fmt.Sprint(arr[0]), nil
}
