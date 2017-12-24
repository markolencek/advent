package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	ModifyReg  string
	ByValue    int
	CheckReg   string
	CheckCond  string
	CheckVlaue int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	instructions := []instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		tmpIns := instruction{
			ModifyReg: tmp[0],
			CheckReg:  tmp[4],
			CheckCond: tmp[5],
		}
		tmpIns.ByValue, err = strconv.Atoi(tmp[2])
		if err != nil {
			panic(err)
		}
		if tmp[1] == "dec" {
			tmpIns.ByValue = -tmpIns.ByValue
		}
		tmpIns.CheckVlaue, err = strconv.Atoi(tmp[6])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, tmpIns)
	}

	register := map[string]int{}

	max := 0

	for _, inst := range instructions {
		tmp := inst.Exec(register)
		if tmp > max {
			max = tmp
		}
	}
	fmt.Println(max)

	max = 0
	t := true

	for _, val := range register {
		if t {
			max = val
			t = false
		}
		if val > max {
			max = val
		}
	}

	fmt.Println(max)
}

func (i *instruction) Exec(reg map[string]int) int {
	switch i.CheckCond {
	case ">":
		if reg[i.CheckReg] <= i.CheckVlaue {
			return reg[i.ModifyReg]
		}
	case "<":
		if reg[i.CheckReg] >= i.CheckVlaue {
			return reg[i.ModifyReg]
		}
	case ">=":
		if reg[i.CheckReg] < i.CheckVlaue {
			return reg[i.ModifyReg]
		}
	case "<=":
		if reg[i.CheckReg] > i.CheckVlaue {
			return reg[i.ModifyReg]
		}
	case "==":
		if reg[i.CheckReg] != i.CheckVlaue {
			return reg[i.ModifyReg]
		}
	case "!=":
		if reg[i.CheckReg] == i.CheckVlaue {
			return reg[i.ModifyReg]
		}
	}

	reg[i.ModifyReg] += i.ByValue
	return reg[i.ModifyReg]
}
