package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	var maskAnd uint64
	var maskOr uint64
	var floats []uint64
	mem := make(map[uint64]uint64)
	for s.Scan() {
		l := s.Text()
		ls := strings.Split(l, " = ")
		if len(ls) < 2 {
			fmt.Println("invalid line")
			os.Exit(1)
		}
		if ls[0] == "mask" {
			lbs := []byte(ls[1])
			maskAnd, maskOr = 0, 0
			floats = make([]uint64, 0, 64)
			for i := range lbs {
				// If the bitmask bit is 0, the corresponding memory address bit is unchanged.
				// If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
				// If the bitmask bit is X, the corresponding memory address bit is floating.
				if lbs[35-i] == '0' {
					maskAnd |= 1 << uint64(i)
				}
				if lbs[35-i] == '1' {
					maskOr |= 1 << uint64(i)
				}
				if lbs[35-i] == 'X' {
					floats = append(floats, uint64(i))
				}
			}
		} else {
			memS := strings.Split(ls[0], "[")
			if len(ls) < 2 {
				fmt.Println("invalid mem")
				os.Exit(1)
			}
			trim := strings.TrimSuffix(memS[1], "]")
			addr, err := strconv.ParseUint(trim, 10, 0)
			if err != nil {
				fmt.Printf("parsing mem address failed with error %v\n", err)
				os.Exit(1)
			}
			val, err := strconv.ParseUint(ls[1], 10, 0)
			if err != nil {
				fmt.Printf("parsing mem value failed with error %v\n", err)
				os.Exit(1)
			}

			addr &= maskAnd
			addr |= maskOr
			// 2^(len(floats))
			n := uint64(1) << uint64(len(floats))
			var i uint64
			var resultAddr uint64
			var maskFloat uint64
			for i = 0; i < n; i++ {
				resultAddr = addr
				maskFloat = 0
				for j := 0; j < len(floats); j++ {
					if 1<<j&i > 0 {
						maskFloat |= 1 << floats[j]
					}
				}
				resultAddr |= maskFloat
				// fmt.Printf("%b\n", mem[resultAddr])
				mem[resultAddr] = val
			}
		}
	}

	var result uint64
	for _, v := range mem {
		result += v
	}
	fmt.Println(result)
}
