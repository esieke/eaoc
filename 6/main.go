package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	var bl []byte
	for s.Scan() {
		l := s.Text()
		bl = []byte(l)
	}

	n := 4
	lastN := make([]byte, n)
	for i, _ := range bl {
		if i < n {
			continue
		}
		for k, _ := range lastN {
			lastN[k] = bl[i-n+k]
		}
		sort.Slice(lastN, func(k, l int) bool {
			return lastN[k] < lastN[l]
		})
		sop := true
		for k, _ := range lastN {
			if k < 1 {
				continue
			}
			if lastN[k] == lastN[k-1] {
				sop = false
				break
			}
		}
		if sop {
			fmt.Println(i)
			return
		}
	}
}
