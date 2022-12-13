package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	y, x int
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	in := bufio.NewScanner(input)

	M := [][]byte{} // Map
	T := [][]int{}  // Track
	S := []pos{{}}  // State
	E := pos{}
	y := 0
	for in.Scan() {
		l := in.Text()
		m := []byte(l)
		t := make([]int, len(l))
		T = append(T, t)
		for x, _ := range m {
			if m[x] == 'S' {
				m[x] = 'a'
			}
			if m[x] == 'E' {
				E.x, E.y = x, y
				m[x] = 'z'
			}
		}
		M = append(M, m)
		y++
	}

	result := 505 // initialize with result from puzzle 1
	for Y, _ := range M {
		for X, _ := range M[Y] {
			if M[Y][X] == 'a' {
				for y, _ := range T {
					for x, _ := range T[y] {
						T[y][x] = 0
					}
				}
				S = []pos{{x: X, y: Y}}
				step := 1
				done := false
				for step = 1; step < 505; step++ {
					s := make([]pos, len(S))
					for i, v := range S {
						s[i] = v
					}
					S = []pos{}

					for i, _ := range s {
						m := M[s[i].y][s[i].x]
						// left
						y := s[i].y
						x := s[i].x - 1
						if x > -1 {
							if x == E.x && E.y == y {
								done = true
								break
							}
							if (M[y][x] == m || M[y][x] == m+1 || M[y][x] < m) && T[y][x] < 1 {
								S = append(S, pos{y: y, x: x})
								T[y][x] = step
							}
						}
						// right
						x = s[i].x + 1
						if x < len(M[0]) {
							if x == E.x && E.y == y {
								done = true
								break
							}
							if (M[y][x] == m || M[y][x] == m+1 || M[y][x] < m) && T[y][x] < 1 {
								S = append(S, pos{y: y, x: x})
								T[y][x] = step
							}
						}
						// up
						x = s[i].x
						y = s[i].y - 1
						if y > -1 {
							if x == E.x && E.y == y {
								done = true
								break
							}
							if (M[y][x] == m || M[y][x] == m+1 || M[y][x] < m) && T[y][x] < 1 {
								S = append(S, pos{y: y, x: x})
								T[y][x] = step
							}
						}
						// down
						y = s[i].y + 1
						if y < len(M) {
							if x == E.x && E.y == y {
								done = true
								break
							}
							if (M[y][x] == m || M[y][x] == m+1 || M[y][x] < m) && T[y][x] < 1 {
								S = append(S, pos{y: y, x: x})
								T[y][x] = step
							}
						}
					}
					if done {
						break
					}
				}
				if step < result {
					result = step
				}
			}
		}
	}
	fmt.Println(result)
}
