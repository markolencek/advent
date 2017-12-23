package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules := readFile("input.txt")
	pattern := [][]rune{
		[]rune(".#."),
		[]rune("..#"),
		[]rune("###"),
	}

	for i := 0; i < 5; i++ {
		pattern = expand(pattern, rules)
	}
	{
		joined := ""
		for _, p := range pattern {
			joined = joined + string(p)
		}
		count := strings.Replace(joined, ".", "", -1)
		fmt.Println(len(count))
	}
	for i := 0; i < 13; i++ {
		pattern = expand(pattern, rules)
	}
	{
		joined := ""
		for _, p := range pattern {
			joined = joined + string(p)
		}
		count := strings.Replace(joined, ".", "", -1)
		fmt.Println(len(count))
	}
}

func expand(pattern [][]rune, rules map[string][]string) [][]rune {
	n := 3
	m := 4
	if len(pattern)%2 == 0 {
		n = 2
		m = 3
	}
	new := make([][]rune, (len(pattern)/n)*m)
	for i := 0; i < len(pattern)/n; i++ {
		for j := 0; j < len(pattern)/n; j++ {
			sample := ""
			for k := 0; k < n; k++ {
				sample = sample + string(pattern[(i*n)+k][(j*n):(j+1)*n])
			}
			out := rules[sample]
			for k := 0; k < m; k++ {
				new[i*m+k] = append(new[i*m+k], []rune(out[k])...)
			}
		}
	}
	return new
}

func getValue(v string, regs map[string]int) int {
	if val, err := strconv.Atoi(v); err == nil {
		return val
	}
	return regs[v]
}

func readFile(fileName string) map[string][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := map[string][]string{}
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		makeRules(tmp[0], tmp[2], rules)
	}

	return rules
}

func makeRules(input, output string, outMap map[string][]string) {
	p := strings.Split(input, "/")
	o := strings.Split(output, "/")
	patterns := []string{}
	for i := 0; i < 2; i++ {
		patterns = append(patterns, strings.Join(p, ""))
		mirrorH(p)
		patterns = append(patterns, strings.Join(p, ""))
		mirrorV(p)
		patterns = append(patterns, strings.Join(p, ""))
		mirrorH(p)
		patterns = append(patterns, strings.Join(p, ""))
		transpose(p)
	}
	for _, pat := range patterns {
		outMap[pat] = o
	}
}

func transpose(a []string) {
	n := len(a)
	b := [][]rune{}
	for _, aa := range a {
		b = append(b, []rune(aa))
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			b[i][j], b[j][i] = b[j][i], b[i][j]
		}
	}
	for i, bb := range b {
		a[i] = string(bb)
	}
}

func mirrorH(a []string) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func mirrorV(a []string) {
	n := len(a)
	for i, pp := range a {
		r := []rune(pp)
		for i := 0; i < n/2; i++ {
			r[i], r[n-i-1] = r[n-i-1], r[i]
		}
		a[i] = string(r)
	}
}
