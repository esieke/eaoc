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
	BorderTop    borderEnum = iota
	BorderLeft   borderEnum = iota
	BorderBottom borderEnum = iota
	BorderRight  borderEnum = iota
)

type tile struct {
	id  int
	dim int // 10 x 10 dim = 10
	raw [][]byte
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

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
			t = newTile(id)
			tiles = append(tiles, t)
			continue
		}
		// fill
		t.addRow([]byte(l))
	}

	// tiles map
	mapTiles := make([][]*tile, 0, 128)

	// search top/left corner
	t = tiles[0]
	var next *tile
	dir := BorderLeft
	for {
		next = t.getNext(dir, tiles)
		if next != nil {
			t = next
		}
		if next == nil && dir == BorderTop {
			break
		}
		if next == nil && dir == BorderLeft {
			dir = BorderTop
		}
	}

	// fill first column
	mapTilesR := make([]*tile, 0, 128)
	mapTilesR = append(mapTilesR, t)
	mapTiles = append(mapTiles, mapTilesR)
	dir = BorderBottom
	for {
		next = t.getNext(dir, tiles)
		if next != nil {
			t = next
			mapTilesR = make([]*tile, 0, 128)
			mapTilesR = append(mapTilesR, t)
			mapTiles = append(mapTiles, mapTilesR)
		}
		if next == nil {
			break
		}
	}

	// fill rows
	dir = BorderRight
	for r := range mapTiles {
		t = mapTiles[r][0]
		for {
			next = t.getNext(dir, tiles)
			if next != nil {
				t = next
				mapTiles[r] = append(mapTiles[r], t)
			}
			if next == nil {
				break
			}
		}
	}

	if len(mapTiles) != len(mapTiles[0]) {
		panic("expect square")
	}

	picture := newTile(0)
	picture.dim = len(mapTiles) * (mapTiles[0][0].dim - 2)
	picture.initRaw()
	for r := range mapTiles {
		for c := range mapTiles[r] {
			dim := mapTiles[r][c].dim - 2
			for rt := 1; rt <= dim; rt++ {
				for ct := 1; ct <= dim; ct++ {
					rp := (r * dim) + rt - 1
					cp := (c * dim) + ct - 1
					picture.raw[rp][cp] = mapTiles[r][c].raw[rt][ct]
				}
			}
		}
	}

	seaMonsters := 0
	roughness := 0
	for i := 0; i < 2; i++ {
		for k := 0; k < 4; k++ {
			m := picture.searchSeaMonsters()
			if m > seaMonsters {
				seaMonsters = m
				for r := range picture.raw {
					for c := range picture.raw[r] {
						if picture.raw[r][c] == '#' {
							roughness++
						}
					}
				}
				roughness -= (seaMonsters * 15)
			}
			picture.rot()
		}
		picture.flip()
	}
	fmt.Println(roughness)
}

func newTile(id int) *tile {
	return &tile{
		id: id,
	}
}

func (t *tile) addRow(r []byte) {
	if len(t.raw) == 0 {
		t.dim = len(r)
		t.raw = make([][]byte, 0, 128)
	}
	t.raw = append(t.raw, r)
}

func (t *tile) initRaw() {
	t.raw = make([][]byte, t.dim)
	for r := range t.raw {
		t.raw[r] = make([]byte, t.dim)
	}
}

func (t *tile) copy() *tile {
	raw := make([][]byte, t.dim)
	for r := range t.raw {
		raw[r] = make([]byte, t.dim)
		for c := range t.raw[r] {
			raw[r][c] = t.raw[r][c]
		}
	}
	return &tile{
		id:  t.id,
		dim: t.dim,
		raw: raw,
	}
}

func (t *tile) rot() {
	buf := t.copy()
	for r := range t.raw {
		for c := range t.raw[r] {
			t.raw[c][buf.dim-1-r] = buf.raw[r][c]
		}
	}
}

func (t *tile) flip() {
	buf := t.copy()
	for r := range t.raw {
		for c := range t.raw[r] {
			t.raw[r][buf.dim-1-c] = buf.raw[r][c]
		}
	}
}

func (t *tile) getBorder(b borderEnum) string {
	border := make([]byte, t.dim)

	switch b {
	case BorderTop:
		for i := 0; i < t.dim; i++ {
			border[i] = t.raw[0][i]
		}
	case BorderLeft:
		for i := 0; i < t.dim; i++ {
			border[i] = t.raw[i][0]
		}
	case BorderBottom:
		for i := 0; i < t.dim; i++ {
			border[i] = t.raw[t.dim-1][i]
		}
	case BorderRight:
		for i := 0; i < t.dim; i++ {
			border[i] = t.raw[i][t.dim-1]
		}
	}
	return string(border)
}

func (t *tile) print() {
	for _, row := range t.raw {
		fmt.Println(string(row))
	}
}

func (t *tile) getNext(b borderEnum, tiles []*tile) *tile {
	tBorder := string(t.getBorder(b))

	for ti := range tiles {
		if t.id != tiles[ti].id {
			for i := 0; i < 2; i++ {
				for k := 0; k < 4; k++ {
					tCompBorder := tiles[ti].getBorder((b + 2) % 4)
					if tBorder == tCompBorder {
						return tiles[ti]
					}
					tiles[ti].rot()
				}
				tiles[ti].flip()
			}
		}
	}
	return nil
}

func (t *tile) searchSeaMonsters() int {
	//                   #
	// #    ##    ##    ###
	//  #  #  #  #  #  #

	// line 0 to line len()-3 column 18 to len()-2
	ret := 0
	for c := 0; c < t.dim-3; c++ {
		for r := 18; r < t.dim-2; r++ {
			if t.raw[c][r] == '#' &&
				// next row
				t.raw[c+1][r-18] == '#' &&
				t.raw[c+1][r-13] == '#' &&
				t.raw[c+1][r-12] == '#' &&
				t.raw[c+1][r-7] == '#' &&
				t.raw[c+1][r-6] == '#' &&
				t.raw[c+1][r-1] == '#' &&
				t.raw[c+1][r-0] == '#' &&
				t.raw[c+1][r+1] == '#' &&
				// next row
				t.raw[c+2][r-17] == '#' &&
				t.raw[c+2][r-14] == '#' &&
				t.raw[c+2][r-11] == '#' &&
				t.raw[c+2][r-8] == '#' &&
				t.raw[c+2][r-5] == '#' &&
				t.raw[c+2][r-2] == '#' {
				ret++
			}
		}
	}
	return ret
}
