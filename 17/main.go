package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cube struct {
	lx int
	ly int
	lz int
	// z y x
	c [][][]bool
}

func new(lx, ly, lz int) *cube {
	c := make([][][]bool, lz)
	for z := 0; z < lz; z++ {
		c[z] = append(c[z], make([][]bool, ly)...)
		for y := 0; y < ly; y++ {
			c[z][y] = append(c[z][y], make([]bool, lx)...)
		}
	}
	return &cube{
		lx: lx,
		ly: ly,
		lz: lz,
		c:  c,
	}
}

func (c *cube) validPos(x, y, z int) bool {
	return x >= 0 && x < c.lx &&
		y >= 0 && y < c.ly &&
		z >= 0 && z < c.lz
}

func (c *cube) setState(x, y, z int, state bool) {
	if c.validPos(x, y, z) {
		c.c[z][y][x] = state
	}
}

func (c *cube) getState(x, y, z int) bool {
	if c.validPos(x, y, z) {
		return c.c[z][y][x]
	}
	return false
}

func (c *cube) copy(dest *cube) {
	for z := 0; z < c.lz; z++ {
		for y := 0; y < c.ly; y++ {
			for x := 0; x < c.lx; x++ {
				dest.setState(x, y, z, c.getState(x, y, z))
			}
		}
	}
}

func (c *cube) numActiveNeighbors(x, y, z int) int {
	r := 0
	for zz := z - 1; zz < z+2; zz++ {
		for yy := y - 1; yy < y+2; yy++ {
			for xx := x - 1; xx < x+2; xx++ {
				if !(xx == x && yy == y && zz == z) {
					if c.getState(xx, yy, zz) {
						r++
					}
				}
			}
		}
	}
	return r
}

func (c *cube) nextState(next *cube) {
	for z := 0; z < c.lz; z++ {
		for y := 0; y < c.ly; y++ {
			for x := 0; x < c.lx; x++ {
				num := c.numActiveNeighbors(x, y, z)
				state := c.getState(x, y, z)
				if !(state && num >= 2 && num <= 3) {
					next.setState(x, y, z, false)
				}
				if !state && num == 3 {
					next.setState(x, y, z, true)
				}
			}
		}
	}
}

func (c *cube) numActive() int {
	r := 0
	for z := 0; z < c.lz; z++ {
		for y := 0; y < c.ly; y++ {
			for x := 0; x < c.lx; x++ {
				if c.getState(x, y, z) {
					r++
				}
			}
		}
	}
	return r
}

func (c *cube) addXYFrame(x, y, z int, frame [][]bool) {
	for yy := 0; yy < len(frame); yy++ {
		for xx := 0; xx < len(frame[0]); xx++ {
			c.setState(xx+x, yy+y, z, frame[yy][xx])
		}
	}
}

func (c *cube) print() int {
	r := 0
	for z := 0; z < c.lz; z++ {
		fmt.Printf("z=%d\n", z)
		for y := 0; y < c.ly; y++ {
			for x := 0; x < c.lx; x++ {
				if c.getState(x, y, z) {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	return r
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	initState := make([][]bool, 0, 1024)
	for s.Scan() {
		l := []byte(s.Text())
		x := make([]bool, len(l))
		for i, s := range l {
			if s == '#' {
				x[i] = true
			}
		}
		initState = append(initState, x)
	}

	// init
	cycles := 6 // 6
	lx := len(initState[0]) + 2*cycles
	ly := len(initState) + 2*cycles
	lz := 1 + 2*cycles
	c1 := new(lx, ly, lz)
	c1.addXYFrame(cycles, cycles, cycles, initState)
	c2 := new(lx, ly, lz)
	c1.copy(c2)

	for ci := 0; ci < cycles; ci++ {
		// c1.print()
		c1.nextState(c2)
		c2.copy(c1)
		// fmt.Println("###################################")
	}

	fmt.Println(c1.numActive())
}
