package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	var vals []int
	var val int = 0
	for s.Scan() {
		l := s.Text()
		i, err := strconv.Atoi(l)
		if err != nil {
			vals = append(vals, val)
			val = 0
		} else {
			val += int(i)
		}
	}

	sort.Ints(vals)

	fmt.Print(vals[len(vals)-1] + vals[len(vals)-2] + vals[len(vals)-3])
}
