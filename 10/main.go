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

	var cycle int
	var X int = 1
	sprite := make([]byte, 40)
	screen := make([][]byte, 6)
	spriteInit(sprite)
	screenInit(screen)
	for s.Scan() {
		l := s.Text()
		ls := strings.Split(l, " ")

		// istruction parser
		if len(ls) == 1 {
			// noop
			cycle += 1
			screenUpdate(screen, sprite, cycle)
		}
		if len(ls) == 2 {
			// addx
			v, e := strconv.Atoi(ls[1])
			if e != nil {
				panic("syntax error. value must be an integer")
			}
			cycle += 1
			screenUpdate(screen, sprite, cycle)
			cycle += 1
			screenUpdate(screen, sprite, cycle)
			X += v
			spriteUpdate(sprite, X, cycle)
		}
		if len(ls) > 2 || len(ls) < 1 {
			panic("syntax error")
		}
	}
}

func screenInit(screen [][]byte) {
	for y, _ := range screen {
		screen[y] = make([]byte, 40)
		for x, _ := range screen[y] {
			screen[y][x] = 'x'
		}
	}
}

func screenUpdate(screen [][]byte, sprite []byte, cycle int) {
	x := (cycle - 1) % 40
	y := (cycle - 1) % 240 / 40
	screen[y][x] = sprite[x]
	if x == 39 && y == 5 {
		for py, _ := range screen {
			for px, _ := range screen[py] {
				fmt.Printf("%s", string(screen[py][px]))
			}
			fmt.Printf("\n")
		}
	}
}

func spriteInit(sprite []byte) {
	for i, _ := range sprite {
		sprite[i] = '.'
		if i < 3 {
			sprite[i] = '#'
		}
	}
}

func spriteUpdate(sprite []byte, center, cycle int) {
	for i, _ := range sprite {
		sprite[i] = '.'
		if i == center-1 && center-1 > -1 {
			sprite[i] = '#'
		}
		if i == center {
			sprite[i] = '#'
		}
		if i == center+1 && center+1 < len(sprite) {
			sprite[i] = '#'
		}
	}
}
