package main

import (
	"fmt"
)

func main() {
	input := 371
	steps := input
	position := 0
	buffer := []int{0}

	for i := 1; i <= 2017; i++ {
		position = ((position + steps) % len(buffer))
		buffer = insert(buffer, position, i)
		position++
	}
	position = (position + 1) % len(buffer)
	fmt.Println("part 1:", buffer[position])

	size := 1
	position = 0
	pos1 := 0
	for i := 1; i <= 50000000; i++ {
		position = ((position + steps) % size)
		if position == 0 {
			pos1 = i
		}
		position++
		size++
	}
	fmt.Println("part 2:", pos1)
}

func insert(in []int, pos, val int) []int {
	out := make([]int, len(in)+1)
	copy(out[:pos+1], in[:pos+1])
	out[pos+1] = val
	if len(in) > pos+1 {
		copy(out[pos+2:], in[pos+1:])
	}
	return out
}
