package main

import (
	"fmt"
)

type action struct {
	setValue  int
	moveSlots int
	nextState string
}

func main() {
	actions := initActions()
	steps := 12629077

	tape := map[int]int{}
	position := 0
	state := "A"

	for i := 0; i < steps; i++ {
		a := actions[state][tape[position]]
		tape[position] = a.setValue
		state = a.nextState
		position += a.moveSlots
	}
	checksum := 0
	for _, t := range tape {
		checksum += t
	}
	fmt.Println("Part 1:", checksum)
}

func initActions() map[string]map[int]action {
	actions := map[string]map[int]action{
		"A": map[int]action{
			0: action{
				setValue:  1,
				moveSlots: 1,
				nextState: "B",
			},
			1: action{
				setValue:  0,
				moveSlots: -1,
				nextState: "B",
			},
		},
		"B": map[int]action{
			0: action{
				setValue:  0,
				moveSlots: 1,
				nextState: "C",
			},
			1: action{
				setValue:  1,
				moveSlots: -1,
				nextState: "B",
			},
		},
		"C": map[int]action{
			0: action{
				setValue:  1,
				moveSlots: 1,
				nextState: "D",
			},
			1: action{
				setValue:  0,
				moveSlots: -1,
				nextState: "A",
			},
		},
		"D": map[int]action{
			0: action{
				setValue:  1,
				moveSlots: -1,
				nextState: "E",
			},
			1: action{
				setValue:  1,
				moveSlots: -1,
				nextState: "F",
			},
		},
		"E": map[int]action{
			0: action{
				setValue:  1,
				moveSlots: -1,
				nextState: "A",
			},
			1: action{
				setValue:  0,
				moveSlots: -1,
				nextState: "D",
			},
		},
		"F": map[int]action{
			0: action{
				setValue:  1,
				moveSlots: 1,
				nextState: "A",
			},
			1: action{
				setValue:  1,
				moveSlots: -1,
				nextState: "E",
			},
		},
	}
	return actions
}
