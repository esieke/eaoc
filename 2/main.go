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

	totalScore := 0
	for s.Scan() {
		l := s.Text()
		totalScore += getScore2(l)
	}
	fmt.Println(totalScore)
}

func getScore(game string) int {
	// Rock, Paper, Scissors
	// 'A' == 65
	// 'X' == 88
	b := []byte(game)
	// draw
	if b[2]-87 == b[0]-64 {
		return int(b[2] - 87 + 3)
	}
	// won
	if (b[2]-87 == 1 && b[0]-64 == 3) ||
		(b[2]-87 == 2 && b[0]-64 == 1) ||
		(b[2]-87 == 3 && b[0]-64 == 2) {
		return int(b[2] - 87 + 6)
	}
	// lost
	return int(b[2] - 87 + 0)
}

// X lose
// Y draw
// Z win
func getScore2(game string) int {
	// Rock, Paper, Scissors
	// 'A' == 65
	// 'X' == 88
	b := []byte(game)
	if b[2]-87 == 1 {
		// lose
		my := int(b[0] - 64 - 1)
		if my < 1 {
			my = 3
		}
		return my + 0
	}
	if b[2]-87 == 2 {
		// draw
		return int(b[0] - 64 + 3)
	}
	// win
	my := int(b[0] - 64 + 1)
	if my > 3 {
		my = 1
	}
	return my + 6
}
