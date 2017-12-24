package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tower struct {
	Name     string
	Weight   int
	ChStr    []string
	Children []*tower
	Parent   *tower
	sum      int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		lines = append(lines, tmp)
	}

	towers := parse(lines)
	for _, t := range towers {
		if t.Parent == nil {
			fmt.Println("Bottom tower: ", t.Name)
			t.Check(0)
			break
		}
	}
}

func parse(lines [][]string) map[string]*tower {
	towers := map[string]*tower{}
	for _, ln := range lines {
		w, err := strconv.Atoi(strings.Trim(ln[1], "()"))
		if err != nil {
			panic(err)
		}
		t := tower{
			Name:   ln[0],
			Weight: w,
		}
		if len(ln) > 3 {
			for _, c := range ln[3:] {
				t.ChStr = append(t.ChStr, strings.Trim(c, ","))
			}
		}
		towers[t.Name] = &t
	}
	for n, t := range towers {
		if len(t.ChStr) > 0 {
			for _, c := range t.ChStr {
				towers[c].Parent = towers[n]
				t.Children = append(t.Children, towers[c])
			}
		}
	}
	return towers
}

func (t *tower) sumWeight() int {
	if t.sum != 0 {
		return t.sum
	}
	sum := 0
	for _, c := range t.Children {
		sum += c.sumWeight()
	}
	t.sum = t.Weight + sum
	return t.sum
}

func (t *tower) Check(diff int) {
	w, w2, wc, w2c := 0, 0, 0, 0
	for _, c := range t.Children {
		if w == 0 {
			w = c.sumWeight()
			wc++
			continue
		}
		if w != c.sumWeight() {
			w2 = c.sumWeight()
			w2c++
			continue
		}
		wc++
	}
	if w2c > 0 {
		ok := w2
		if w2c == 1 {
			ok = w
		}
		for _, c := range t.Children {
			if ok != c.sumWeight() {
				c.Check(c.sumWeight() - ok)
				break
			}
		}
	} else if diff != 0 {
		fmt.Println("Wrong: ", t.Name, " Weight: ", t.Weight, " Should be: ", t.Weight-diff)
	}
}
