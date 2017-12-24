package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	cleanNode    = '.'
	infectedNode = '#'
	weakenedNode = 'W'
	flaggedNode  = 'F'
)

type checker interface {
	check(rune, int) (rune, int)
}

type checker1 struct{}
type checker2 struct{}

func main() {
	input := readFile("input.txt")
	doPart1(input)
	doPart2(input)
}

func doPart1(input [][]rune) {
	mid := len(input) / 2
	grid := initMap(input)
	x, y, infs := walk(mid, mid, 10000, grid, checker1{})
	fmt.Println(x, y, infs)
}
func doPart2(input [][]rune) {
	mid := len(input) / 2
	grid := initMap(input)
	x, y, infs := walk(mid, mid, 10000000, grid, checker2{})
	fmt.Println(x, y, infs)
}

func readFile(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := [][]rune{}
	for scanner.Scan() {
		l := []rune(scanner.Text())
		c = append(c, l)
	}

	return c
}

func initMap(input [][]rune) map[string]rune {
	grid := map[string]rune{}
	for x, a := range input {
		for y, b := range a {
			grid[fmt.Sprintf("%d.%d", x, y)] = b
		}
	}
	return grid
}

func walk(x, y, steps int, grid map[string]rune, chk checker) (int, int, int) {
	dir := 0
	infections := 0
	for i := 0; i < steps; i++ {
		coord := fmt.Sprintf("%d.%d", x, y)
		grid[coord], dir = chk.check(grid[coord], dir)
		if grid[coord] == infectedNode {
			infections++
		}
		switch dir {
		case 0:
			x--
		case 1:
			y++
		case 2:
			x++
		case 3:
			y--
		}
	}
	return x, y, infections
}

func (p checker1) check(node rune, dir int) (rune, int) {
	switch node {
	case infectedNode:
		dir = (dir + 1) % 4
		node = cleanNode
	default:
		dir = (dir + 3) % 4
		node = infectedNode
	}
	return node, dir
}

func (p checker2) check(node rune, dir int) (rune, int) {
	switch node {
	case weakenedNode:
		node = infectedNode
	case infectedNode:
		node = flaggedNode
		dir = (dir + 1) % 4
	case flaggedNode:
		node = cleanNode
		dir = (dir + 2) % 4
	default:
		node = weakenedNode
		dir = (dir + 3) % 4
	}
	return node, dir
}
