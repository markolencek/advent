package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Fields(string(buf))
	maze := make([]int, len(input))
	for i, val := range input {
		maze[i], err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	part1(maze)
	part2(maze)
}

func part2(input []int) {
	maze := make([]int, len(input))
	copy(maze, input)
	pos := 0
	last := len(maze) - 1
	steps := 0
	for {
		steps++
		move := maze[pos]
		if maze[pos] < 3 {
			maze[pos] = maze[pos] + 1
		} else {
			maze[pos] = maze[pos] - 1
		}
		pos = pos + move
		if pos < 0 || pos > last {
			break
		}
	}

	fmt.Println("Part 2 steps:", steps)
}

func part1(input []int) {
	maze := make([]int, len(input))
	copy(maze, input)
	pos := 0
	last := len(maze) - 1
	steps := 0
	for {
		steps++
		move := maze[pos]
		maze[pos] = maze[pos] + 1
		pos = pos + move
		if pos < 0 || pos > last {
			break
		}
	}

	fmt.Println("Part 1 steps:", steps)
}
