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
			for i := range lbs {
				// maskAnd all X = 1 all 0 = 0 all 1 = 1
				// maskOr  all X = 0 all 1 = 1 all 0 = 0
				if lbs[35-i] == 'X' || lbs[35-i] == '1' {
					maskAnd |= 1 << uint64(i)
				}
				if lbs[35-i] == '1' {
					maskOr |= 1 << uint64(i)
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
			mem[addr] = val & maskAnd
			mem[addr] |= maskOr
		}
	}

	var result uint64
	for _, v := range mem {
		result += v
	}
	fmt.Println(result)
}
