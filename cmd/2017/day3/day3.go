package main

import (
	"fmt"
	"math"
)

func main() {
	input := 368078
	_ = input

	max := 0
	lastLoop := []int{1}
	lastSize := 1

	for max < input {
		lastSize += 2
		lastLoop = calcLoop(lastSize, lastLoop)
		max = lastLoop[len(lastLoop)-1]
		fmt.Println(max)
	}

	for _, num := range lastLoop {
		if num > input {
			fmt.Println(num)
			return
		}
	}
}

func calc(input float64) int {
	sq := math.Floor(math.Sqrt(input - 1))
	size := sq + math.Mod(sq, 2) + 1
	psize := size - 2
	start := psize*psize + 1
	pos := math.Mod(input-start, size-1)
	dist := int(size) - 2

	for i := 1.0; i <= pos; i++ {
		if i < (size-1.0)/2.0 {
			dist--
		} else {
			dist++
		}
	}

	return dist
}

func calcLoop(size int, previous []int) []int {
	prev := append([]int{previous[len(previous)-1]}, previous...)
	prevSize := size - 2
	loop := make([]int, size*4-4)
	partSize := len(loop) / 4

	for i := 0; i < 4; i++ {
		var tmpPrev []int
		if i == 0 {
			tmpPrev = append([]int{0}, prev[i*(prevSize-1):(i+1)*(prevSize-1)+1]...)
		} else {
			tmpPrev = append([]int{loop[i*partSize-2]}, prev[i*(prevSize-1):(i+1)*(prevSize-1)+1]...)
		}
		if i == 3 {
			tmpPrev = append(tmpPrev, loop[0], 0)
		} else {
			tmpPrev = append(tmpPrev, 0, 0)
		}

		for j := 0; j < partSize; j++ {
			loop[i*partSize+j] = tmpPrev[j] + tmpPrev[j+1] + tmpPrev[j+2]
			if (i*partSize + j) > 0 {
				loop[i*partSize+j] += loop[i*partSize+j-1]
			}
		}
	}

	return loop
}
