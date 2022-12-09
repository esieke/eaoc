package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	n := 2
	pos := make([]position, n)
	track := make(map[position]int)
	track[pos[0]] = 1 // initial position
	for s.Scan() {
		l := s.Text()
		ls := strings.Split(l, " ")
		if len(ls) != 2 {
			panic("parser error")
		}
		steps, err := strconv.Atoi(ls[1])
		if err != nil {
			panic("must be an int")
		}

		for i := 0; i < steps; i++ {
			if ls[0] == "R" {
				pos[0].x += 1
			}
			if ls[0] == "L" {
				pos[0].x -= 1
			}
			if ls[0] == "U" {
				pos[0].y += 1
			}
			if ls[0] == "D" {
				pos[0].y -= 1
			}
			for i, _ := range pos {
				if i < 1 {
					continue
				}
				dx := pos[i-1].x - pos[i].x
				dy := pos[i-1].y - pos[i].y
				if dx > 1 || dx < -1 || dy > 1 || dy < -1 {
					pos[i].x += rangeLimit(dx)
					pos[i].y += rangeLimit(dy)
					if i == len(pos)-1 {
						track[pos[i]] = 1
					}
				}
			}
		}
	}
	var r int
	for _, v := range track {
		r += v
	}
	fmt.Println(r)
}

func rangeLimit(v int) int {
	if v >= 1 {
		return 1
	}
	if v <= -1 {
		return -1
	}
	return 0
}
