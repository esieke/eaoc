package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	var max uint64 = 0
	var val uint64 = 0
	for s.Scan() {
		l := s.Text()
		i, err := strconv.Atoi(l)
		if err != nil {
			if val > max {
				max = val
			}
			val = 0
		} else {
			val += uint64(i)
		}
	}
	fmt.Print(max)
}
