package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type elf struct {
	min, max uint64
}

type pair struct {
	firstElf, secondElf elf
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	count := 0
	for s.Scan() {
		l := s.Text()
		p := pair{}
		fmt.Sscanf(l, "%d-%d,%d-%d", &p.firstElf.min, &p.firstElf.max, &p.secondElf.min, &p.secondElf.max)
		state := true
		if p.firstElf.min <= p.secondElf.min && p.firstElf.max >= p.secondElf.max {
			state = false
			count += 1
		}
		if p.firstElf.min >= p.secondElf.min && p.firstElf.max <= p.secondElf.max && state {
			count += 1
		}
	}

	log.Println(count)
}
