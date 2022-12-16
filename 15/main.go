package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pos struct {
	y, x int
}

type MiMa struct {
	mi, ma int
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	S := Pos{}
	B := Pos{}
	MIN := 0
	MAX := 4000000
	M := make([][]MiMa, MAX+1)
	for i, _ := range M {
		M[i] = []MiMa{{mi: MIN, ma: MAX}}
	}
	_ = M
	for s.Scan() {
		l := s.Text()
		lt := strings.Split(l, "=")
		fmt.Sscanf(lt[1], "%d", &S.x)
		fmt.Sscanf(lt[2], "%d", &S.y)
		fmt.Sscanf(lt[3], "%d", &B.x)
		fmt.Sscanf(lt[4], "%d", &B.y)
		dx := abs(S.x - B.x)
		dy := abs(S.y - B.y)
		r := dx + dy

		for y := -r; y <= r; y++ {
			x := r - abs(y)
			var miN bool
			var maN bool
			ma, yy := S.x+x, S.y+y
			if yy >= MIN && yy <= MAX && ma >= MIN && ma <= MAX {
				maN = true
			}
			mi := S.x - x
			if yy >= MIN && yy <= MAX && mi >= MIN && mi <= MAX {
				miN = true
			}
			if maN == true || miN == true {
				M[yy] = append(M[yy], MiMa{mi: mi, ma: ma})
				continue
			}
			if maN == true {
				M[yy][0].mi = max(ma, M[yy][0].mi)
			}
			if miN == true {
				M[yy][0].ma = min(mi, M[yy][0].ma)
			}
		}
	}
	for y, _ := range M {
		for k := 0; k < len(M[y])-1; k++ {
			for i, _ := range M[y] {
				if i < 1 {
					continue
				}
				if M[y][i].mi <= M[y][0].mi+1 && M[y][i].ma > M[y][0].mi {
					M[y][0].mi = M[y][i].ma
				}
				if M[y][i].ma >= M[y][0].ma-1 && M[y][i].mi < M[y][0].ma {
					M[y][0].ma = M[y][i].mi
				}
			}
		}
		diff := M[y][0].ma - M[y][0].mi
		if diff == 2 {
			fmt.Println((M[y][0].mi+1)*MAX + y)
			return
		}
		if diff > 2 {
			panic("more than one solution")
		}
	}
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
