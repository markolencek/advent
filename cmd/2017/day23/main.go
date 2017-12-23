package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type inst struct {
	inst string
	val1 string
	val2 string
}

var sound int
var deadlockTs = 1 * time.Second

func main() {
	inst1 := readFile("input.txt")
	inst2 := readFile("input2.txt")
	part1 := map[string]int{}
	run(inst1, part1)
	fmt.Println(part1)
	part2 := map[string]int{"a": 1}
	run(inst2, part2)
	fmt.Println(part2)
}

func run(instructions []*inst, regs map[string]int) {
	mul := 0
	for i := 0; i < len(instructions) && i >= 0; {
		ins := instructions[i]
		switch ins.inst {
		case "snd":
			sound = getValue(ins.val1, regs)
		case "set":
			regs[ins.val1] = getValue(ins.val2, regs)
		case "add":
			regs[ins.val1] += getValue(ins.val2, regs)
		case "sub":
			regs[ins.val1] -= getValue(ins.val2, regs)
		case "mul":
			regs[ins.val1] = getValue(ins.val1, regs) * getValue(ins.val2, regs)
			mul++
		case "mod":
			regs[ins.val1] = getValue(ins.val1, regs) % getValue(ins.val2, regs)
		case "rcv":
			if getValue(ins.val1, regs) != 0 {
				regs[ins.val1] = sound
				if sound != 0 {
					fmt.Println("Part 1:", sound)
					return
				}
			}
		case "jgz":
			if getValue(ins.val1, regs) > 0 {
				i = i + getValue(ins.val2, regs)
				continue
			}
		case "jnz":
			if getValue(ins.val1, regs) != 0 {
				i = i + getValue(ins.val2, regs)
				continue
			}
		case "opt1":
			b := regs["b"]
			for d := regs["d"]; d != b; d++ {
				if b%d == 0 {
					regs["f"] = 0
				}
			}
		}
		i++
	}
	fmt.Println("Mul", mul)
}

func getValue(v string, regs map[string]int) int {
	if val, err := strconv.Atoi(v); err == nil {
		return val
	}
	return regs[v]
}

func readFile(fileName string) []*inst {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []*inst{}
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		ins := inst{
			inst: tmp[0],
			val1: tmp[1],
		}
		if len(tmp) > 2 {
			ins.val2 = tmp[2]
		}
		instructions = append(instructions, &ins)
	}

	return instructions
}
