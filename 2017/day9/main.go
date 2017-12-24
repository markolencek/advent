package main

import (
	"io/ioutil"
)

const (
	groupStart     = '{'
	groupEnd       = '}'
	garbageStart   = '<'
	garbageEnd     = '>'
	cancelNextChar = '!'
)

type group struct {
	groups []*group
	value  int
	sum    int
}

var garbageSum int

func main() {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := []rune(string(buf))
	gr, _ := parse(input, 0)
	println(gr.sum)
	println(garbageSum)
}

func parse(input []rune, val int) (*group, int) {
	gr := group{value: val + 1, sum: val + 1}
	for i := 1; i < len(input); i++ {
		switch input[i] {
		case groupStart:
			addGr, skip := parse(input[i:], gr.value)
			gr.groups = append(gr.groups, addGr)
			gr.sum += addGr.sum
			i += skip
		case groupEnd:
			return &gr, i
		case garbageStart:
		garbage:
			for i < len(input) {
				switch input[i] {
				case cancelNextChar:
					i++
				case garbageEnd:
					break garbage
				default:
					garbageSum++
				}
				i++
			}
		}
	}
	return nil, 0
}
