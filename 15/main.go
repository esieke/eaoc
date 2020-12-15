package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type turn struct {
	num int
	val int
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	mem := make(map[int][]turn)
	lastVal := 0
	ctr := 0

	for s.Scan() {
		l := strings.Split(s.Text(), ",")
		for _, n := range l {
			val, err := strconv.Atoi(n)
			if err != nil {
				fmt.Println("parse string failed with error", err)
			}
			mem[val] = append(mem[val], turn{
				num: ctr + 1,
				val: val,
			})
			lastVal = val
			ctr++
		}
	}

	for {
		val := 0
		if len(mem[lastVal]) > 1 {
			v1 := mem[lastVal][len(mem[lastVal])-1].num
			v2 := mem[lastVal][len(mem[lastVal])-2].num
			val = v1 - v2

		} else {
			val = 0
		}

		if ctr == 2020-1 {
			fmt.Println(val)
			return
		}

		mem[val] = append(mem[val], turn{
			num: ctr + 1,
			val: val,
		})
		lastVal = val
		ctr++
	}

}
