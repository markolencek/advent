package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type program struct {
	visited   bool
	inGroup   int
	connected []*program
}

func main() {
	input := "nbysizxe"

	lens := 0
	grid := make([][]int, 128)
	for i := range grid {
		grid[i] = make([]int, 128)
	}
	for i := 0; i < 128; i++ {
		h := hex.EncodeToString(calculateKnotHash([]byte(fmt.Sprintf("%s-%d", input, i)), 256))
		line := ""
		for _, j := range h {
			ui, _ := strconv.ParseUint(string(j), 16, 64)
			line = fmt.Sprintf("%s%004b", line, ui)
		}
		lens += len(strings.Replace(line, "0", "", -1))
		for x, a := range line {
			if a == '1' {
				grid[i][x] = 1
			}
		}
	}
	fmt.Println(lens)
	next := 2
	for x, line := range grid {
		for y, field := range line {
			if field == 1 {
				flood(grid, x, y, next)
				next++
			}
		}
	}
	for _, line := range grid {
		for _, field := range line {
			fmt.Printf("%04d ", field)
		}
		fmt.Printf("\n")
	}
	fmt.Println(next)
}

func flood(grid [][]int, x, y, next int) {
	grid[x][y] = next
	if x > 0 {
		if grid[x-1][y] == 1 {
			flood(grid, x-1, y, next)
		}
	}
	if x < 127 {
		if grid[x+1][y] == 1 {
			flood(grid, x+1, y, next)
		}
	}
	if y > 0 {
		if grid[x][y-1] == 1 {
			flood(grid, x, y-1, next)
		}
	}
	if y < 127 {
		if grid[x][y+1] == 1 {
			flood(grid, x, y+1, next)
		}
	}
}

func calculateKnotHash(lengths []byte, len int) []byte {
	lengths = append(lengths, []byte{17, 31, 73, 47, 23}...)
	list := make([]byte, len)
	for i := range list {
		list[i] = byte(i)
	}

	start := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			iLen := int(length)
			twist(list, start, iLen, skip)
			start = (start + iLen + skip)
			skip++
		}
	}
	denseHash := make([]byte, len/16)
	for i := range denseHash {
		denseHash[i] = xor(list[i*16 : (i+1)*16])
	}
	return denseHash
}

func twist(input []byte, start, length, skip int) {
	for i := 0; i < length/2; i++ {
		s1 := (start + i) % len(input)
		s2 := (start + length - i - 1) % len(input)
		input[s1], input[s2] = input[s2], input[s1]
	}
}

func xor(data []byte) byte {
	out := byte(0)
	for _, num := range data {
		out = out ^ num
	}

	return out
}
