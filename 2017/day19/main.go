package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	moveY = '-'
	moveX = '|'
	moveC = '+'
	moveS = ' '
)

func main() {
	input := readFile("input.txt")

	startY := -1
	for x, c := range input[0] {
		if c == moveX {
			startY = x
			break
		}
	}
	if startY < 0 {
		panic("eek")
	}

	x, y, dirX, dirY := 0, startY, 1, 0
	letters := []rune{}
	maxX := len(input) - 1
	maxY := len(input[0]) - 1
	steps := 0
	for {
		if input[x][y] == moveS {
			break
		}
		steps++
		switch input[x][y] {
		case moveC:
		case moveS:
		case moveX:
		case moveY:
		default:
			letters = append(letters, input[x][y])
		}
		x += dirX
		y += dirY
		if input[x][y] == moveC {
			if dirX != 1 && x > 0 && input[x-1][y] != moveS {
				dirX, dirY = -1, 0
			} else if dirY != 1 && y > 0 && input[x][y-1] != moveS {
				dirX, dirY = 0, -1
			} else if dirX != -1 && x < maxX-1 && input[x+1][y] != moveS {
				dirX, dirY = 1, 0
			} else if dirY != -1 && y < maxY-1 && input[x][y+1] != moveS {
				dirX, dirY = 0, 1
			}
		}
	}
	fmt.Println("Part 1:", string(letters))
	fmt.Println("Part 2:", steps)
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
