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

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	S := Pos{}
	B := Pos{}
	Bs := []Pos{}
	M := make(map[Pos]bool)
	//Y := 10
	Y := 2000000
	for s.Scan() {
		l := s.Text()
		lt := strings.Split(l, "=")
		fmt.Sscanf(lt[1], "%d", &S.x)
		fmt.Sscanf(lt[2], "%d", &S.y)
		fmt.Sscanf(lt[3], "%d", &B.x)
		fmt.Sscanf(lt[4], "%d", &B.y)
		Bs = append(Bs, B)
		dx := abs(S.x - B.x)
		dy := abs(S.y - B.y)
		r := dx + dy
		if S.y+r >= Y && S.y-r <= Y {
			for x := -1 * (r - abs(S.y-Y)); x <= (r - abs(S.y-Y)); x++ {
				M[Pos{y: Y, x: S.x + x}] = true
			}
		}

	}
	for _, b := range Bs {
		M[b] = false
	}
	result := 0
	for k, v := range M {
		if k.y == Y && v {
			//fmt.Println(k)
			result++
		}
	}
	fmt.Println(result)
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}
