package main

import (
	"encoding/hex"
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
	part1(string(buf))
	part2(buf)
}

func part2(input []byte) {
	lengths := append(input, []byte{17, 31, 73, 47, 23}...)

	list := make([]byte, 256)
	for i := range list {
		list[i] = byte(i)
	}

	start := 0
	skip := 0
	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			iLen := int(length)
			twist(list, start, iLen, skip)
			start = (start + iLen + skip)
			skip++
		}
	}
	denseHash := make([]byte, 16)
	for i := range denseHash {
		denseHash[i] = xor(list[i*16 : (i+1)*16])
	}
	fmt.Println("part2 :", hex.EncodeToString(denseHash))
}

func xor(data []byte) byte {
	out := byte(0)
	for _, num := range data {
		out = out ^ num
	}

	return out
}

func part1(input string) {
	lengths := strings.Split(input, ",")

	list := make([]byte, 256)
	for i := range list {
		list[i] = byte(i)
	}

	start := 0
	for skip, length := range lengths {
		iLen, err := strconv.Atoi(length)
		if err != nil {
			panic(err)
		}
		twist(list, start, iLen, skip)
		start = (start + iLen + skip)
	}
	fmt.Println("part1 :", int(list[0])*int(list[1]))
}

func twist(input []byte, start, length, skip int) {
	for i := 0; i < length/2; i++ {
		s1 := (start + i) % len(input)
		s2 := (start + length - i - 1) % len(input)
		input[s1], input[s2] = input[s2], input[s1]
	}
}
