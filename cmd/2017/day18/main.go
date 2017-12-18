package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type inst struct {
	inst string
	val1 string
	val2 string
}

var instructions = []*inst{}
var sound int
var deadlockTs = 1 * time.Second

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
	part1()

	chansize := 1000
	prg1 := make(chan int, chansize)
	prg2 := make(chan int, chansize)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go part2(prg2, prg1, 0, &wg)
	go part2(prg1, prg2, 1, &wg)
	wg.Wait()
	fmt.Println("Done!")
}

func part2(in, out chan (int), id int, wg *sync.WaitGroup) {
	sent, rcv := 0, 0
	defer wg.Done()
	var regs = map[string]int{}
	regs["p"] = id
	for i := 0; i < len(instructions) && i >= 0; {
		ins := instructions[i]
		switch ins.inst {
		case "snd":
			out <- getValue(ins.val1, regs)
			sent++
		case "set":
			regs[ins.val1] = getValue(ins.val2, regs)
		case "add":
			regs[ins.val1] += getValue(ins.val2, regs)
		case "mul":
			regs[ins.val1] = getValue(ins.val1, regs) * getValue(ins.val2, regs)
		case "mod":
			regs[ins.val1] = getValue(ins.val1, regs) % getValue(ins.val2, regs)
		case "rcv":
			timeout := timer(deadlockTs)
			select {
			case regs[ins.val1] = <-in:
				rcv++
			case <-timeout:
				fmt.Println(id, "sent:", sent)
				fmt.Println(id, "rcv:", rcv)
				return
			}
		case "jgz":
			if getValue(ins.val1, regs) > 0 {
				i = i + getValue(ins.val2, regs)
				continue
			}
		}
		i++
	}
	fmt.Println(id, "sent:", sent)
	fmt.Println(id, "rcv:", rcv)
}

func timer(wt time.Duration) chan (bool) {
	timeout := make(chan bool)
	go func() {
		time.Sleep(wt)
		timeout <- true
	}()
	return timeout
}

func part1() {
	var regs = map[string]int{}
	for i := 0; i < len(instructions) && i >= 0; {
		ins := instructions[i]
		switch ins.inst {
		case "snd":
			sound = getValue(ins.val1, regs)
		case "set":
			regs[ins.val1] = getValue(ins.val2, regs)
		case "add":
			regs[ins.val1] += getValue(ins.val2, regs)
		case "mul":
			regs[ins.val1] = getValue(ins.val1, regs) * getValue(ins.val2, regs)
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
		}
		i++
	}
}

func getValue(v string, regs map[string]int) int {
	if val, err := strconv.Atoi(v); err == nil {
		return val
	}
	return regs[v]
}
