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

	var r int
	var cycle int
	var X int = 1
	for s.Scan() {
		l := s.Text()
		ls := strings.Split(l, " ")

		// istruction parser
		if len(ls) == 1 {
			// noop
			cycle += 1
			r += signalStrength(cycle, X)
		}
		if len(ls) == 2 {
			// addx
			v, e := strconv.Atoi(ls[1])
			if e != nil {
				panic("syntax error. value must be an integer")
			}
			cycle += 1
			r += signalStrength(cycle, X)
			cycle += 1
			r += signalStrength(cycle, X)
			X += v
		}
		if len(ls) > 2 || len(ls) < 1 {
			panic("syntax error")
		}
	}
	fmt.Println(r)
}

func signalStrength(cycle, X int) int {
	var r int
	if cycle%40 == 20 {
		r = cycle * X
	}
	return r
}
