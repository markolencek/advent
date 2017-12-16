package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	start := []rune("abcdefghijklmnop")
	programs := []rune("abcdefghijklmnop")

	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	moves := strings.Split(string(buf), ",")

	for _, m := range moves {
		switch m[0] {
		case 'x':
			exchange(programs, m)
		case 'p':
			partner(programs, m)
		case 's':
			programs = spin(programs, m)
		}
	}

	fmt.Println("Part 1:", string(programs))

	copy(programs, start)

	repeats := 0
	for i := 1; ; i++ {
		for _, m := range moves {
			switch m[0] {
			case 'x':
				exchange(programs, m)
			case 'p':
				partner(programs, m)
			case 's':
				programs = spin(programs, m)
			}
		}
		if string(programs) == string(start) {
			fmt.Println("Repeats after: ", i)
			repeats = i
			break
		}
	}

	runs := 1000000000 % repeats
	copy(programs, start)

	for i := 0; i < runs; i++ {
		for _, m := range moves {
			switch m[0] {
			case 'x':
				exchange(programs, m)
			case 'p':
				partner(programs, m)
			case 's':
				programs = spin(programs, m)
			}
		}
	}

	fmt.Println("Part 2:", string(programs))
}

func spin(programs []rune, move string) []rune {
	steps, err := strconv.Atoi(strings.TrimPrefix(move, "s"))
	if err != nil {
		panic(err)
	}
	return append(programs[len(programs)-steps:], programs[:len(programs)-steps]...)
}

func exchange(programs []rune, move string) {
	s := strings.Split(strings.TrimPrefix(move, "x"), "/")
	if len(s) == 2 {
		i1, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		i2, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		programs[i1], programs[i2] = programs[i2], programs[i1]
	}
}

func partner(programs []rune, move string) {
	s := strings.Split(strings.TrimPrefix(move, "p"), "/")
	if len(s) == 2 {
		var i []int
		for x, p := range programs {
			if string(p) == s[0] || string(p) == s[1] {
				i = append(i, x)
			}
		}
		programs[i[0]], programs[i[1]] = programs[i[1]], programs[i[0]]
	}
}
