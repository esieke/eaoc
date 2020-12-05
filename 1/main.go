package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read input
	in := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("load input data failed with error %v\n", err)
			os.Exit(1)
		}
		in = append(in, val)
	}

	// puzzle one
	done := false
	for i := range in {
		if done {
			break
		}
		for j := range in {
			if j != i {
				if (in[i] + in[j]) == 2020 {
					fmt.Printf("result puzzle one: %v\n", in[i]*in[j])
					done = true
					break
				}
			}
		}
	}

	// puzzle two
	done = false
	for i := range in {
		if done {
			break
		}
		for j := range in {
			if done {
				break
			}
			for k := range in {
				if k != j && k != i {
					if in[i]+in[j]+in[k] == 2020 {
						fmt.Printf("result puzzle two: %v\n", in[i]*in[j]*in[k])
						done = true
						break
					}
				}
			}
		}
	}
}
