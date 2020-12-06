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

	var group map[byte]bool
	result := 0
	group = make(map[byte]bool)
	for s.Scan() {
		l := []byte(s.Text())
		if len(l) == 0 {
			for _, y := range group {
				if y {
					result++
				}
			}
			group = make(map[byte]bool)
			continue
		}
		for _, b := range l {
			group[b] = true
		}

	}
	// last group
	for _, y := range group {
		if y {
			result++
		}
	}

	fmt.Printf("result puzzle one: %d\n", result)
}
