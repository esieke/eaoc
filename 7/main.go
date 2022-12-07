package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	dirs := make(map[string]int)
	pwd := []string{}

	for s.Scan() {
		l := s.Text()

		// cd
		e := regexp.MustCompile(`^\$ cd (.*)`)
		instruction := e.FindStringSubmatch(l)
		if len(instruction) > 1 {
			if instruction[1] == ".." {
				// one up
				pwd = pwd[:len(pwd)-1]
				continue
			}
			// one down
			pwd = append(pwd, instruction[1])
			continue
		}

		// file size
		e = regexp.MustCompile(`^([0-9]+) `)
		instruction = e.FindStringSubmatch(l)
		if len(instruction) > 1 {
			s, err := strconv.Atoi(instruction[1])
			if err != nil {
				panic("size must be an int")
			}
			var key string
			for _, d := range pwd {
				key = fmt.Sprintf("%s/%s", key, d)
				dirs[key] += s
			}
			continue

		}
	}

	size := []int{}
	for _, v := range dirs {
		size = append(size, v)
	}
	sort.Ints(size)

	result := 0
	for _, v := range size {
		if v > 100000 {
			break
		}
		result += v
	}
	fmt.Println(result)
}
