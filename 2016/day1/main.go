package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := "R8, R4, R4, R8"
	input = "L5, R1, L5, L1, R5, R1, R1, L4, L1, L3, R2, R4, L4, L1, L1, R2, R4, R3, L1, R4, L4, L5, L4, R4, L5, R1, R5, L2, R1, R3, L2, L4, L4, R1, L192, R5, R1, R4, L5, L4, R5, L1, L1, R48, R5, R5, L2, R4, R4, R1, R3, L1, L4, L5, R1, L4, L2, L5, R5, L2, R74, R4, L1, R188, R5, L4, L2, R5, R2, L4, R4, R3, R3, R2, R1, L3, L2, L5, L5, L2, L1, R1, R5, R4, L3, R5, L1, L3, R4, L1, L3, L2, R1, R3, R2, R5, L3, L1, L1, R5, L4, L5, R5, R2, L5, R2, L1, L5, L3, L5, L5, L1, R1, L4, L3, L1, R2, R5, L1, L3, R4, R5, L4, L1, R5, L1, R5, R5, R5, R2, R1, R2, L5, L5, L5, R4, L5, L4, L4, R5, L2, R1, R5, L1, L5, R4, L3, R4, L2, R3, R3, R3, L2, L2, L2, L1, L4, R3, L4, L2, R2, R5, L1, R2"

	fmt.Println(calc(input))
}

func calc(input string) float64 {
	directions := strings.Split(input, ",")

	orientation := 0
	x := 0
	y := 0

	locations := make(map[string]bool)
mainloop:
	for _, dir := range directions {
		runes := []rune(strings.Trim(dir, " "))
		if runes[0] == 'R' {
			orientation = (orientation + 1) % 4
		} else {
			orientation = (orientation - 1) % 4
		}
		if orientation < 0 {
			orientation += 4
		}
		distance, _ := strconv.Atoi(string(runes[1:]))

		for i := 1; i <= distance; i++ {
			switch orientation {
			case 0:
				x++
			case 1:
				y++
			case 2:
				x--
			case 3:
				y--
			}
			if locations[fmt.Sprint(x, ",", y)] == true {
				break mainloop
			}
			locations[fmt.Sprint(x, ",", y)] = true
		}
	}

	return math.Abs(float64(x)) + math.Abs(float64(y))
}
