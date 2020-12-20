package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type borderEnum int

const (
	BorderTop          borderEnum = iota
	BorderLeft         borderEnum = iota
	BorderBottom       borderEnum = iota
	BorderRight        borderEnum = iota
	BorderFlipedTop    borderEnum = iota // inverse BorderBottom
	BorderFlipedLeft   borderEnum = iota
	BorderFlipedBottom borderEnum = iota // inverse BorderTop
	BorderFlipedRight  borderEnum = iota
	BorderNum          borderEnum = iota // Number of borders
)

type tile struct {
	id      int
	rotCW   int
	dim     int // 10 x 10
	borders []uint16
	raw     []string
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	tileLine := 0
	tiles := []*tile{}
	var t *tile
	for s.Scan() {
		l := s.Text()
		if l == "" {
			continue
		}
		if strings.Contains(l, "Tile") {
			lp := strings.Split(l, " ")
			lp = strings.Split(lp[1], ":")
			id, err := strconv.Atoi(lp[0])
			if err != nil {
				fmt.Printf("parse tile id failed with error %v\n", err)
				continue
			}
			tileLine = 0
			t = newTile(id, 10)
			continue
		}
		// fill borders
		lb := []byte(l)
		if tileLine == 0 {
			for i, b := range lb {
				if b == '#' {
					t.borders[BorderTop] |= 1 << i
					t.borders[BorderFlipedBottom] |= 1 << (t.dim - 1 - i)
				}
			}
		}
		if tileLine == t.dim-1 {
			for i, b := range lb {
				if b == '#' {
					t.borders[BorderBottom] |= 1 << (t.dim - 1 - i)
					t.borders[BorderFlipedTop] |= 1 << i
				}
			}
		}
		if lb[0] == '#' {
			t.borders[BorderLeft] |= 1 << (t.dim - 1 - tileLine)
			t.borders[BorderFlipedLeft] |= 1 << tileLine
		}
		if lb[t.dim-1] == '#' {
			t.borders[BorderRight] |= 1 << tileLine
			t.borders[BorderFlipedRight] |= 1 << (t.dim - 1 - tileLine)
		}

		// fill raw
		t.raw[tileLine] = l

		if tileLine == t.dim-1 {
			tiles = append(tiles, t)
		}

		tileLine++
	}
	// for _, v := range t.borders {
	// 	fmt.Printf("%010b\n", v)
	// }

	result := 1
	for _, t := range tiles {
		num := findNumBorders(t, tiles)
		if num == 2 {
			result *= t.id
			fmt.Println(t.id)
		}

	}

	fmt.Println(result)
}

func newTile(id, dim int) *tile {
	borders := make([]uint16, int(BorderNum))
	raw := make([]string, dim)
	return &tile{
		id:      id,
		dim:     dim,
		borders: borders,
		raw:     raw,
	}
}

func findNumBorders(t *tile, ts []*tile) int {
	ret := 0
	for _, tin := range ts {
		if tin.id != t.id {
			for i := 0; i < 4; i++ {
				for _, b := range tin.borders {
					if t.borders[i] == b {
						ret++
					}
				}
			}
		}
	}
	return ret
}
