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

	var ls [][]byte
	prioritiesSum := 0
	for s.Scan() {
		l := s.Text()
		bl := []byte(l)
		ls = append(ls, bl)
	}
	var priorities []int
	for i := 0; i < len(ls); i += 3 {
		m1 := make(map[byte]bool)
		for k := 0; k < len(ls[i]); k++ {
			m1[ls[i][k]] = true
		}
		m2 := make(map[byte]bool)
		for k := 0; k < len(ls[i+1]); k++ {
			m2[ls[i+1][k]] = true
		}
		for k := 0; k < len(ls[i+2]); k++ {
			v1 := m1[ls[i+2][k]]
			v2 := m2[ls[i+2][k]]
			if v1 && v2 {
				if ls[i+2][k] >= 'a' && ls[i+2][k] <= 'z' {
					priorities = append(priorities, int(ls[i+2][k]-'a')+1)
					prioritiesSum += int(ls[i+2][k]-'a') + 1
				}
				if ls[i+2][k] >= 'A' && ls[i+2][k] <= 'Z' {
					priorities = append(priorities, int(ls[i+2][k]-'A')+27)
					prioritiesSum += int(ls[i+2][k]-'A') + 27
				}
				break
			}
		}
	}
	log.Println(prioritiesSum)
}
