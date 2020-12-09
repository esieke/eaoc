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

	out := make([]int, 0, 1000)
	for s.Scan() {
		l := s.Text()
		val, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("wrong input datatype skip with error %v\n", err)
			continue
		}
		out = append(out, val)
	}
	preamp := 25
	invalid := 0
	for i := preamp; i < len(out); i++ {
		preampS := out[i-preamp : i]
		valid := false
		for j := 1; j < preamp && !valid; j++ {
			for k := 0; k < j && !valid; k++ {
				sum := preampS[j] + preampS[k]
				if sum == out[i] {
					valid = true
				}
			}
		}
		if !valid {
			invalid = out[i]
			fmt.Printf("result puzzle one: %d\n", out[i])
			break
		}
	}

	// puzzle two
	for i := range out {
		sum := 0
		min, max := out[i], out[i]
		for j := i; j < len(out); j++ {
			if out[j] < min {
				min = out[j]
			}
			if out[j] > max {
				max = out[j]
			}
			sum += out[j]
			if sum == invalid {
				fmt.Printf("result puzzle two %d", min+max)
				return
			}
			if sum > invalid {
				break
			}
		}
	}
}
