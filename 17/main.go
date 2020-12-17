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
	lw int
	// w z y x
	c [][][][]bool
}

func new(lx, ly, lz, lw int) *cube {
	c := make([][][][]bool, lw)
	for w := 0; w < lw; w++ {
		c[w] = append(c[w], make([][][]bool, ly)...)
		for z := 0; z < lz; z++ {
			c[w][z] = append(c[w][z], make([][]bool, ly)...)
			for y := 0; y < ly; y++ {
				c[w][z][y] = append(c[w][z][y], make([]bool, lx)...)
			}
		}
	}
	return &cube{
		lx: lx,
		ly: ly,
		lz: lz,
		lw: lw,
		c:  c,
	}
}

func (c *cube) validPos(x, y, z, w int) bool {
	return x >= 0 && x < c.lx &&
		y >= 0 && y < c.ly &&
		z >= 0 && z < c.lz &&
		w >= 0 && w < c.lw
}

func (c *cube) setState(x, y, z, w int, state bool) {
	if c.validPos(x, y, z, w) {
		c.c[w][z][y][x] = state
	}
}

func (c *cube) getState(x, y, z, w int) bool {
	if c.validPos(x, y, z, w) {
		return c.c[w][z][y][x]
	}
	return false
}

func (c *cube) copy(dest *cube) {
	for w := 0; w < c.lw; w++ {
		for z := 0; z < c.lz; z++ {
			for y := 0; y < c.ly; y++ {
				for x := 0; x < c.lx; x++ {
					dest.setState(x, y, z, w, c.getState(x, y, z, w))
				}
			}
		}
	}
}

func (c *cube) numActiveNeighbors(x, y, z, w int) int {
	r := 0
	for ww := w - 1; ww < w+2; ww++ {
		for zz := z - 1; zz < z+2; zz++ {
			for yy := y - 1; yy < y+2; yy++ {
				for xx := x - 1; xx < x+2; xx++ {
					if !(xx == x && yy == y && zz == z && ww == w) {
						if c.getState(xx, yy, zz, ww) {
							r++
						}
					}
				}
			}
		}
	}
	return r
}

func (c *cube) nextState(next *cube) {
	for w := 0; w < c.lw; w++ {
		for z := 0; z < c.lz; z++ {
			for y := 0; y < c.ly; y++ {
				for x := 0; x < c.lx; x++ {
					num := c.numActiveNeighbors(x, y, z, w)
					state := c.getState(x, y, z, w)
					if !(state && num >= 2 && num <= 3) {
						next.setState(x, y, z, w, false)
					}
					if !state && num == 3 {
						next.setState(x, y, z, w, true)
					}
				}
			}
		}
	}
}

func (c *cube) numActive() int {
	r := 0
	for w := 0; w < c.lw; w++ {
		for z := 0; z < c.lz; z++ {
			for y := 0; y < c.ly; y++ {
				for x := 0; x < c.lx; x++ {
					if c.getState(x, y, z, w) {
						r++
					}
				}
			}
		}
	}
	return r
}

func (c *cube) addXYFrame(x, y, z int, frame [][]bool) {
	w := z
	for yy := 0; yy < len(frame); yy++ {
		for xx := 0; xx < len(frame[0]); xx++ {
			c.setState(xx+x, yy+y, z, w, frame[yy][xx])
		}
	}
}

func (c *cube) print() int {
	r := 0
	for w := 0; w < c.lw; w++ {
		for z := 0; z < c.lz; z++ {
			fmt.Printf("w=%d, z=%d\n", w, z)
			for y := 0; y < c.ly; y++ {
				for x := 0; x < c.lx; x++ {
					if c.getState(x, y, z, w) {
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
	lw := lz
	c1 := new(lx, ly, lz, lw)
	c1.addXYFrame(cycles, cycles, cycles, initState)
	c2 := new(lx, ly, lz, lw)
	c1.copy(c2)

	for ci := 0; ci < cycles; ci++ {
		// c1.print()
		c1.nextState(c2)
		c2.copy(c1)
		// fmt.Println("###################################")
	}

	fmt.Println(c1.numActive())
}
