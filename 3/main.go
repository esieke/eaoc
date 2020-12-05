package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	m := make([][]byte, 0)
	r := 0
	for s.Scan() {
		m = append(m, []byte(s.Text()))
		// Doesn't allocate memory !!! m = append(m, s.Bytes())
	}

	for i := range m {
		if m[i][i*3%len(m[0])] == '#' {
			r++
		}
	}

	fmt.Printf("result puzzle one: %d\n", r)

	slopes := []struct{ r, d int }{
		{r: 1, d: 1},
		{r: 3, d: 1},
		{r: 5, d: 1},
		{r: 7, d: 1},
		{r: 1, d: 2}}

	r, step := 0, 0
	rGlob := 1
	for _, sl := range slopes {
		for i := 0; i < len(m); i += sl.d {
			if m[i][(step*sl.r)%len(m[0])] == '#' {
				r++
			}
			step++
		}
		rGlob = rGlob * r
		r, step = 0, 0
	}

	fmt.Printf("result puzzle two: %d\n", rGlob)
}
