package main

import (
	"fmt"
	"math"
)

type program struct {
	visited   bool
	inGroup   int
	connected []*program
}

const mA = 16807
const mB = 48271

const mod = math.MaxInt32

func main() {
	low := int(math.Pow(2, 16))
	startA := 873
	startB := 583
	inputA := startA
	inputB := startB

	match := 0

	for i := 0; i < 40000000; i++ {
		inputA = (inputA * mA) % mod
		inputB = (inputB * mB) % mod
		if inputA%low == inputB%low {
			match++
		}
	}

	fmt.Println("Part 1:", match)

	match = 0
	cA := make(chan int)
	cB := make(chan int)

	go calc(startB, mB, 8, cB)
	go calc(startA, mA, 4, cA)

	for i := 0; i < 5000000; i++ {
		inputA, inputB = <-cA, <-cB

		if inputA%low == inputB%low {
			match++
		}
	}
	fmt.Println("Part 2:", match)
}

func calc(input, m, cond int, c chan int) {
	for {
		input = (input * m) % mod
		if input%cond == 0 {
			c <- input
		}
	}
}
