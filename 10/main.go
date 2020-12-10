package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	parts := make([][]int, 0, 1024)
	var last int
	for i := range joltAdapt {
		if i == 0 {
			continue
		}
		if joltAdapt[i]-joltAdapt[i-1] == 3 {
			parts = append(parts, joltAdapt[last:i])
			last = i
		}
	}

	result := 1
	for _, joltAdapGroup := range parts {
		if len(joltAdapGroup) <= 2 {
			continue
		}
		partResult := 0
		solutions := int(math.Pow(2.0, float64(len(joltAdapGroup)-2)))
		for bitMask := 0; bitMask < solutions; bitMask++ {
			last = joltAdapGroup[0]
			invalid := false
			for i := 1; i < len(joltAdapGroup); i++ {
				if 1<<(i-1)&bitMask != 0 || i == len(joltAdapGroup)-1 {
					if joltAdapGroup[i]-last > 3 {
						invalid = true
						break
					}
					last = joltAdapGroup[i]
				}
			}
			if invalid {
				continue
			}
			partResult++
		}
		result *= partResult
	}
	fmt.Println(result)
}
