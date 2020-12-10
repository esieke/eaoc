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

	joltAdapt := make([]int, 0, 1024)
	joltAdapt = append(joltAdapt, 0)
	for s.Scan() {
		l := s.Text()
		val, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("wrong input datatype skip with error %v\n", err)
			continue
		}
		joltAdapt = append(joltAdapt, val)
	}

	sort.Ints(joltAdapt)
	joltAdapt = append(joltAdapt, joltAdapt[len(joltAdapt)-1]+3)

	diffCtr := make(map[int]int)
	for i := range joltAdapt {
		if i == 0 {
			continue
		}
		diffCtr[joltAdapt[i]-joltAdapt[i-1]]++
	}
	fmt.Println(diffCtr[1] * diffCtr[3])
}
