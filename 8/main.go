package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	m := [][]byte{}
	v := [][]byte{}
	for s.Scan() {
		l := s.Text()
		lb := []byte(l)
		for i, _ := range lb {
			lb[i] = lb[i] - '0'
		}
		vb := make([]byte, len(lb))
		m = append(m, lb)
		v = append(v, vb)
	}

	// check rows
	for y, _ := range m {
		var x_max_pos byte // right to left
		var x_max_neg byte // left to right
		for x, _ := range m[y] {
			if m[y][x] > x_max_pos || x == 0 {
				x_max_pos = m[y][x]
				v[y][x] = 1
			}
			if m[y][len(m[y])-1-x] > x_max_neg || x == 0 {
				x_max_neg = m[y][len(m[y])-1-x]
				v[y][len(m[y])-1-x] = 1
			}
		}
	}

	// check colums
	for x, _ := range m[0] {
		var y_max_pos byte // up to down
		var y_max_neg byte // down to up
		for y, _ := range m {
			if m[y][x] > y_max_pos || y == 0 {
				y_max_pos = m[y][x]
				v[y][x] = 1
			}
			if m[len(m)-1-y][x] > y_max_neg || y == 0 {
				y_max_neg = m[len(m)-1-y][x]
				v[len(m)-1-y][x] = 1
			}
		}
	}

	var r uint64
	for y, _ := range m {
		for x, _ := range m[y] {
			r += uint64(v[y][x])
		}
	}

	// fmt.Println(m)
	// fmt.Println(v)
	fmt.Println(r)
}
