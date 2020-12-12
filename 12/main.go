package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type instr struct {
	action byte
	value  int
}

type ferry struct {
	posN int
	posE int
	dir  int
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	track := make([]instr, 0, 1024)
	for s.Scan() {
		l := []byte(s.Text())
		v, err := strconv.Atoi(string(l[1:]))
		if err != nil {
			fmt.Println("error parsing line")
			continue
		}
		track = append(track, instr{
			action: l[0],
			value:  v,
		})
	}

	f := ferry{
		posN: 0,
		posE: 0,
		dir:  0,
	}
	for _, v := range track {
		switch v.action {
		case 'N':
			f.posN += v.value
		case 'S':
			f.posN -= v.value
		case 'E':
			f.posE += v.value
		case 'W':
			f.posE -= v.value
		case 'L':
			f.dir -= v.value
			if f.dir < 0 {
				f.dir = 360 - (f.dir*-1)%360
			} else {
				f.dir = f.dir % 360
			}
		case 'R':
			f.dir += v.value
			f.dir = f.dir % 360
		case 'F':
			switch f.dir {
			case 0:
				f.posE += v.value
			case 90:
				f.posN -= v.value
			case 180:
				f.posE -= v.value
			case 270:
				f.posN += v.value
			}
		}
	}

	if f.posE < 0 {
		f.posE *= -1
	}
	if f.posN < 0 {
		f.posN *= -1
	}

	fmt.Println(f.posE + f.posN)
}
