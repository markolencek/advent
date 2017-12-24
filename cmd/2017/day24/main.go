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

type comp struct {
	ID    int
	port1 int
	port2 int
	str   int
	used  bool
}

var sound int
var deadlockTs = 1 * time.Second

func main() {
	comps := readFile("input.txt")
	maxStr, maxLen, maxMax := build(0, 0, 0, comps)
	fmt.Println(maxMax)
	fmt.Println(maxStr, maxLen)
}

func build(port, str, len int, comps []*comp) (int, int, int) {
	maxStr := 0
	maxLen := 0
	maxMax := 0

	maxes := func(nStr, nLen, nMax int) {
		maxMax = max(maxMax, nMax)
		if nLen > maxLen {
			maxLen = nLen
			maxStr = nStr
		} else if nLen == maxLen {
			maxStr = max(maxStr, nStr)
		}
	}

	for _, c := range comps {
		if c.used {
			continue
		}
		if port == c.port1 {
			c.used = true
			maxes(build(c.port2, c.str, len+1, comps))
			c.used = false
		} else if port == c.port2 {
			c.used = true
			maxes(build(c.port1, c.str, len+1, comps))
			c.used = false
		}
	}
	return maxStr + str, maxLen + len, maxMax + str
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func readFile(fileName string) []*comp {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	comps := []*comp{}

	scanner := bufio.NewScanner(file)
	for id := 0; scanner.Scan(); id++ {
		tmp := strings.Split(strings.TrimSpace(scanner.Text()), "/")
		p1, err := strconv.Atoi(tmp[0])
		if err != nil {
			log.Fatal(err)
		}
		p2, err := strconv.Atoi(tmp[1])
		if err != nil {
			log.Fatal(err)
		}
		c := comp{
			ID:    id,
			port1: p1,
			port2: p2,
			str:   p1 + p2,
		}
		comps = append(comps, &c)
	}

	return comps
}
