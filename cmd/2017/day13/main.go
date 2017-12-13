package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	layers := map[int]int{}
	severity := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ":")
		l, err := strconv.Atoi(strings.Trim(tmp[0], " \t"))
		if err != nil {
			log.Fatal(err)
		}
		d, err := strconv.Atoi(strings.Trim(tmp[1], " \t"))
		if err != nil {
			log.Fatal(err)
		}
		layers[l] = d
		if (l % (d + d - 2)) == 0 {
			severity += l * d
		}
	}

	fmt.Println("Part 1:", severity)

wait:
	for delay := 0; delay < math.MaxInt32; delay++ {
		for l, d := range layers {
			if ((l + delay) % (d + d - 2)) == 0 {
				continue wait
			}
		}
		fmt.Println("Part 2:", delay)
		break
	}
}
