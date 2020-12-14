package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	m := make([][]byte, 0, 1024)
	for s.Scan() {
		l := s.Text()
		m = append(m, []byte(l))
	}

	mNext := make([][]byte, len(m), 1024)
	for y, xs := range m {
		mNext[y] = make([]byte, len(xs))
	}
	deepCopy(mNext, m)

	dirs := [][]int{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1},
	}
	for {
		for y, xs := range mNext {
			for x := range xs {
				free, occup := 0, 0
				for _, dir := range dirs {
					xi, yi := x, y
					for {
						xi += dir[0]
						yi += dir[1]
						if xi < 0 ||
							xi > len(xs)-1 ||
							yi < 0 ||
							yi > len(m)-1 {
							break
						}
						if m[yi][xi] == 'L' {
							free++
							break
						}
						if m[yi][xi] == '#' {
							occup++
							break
						}
					}
				}

				if occup == 0 && m[y][x] == 'L' {
					mNext[y][x] = '#'
				}
				if occup >= 5 && m[y][x] == '#' {
					mNext[y][x] = 'L'
				}
			}
		}

		//print(m, mNext)
		if compare(mNext, m) {
			break
		}

		deepCopy(m, mNext)
	}
	fmt.Println(count(m))
}

func deepCopy(dst, src [][]byte) {
	for y, xs := range src {
		for x := range xs {
			b := src[y][x]
			dst[y][x] = b
		}
	}
}

func count(m [][]byte) (ret int) {
	for y, xs := range m {
		for x := range xs {
			if m[y][x] == '#' {
				ret++
			}
		}
	}
	return ret
}

func compare(a, b [][]byte) (ret bool) {
	ret = true
	for y, xs := range b {
		for x := range xs {
			if a[y][x] != b[y][x] {
				ret = false
				break
			}
		}
		if !ret {
			break
		}
	}
	return ret
}

func print(m, mNext [][]byte) {
	fmt.Println("\033[2J")
	for y, xs := range m {
		for x := range xs {
			fmt.Printf("%s", string(m[y][x]))
		}
		fmt.Printf("\t")
		for x := range xs {
			fmt.Printf("%s", string(mNext[y][x]))
		}
		fmt.Printf("\n")
	}
	time.Sleep(time.Millisecond * 250)
}
