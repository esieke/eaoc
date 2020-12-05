package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	seatID := 0
	var store [128 * 8]bool
	for s.Scan() {
		l := []byte(s.Text())

		row, column, maxRow := 0, 0, 6
		if len(l) != 10 {
			log.Fatal("error boarding pass not valid")
		}
		for i := 0; i <= maxRow; i++ {
			if l[i] == 'B' {
				row |= 1 << (maxRow - i)
			}
		}
		maxColumn := 2
		for i := 0; i <= maxColumn; i++ {
			if l[i+maxRow+1] == 'R' {
				column |= 1 << (maxColumn - i)
			}
		}
		sID := row*8 + column
		if sID > seatID {
			seatID = sID
		}

		// puzzle two
		store[sID] = true

	}

	fmt.Printf("result puzzle one: %d\n", seatID)

	for i := range store {
		if i > 0 && i < 128*8-1 {
			if store[i-1] && !store[i] && store[i+1] {
				seatID = i
			}
		}
	}

	fmt.Printf("result puzzle two: %d\n", seatID)
}
