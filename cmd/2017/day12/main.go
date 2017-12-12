package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type program struct {
	visited   bool
	inGroup   int
	connected []*program
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		lines = append(lines, tmp)
	}
	part1(lines)
}

func part1(input [][]string) {
	programs := make([]program, len(input))
	for i, line := range input {
		if len(line) < 3 {
			continue
		}
		for _, conn := range line[2:] {
			j, err := strconv.Atoi(strings.Trim(conn, ","))
			if err != nil {
				log.Fatal(err)
			}
			programs[i].connected = append(programs[i].connected, &programs[j])
		}
	}

	groups := []int{}
	for i := 0; i < len(programs); i++ {
		if programs[i].inGroup == 0 {
			count := walk(&programs[i], len(groups)+1)
			groups = append(groups, count)
		}
	}
	fmt.Println("Group 0:", groups[0])
	fmt.Println("All groups:", len(groups))
}

func walk(prg *program, group int) int {
	if prg.visited {
		return 0
	}
	prg.visited = true
	count := 1
	prg.inGroup = group
	for _, p := range prg.connected {
		if p == prg {
			continue
		}
		count += walk(p, group)
	}
	return count
}
