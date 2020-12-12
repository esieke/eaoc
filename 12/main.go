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
	wp := ferry{
		posN: 1,
		posE: 10,
		dir:  0,
	}
	for _, v := range track {
		switch v.action {
		case 'N':
			wp.posN += v.value
		case 'S':
			wp.posN -= v.value
		case 'E':
			wp.posE += v.value
		case 'W':
			wp.posE -= v.value
		case 'L':
			deg := v.value % 360
			trans(deg*-1, &wp)
		case 'R':
			deg := v.value % 360
			trans(deg, &wp)
		case 'F':
			f.posE += wp.posE * v.value
			f.posN += wp.posN * v.value
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

// trans deg: 0, +/-90, +/-180, +/-270
func trans(deg int, wp *ferry) {
	if deg < 0 {
		deg = 360 + deg
	}
	switch deg {
	case 90:
		e := wp.posE * -1
		n := wp.posN
		wp.posN = e
		wp.posE = n
	case 180:
		wp.posE *= -1
		wp.posN *= -1
	case 270:
		e := wp.posE
		n := wp.posN * -1
		wp.posE = n
		wp.posN = e
	}
}
