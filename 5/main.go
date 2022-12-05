package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	var stacks [9][]byte
	var rules [][3]int
	state := 0
	for s.Scan() {
		l := s.Text()
		if l == "" {
			state = 1
			continue
		}

		// stacks
		if state == 0 {
			bl := []byte(l)
			stackN := 0
			for i := 1; i < len(bl); i += 4 {
				if bl[i] >= 'A' && bl[i] <= 'Z' {
					stacks[stackN] = append([]byte{bl[i]}, stacks[stackN]...)
				}
				stackN += 1
			}
		}

		// rules
		if state == 1 {
			var move, from, to int
			fmt.Sscanf(l, "move %d from %d to %d", &move, &from, &to)
			rules = append(rules, [3]int{move, from - 1, to - 1})
		}
	}

	for _, r := range rules {
		stacks[r[2]] = append(stacks[r[2]], stacks[r[1]][len(stacks[r[1]])-r[0]:len(stacks[r[1]])]...)
		stacks[r[1]] = stacks[r[1]][:len(stacks[r[1]])-r[0]]
	}

	for _, v := range stacks {
		if len(v) > 0 {
			fmt.Printf("%s", string(v[len(v)-1]))
		}
	}
}
