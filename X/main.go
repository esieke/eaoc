package main

import (
	"bufio"
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

	for s.Scan() {
		l := s.Text()
		println(l)
	}
}
