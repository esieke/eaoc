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

type tile struct {
	color bool
	p     pos
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
	xMin := 0
	xMax := 0
	yMin := 0
	yMax := 0
	for _, poss := range tiles {
		p := pos{}
		for _, pos := range poss {
			p.x += pos.x
			p.y += pos.y
		}
		if p.x < xMin {
			xMin = p.x
		}
		if p.x > xMax {
			xMax = p.x
		}
		if p.y < yMin {
			yMin = p.y
		}
		if p.y > yMax {
			yMax = p.y
		}
		_, exist := isBlack[p]
		if exist {
			isBlack[p] = !isBlack[p]
		} else {
			isBlack[p] = true
		}
	}

	// initialize grid
	xOffset := xMax - xMin + 200 + 20
	yOffset := yMax - yMin + 1 + 200 + 20
	tilesGrid := make([][]tile, yOffset)
	for y := range tilesGrid {
		x := make([]tile, xOffset)
		tilesGrid[y] = x
	}
	xOffset /= 2
	yOffset /= 2
	for y := range tilesGrid {
		for x := range tilesGrid[y] {
			p := pos{
				x: x - xOffset,
				y: y - yOffset,
			}
			black, exist := isBlack[p]
			if exist {
				tilesGrid[y][x].color = black
			}
		}
	}

	// initialize buffer
	buf := make([][]tile, len(tilesGrid))
	for y := range tilesGrid {
		x := make([]tile, len(tilesGrid[0]))
		buf[y] = x
	}
	// result
	result := 0
	for day := 0; day < 101; day++ {
		result = 0
		for y := range tilesGrid {
			for x := range tilesGrid[y] {
				if tilesGrid[y][x].color {
					result++
				}
			}
		}

		// copy to buffer
		for y := range tilesGrid {
			for x := range tilesGrid[y] {
				buf[y][x] = tilesGrid[y][x]
			}
		}

		// calc state
		for y := range tilesGrid {
			if y == 0 || y == len(tilesGrid)-1 {
				continue
			}

			// y == 0 --> x 0, 2, 4, 6
			// y == 1 --> x 1, 3, 5, 7
			offset := 2
			if y%2 == 0 {
				offset++
			}

			for x := offset; x < len(tilesGrid[0])-offset; x += 2 {
				blackAdjacent := 0
				// e (2,0), se (1,-1), sw (-1, -1), w (-2, 0), nw (-1, 1), ne (1, 1)
				if tilesGrid[y][x+2].color {
					blackAdjacent++
				}
				if tilesGrid[y-1][x+1].color {
					blackAdjacent++
				}
				if tilesGrid[y-1][x-1].color {
					blackAdjacent++
				}
				if tilesGrid[y][x-2].color {
					blackAdjacent++
				}
				if tilesGrid[y+1][x-1].color {
					blackAdjacent++
				}
				if tilesGrid[y+1][x+1].color {
					blackAdjacent++
				}
				// Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
				if tilesGrid[y][x].color && (blackAdjacent == 0 || blackAdjacent > 2) {
					buf[y][x].color = false
				}
				// Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
				if !tilesGrid[y][x].color && blackAdjacent == 2 {
					buf[y][x].color = true
				}
			}
		}

		// copy buffer
		for y := range tilesGrid {
			for x := range tilesGrid[y] {
				tilesGrid[y][x] = buf[y][x]
			}
		}
	}
	fmt.Println(result)
}
