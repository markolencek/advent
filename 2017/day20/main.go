package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type v3 struct {
	x, y, z int
}

func (v v3) dist() int {
	return int(math.Abs(float64(v.x)) + math.Abs(float64(v.y)) + math.Abs(float64(v.z)))
}

type particle struct {
	p, v, a v3
	active  bool
}

func (p *particle) move() {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z
	p.p.x += p.v.x
	p.p.y += p.v.y
	p.p.z += p.v.z
}

func (p *particle) dA() int {
	return p.a.dist()
}
func (p *particle) dV() int {
	return p.v.dist()
}
func (p *particle) dP() int {
	return p.p.dist()
}

func main() {
	ps := readFile("input.txt")

	minA := math.MaxInt64
	for _, p := range ps {
		tmp := p.dA()
		if tmp < minA {
			minA = tmp
		}
	}

	minAI := []int{}
	for i, p := range ps {
		if p.dA() == minA {
			minAI = append(minAI, i)
		}
	}
	fmt.Println(minAI)

	for i := 0; ; i++ {
		for _, p := range ps {
			if p.active {
				p.move()
			}
		}
		for _, p1 := range ps {
			if !p1.active {
				continue
			}
			for _, p2 := range ps {
				if !p2.active {
					continue
				}
				if p1 == p2 {
					continue
				}
				if p1.p.x == p2.p.x && p1.p.y == p2.p.y && p1.p.z == p2.p.z {
					p1.active = false
					p2.active = false
				}
			}
		}
		actives := 0
		for _, p := range ps {
			if p.active {
				actives++
			}
		}
		fmt.Println(i, actives)
	}

}

func readFile(filename string) []*particle {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ps := []*particle{}
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " ")
		p := particle{active: true}
		for _, cc := range c {
			t := strings.Split(cc, "<")
			t = strings.Split(t[1], ">")
			v := strings.Split(t[0], ",")
			p1, err := strconv.Atoi(v[0])
			if err != nil {
				log.Fatal(err)
			}
			p2, err := strconv.Atoi(v[1])
			if err != nil {
				log.Fatal(err)
			}
			p3, err := strconv.Atoi(v[2])
			if err != nil {
				log.Fatal(err)
			}
			switch cc[0] {
			case 'p':
				p.p = v3{x: p1, y: p2, z: p3}
			case 'v':
				p.v = v3{x: p1, y: p2, z: p3}
			case 'a':
				p.a = v3{x: p1, y: p2, z: p3}
			default:
				log.Fatal("parsing")
			}
		}
		ps = append(ps, &p)
	}

	return ps
}
