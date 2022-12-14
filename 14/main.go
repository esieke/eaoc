package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type Map struct {
	min, max Pos
	t        [][]byte // tiles
}

func (m *Map) InitMap(max Pos) {
	m.min.x = max.x
	m.t = make([][]byte, max.y)
	for i, _ := range m.t {
		t := make([]byte, max.x)
		m.t[i] = t
	}
	m.t[0][500] = '+'
}
func (m *Map) DrawLine(start, end Pos) {
	// vert
	y0, y1 := start.y, end.y
	if start.y > end.y {
		y0, y1 = end.y, start.y
	}
	if start.x == end.x {
		x := start.x
		for y := y0; y <= y1; y++ {
			m.t[y][x] = '#'
		}
		return
	}
	// hor
	x0, x1 := start.x, end.x
	if start.x > end.x {
		x0, x1 = end.x, start.x
	}
	y := start.y
	for x := x0; x <= x1; x++ {
		m.t[y][x] = '#'
	}
	m.MinMax(start)
	m.MinMax(end)
	return
}

func (m *Map) MinMax(pos Pos) {
	if pos.x > m.max.x {
		m.max.x = pos.x + 1
		if m.max.x+1 > len(m.t[0]) {
			panic("x max out of range")
		}
	}
	if pos.y > m.max.y {
		m.max.y = pos.y
		if m.max.y+1 > len(m.t) {
			panic("y max out of range")
		}
	}
	if pos.x < m.min.x {
		m.min.x = pos.x - 1
		if m.min.x < 0 {
			panic("x min out of range")
		}
	}
	if pos.y < m.min.y {
		m.min.y = pos.y
		if m.min.y < 0 {
			panic("y min out of range")
		}
	}
}

func (m *Map) Print() {
	for y := m.min.y; y <= m.max.y; y++ {
		for x := m.min.x; x <= m.max.x; x++ {
			if m.t[y][x] == 0 {
				fmt.Printf(".")
				continue
			}
			fmt.Printf("%s", string(m.t[y][x]))
		}
		fmt.Printf("\n")
	}
}

func (m *Map) SearchAbyss() {
	for y := m.min.y; y <= m.max.y; y++ {
		for x := m.min.x; x <= m.max.x; x++ {
			if m.t[y][x] == 0 && m.t[y+1][x] == '#' {
				// left
				if m.IsAbyss(Pos{y: y + 1, x: x - 1}) {
					m.t[y+1][x-1] = 'a'
				}
				// right
				if m.IsAbyss(Pos{y: y + 1, x: x + 1}) {
					m.t[y+1][x+1] = 'a'
				}
			}
		}
	}
}

func (m *Map) IsAbyss(pos Pos) bool {
	x := pos.x
	for y := pos.y; y <= m.max.y; y++ {
		if m.t[y][x] != 0 {
			return false
		}
	}
	return true
}

func (m *Map) Simulate(pos Pos) int {
	ret := 0
	for true {
		p := pos
		ret++
		for true {
			var s int
			p, s = m.Step(p)
			if s == 3 {
				return ret
			}
			if s == 0 {
				break
			}
		}
		//fmt.Println("step ", ret-1)
		//m.Print()
	}
	return ret
}

func (m *Map) Step(pos Pos) (Pos, int) {
	if m.t[pos.y+1][pos.x] == 0 {
		pos.y += 1
		return pos, 1 // moving
	}
	if m.t[pos.y+1][pos.x-1] == 'a' {
		return pos, 3 // abyss
	}
	if m.t[pos.y+1][pos.x-1] == 0 {
		pos.y += 1
		pos.x -= 1
		return pos, 1 // moving
	}
	if m.t[pos.y+1][pos.x+1] == 'a' {
		return pos, 3 // abyss
	}
	if m.t[pos.y+1][pos.x+1] == 0 {
		pos.y += 1
		pos.x += 1
		return pos, 1 // moving
	}
	m.t[pos.y][pos.x] = 'O'
	return pos, 0 // stable position
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	m := Map{}
	m.InitMap(Pos{x: 1000, y: 1000})
	for s.Scan() {
		l := s.Text()
		pathS := strings.Split(l, " -> ")
		path := []Pos{}
		for i, v := range pathS {
			posS := strings.Split(v, ",")
			x, _ := strconv.Atoi(posS[0])
			y, _ := strconv.Atoi(posS[1])
			path = append(path, Pos{x: x, y: y})
			if i > 0 {
				m.DrawLine(path[i-1], path[i])
			}
		}
	}
	m.SearchAbyss()
	//m.Print()
	r := m.Simulate(Pos{y: 0, x: 500})
	fmt.Println(" ---------------- ")
	m.Print()
	fmt.Println(r - 1)
}
