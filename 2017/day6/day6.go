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
	banks := make([]int, len(input))
	for i, val := range input {
		banks[i], err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
	}
	iter, size := part1(banks)
	fmt.Println("part1: ", iter)
	fmt.Println("part2: ", size)
}

func part1(banks []int) (int, int) {
	iter := 0
	configs := make(map[string]int)
	prev := 0
	for {
		max, index, config := maxAndConfig(banks)
		if configs[config] != 0 {
			prev = configs[config]
			break
		}
		configs[config] = iter
		iter++
		banks[index] = 0
		for i := max; i > 0; i-- {
			index++
			if index >= len(banks) {
				index = index % len(banks)
			}
			banks[index]++
		}
	}
	return iter, iter - prev
}

func maxAndConfig(ints []int) (int, int, string) {
	max := ints[0]
	index := 0
	config := make([]string, len(ints))
	for i, num := range ints {
		config[i] = strconv.Itoa(num)
		if num > max {
			max = num
			index = i
		}
	}
	return max, index, strings.Join(config, ":")
}
