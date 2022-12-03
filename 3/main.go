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

	prioritiesSum := 0
	for s.Scan() {
		l := s.Text()
		bl := []byte(l)
		blLen := len(bl)
		m := make(map[byte]bool)
		for i := 0; i < blLen/2; i++ {
			m[bl[i]] = true
		}
		for i := blLen / 2; i < blLen; i++ {
			v := m[bl[i]]
			if v {
				if bl[i] >= 'a' && bl[i] <= 'z' {
					prioritiesSum += int(bl[i] - 'a' + 1)
				}
				if bl[i] >= 'A' && bl[i] <= 'Z' {
					prioritiesSum += int(bl[i] - 'A' + 27)
				}
				break
			}
		}
	}
	log.Println(prioritiesSum)
}
