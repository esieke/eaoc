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

	var group map[byte]int
	result := 0
	group = make(map[byte]int)
	for s.Scan() {
		l := []byte(s.Text())
		if len(l) == 0 {
			for k, y := range group {
				if k != '!' {
					if group['!'] == y {
						result++
					}
				}
			}
			group = make(map[byte]int)
			continue
		}
		group['!']++
		for _, b := range l {
			group[b]++
		}

	}
	// last group
	for k, y := range group {
		if k != '!' {
			if group['!'] == y {
				result++
			}
		}
	}

	fmt.Printf("result puzzle one: %d\n", result)
}
