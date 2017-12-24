package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	phrases := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := strings.Fields(scanner.Text())
		phrases = append(phrases, tmp)
	}

	a := "gasff"
	b := strings.Split(a, "")
	sort.Strings(b)
	fmt.Println(b, strings.Join(b, ""))

	part1(phrases)
	part2(phrases)
}

func part2(phrases [][]string) {
	valid := 0
	duplicates := 0

	for _, phrase := range phrases {
		counted := map[string]int{}
		duplicate := false
		for _, str := range phrase {
			s := strings.Split(str, "")
			sort.Strings(s)
			str = strings.Join(s, "")
			if _, ok := counted[str]; ok {
				duplicate = true
				break
			}
			counted[str] = 1
		}
		if duplicate {
			duplicates++
			continue
		}
		valid++
	}
	fmt.Println("valid: ", valid)
	fmt.Println("duplicate: ", duplicates)
}

func part1(phrases [][]string) {
	valid := 0
	duplicates := 0

	for _, phrase := range phrases {
		counted := map[string]int{}
		duplicate := false
		for _, str := range phrase {
			if _, ok := counted[str]; ok {
				duplicate = true
				break
			}
			counted[str] = 1
		}
		if duplicate {
			duplicates++
			continue
		}
		valid++
	}
	fmt.Println("valid: ", valid)
	fmt.Println("duplicate: ", duplicates)
}
