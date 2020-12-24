package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	x, y int
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	tiles := make([][]pos, 0, 1024)

	for s.Scan() {
		t := make([]pos, 0, 255)
		l := []byte(s.Text())
		for i := 0; i < len(l); i++ {
			// e (2,0), se (1,-1), sw (-1, -1), w (-2, 0), nw (-1, 1), ne (1, 1)
			var p pos
			switch l[i] {
			case 'e':
				p = pos{
					x: 2,
					y: 0,
				}
			case 's':
				i++
				switch l[i] {
				case 'e':
					p = pos{
						x: 1,
						y: -1,
					}
				case 'w':
					p = pos{
						x: -1,
						y: -1,
					}
				}
			case 'w':
				p = pos{
					x: -2,
					y: 0,
				}
			case 'n':
				i++
				switch l[i] {
				case 'w':
					p = pos{
						x: -1,
						y: 1,
					}
				case 'e':
					p = pos{
						x: 1,
						y: 1,
					}
				}
			}
			t = append(t, p)
		}
		tiles = append(tiles, t)
	}

	isBlack := make(map[pos]bool)
	for _, poss := range tiles {
		p := pos{}
		for _, pos := range poss {
			p.x += pos.x
			p.y += pos.y
		}
		_, exist := isBlack[p]
		if exist {
			isBlack[p] = !isBlack[p]
		} else {
			isBlack[p] = true
		}
	}

	result := 0
	for _, val := range isBlack {
		if val {
			result++
		}
	}

	fmt.Println(result)
}
