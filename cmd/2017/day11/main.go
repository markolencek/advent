package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	moves := strings.Split(string(buf), ",")
	part1(moves)
}

func part1(input []string) {
	x, y, z, dist, maxDist := 0.0, 0.0, 0.0, 0.0, 0.0
	for _, step := range input {
		switch step {
		case "n":
			y++
			z--
		case "s":
			y--
			z++
		case "ne":
			x++
			z--
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		case "se":
			x++
			y--
		}
		dist = (math.Abs(x) + math.Abs(y) + math.Abs(z)) / 2
		if dist > maxDist {
			maxDist = dist
		}
	}
	fmt.Println("part1:", dist)
	fmt.Println("part2:", maxDist)
}
