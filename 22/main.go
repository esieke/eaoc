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

	player := make([][]int, 2)
	player[0] = []int{}
	player[1] = []int{}
	parserState := 0
	for s.Scan() {
		l := s.Text()
		if l == "" {
			continue
		}
		if l == "Player 1:" {
			parserState = 1
			continue
		}
		if l == "Player 2:" {
			parserState = 2
			continue
		}
		if parserState == 1 || parserState == 2 {
			num, err := strconv.Atoi(l)
			if err != nil {
				fmt.Printf("convert to integer failed with error %s\n", err)
				continue
			}
			if parserState == 1 {
				player[0] = append(player[0], num)
			}
			if parserState == 2 {
				player[1] = append(player[1], num)
			}
		}
	}

	// puzzle one
	winner := 1
	for {
		winner = 1
		if player[0][0] > player[1][0] {
			winner = 0
		}
		player[winner] = append(player[winner], player[winner][0], player[(winner+1)%2][0])
		player[winner] = player[winner][1:len(player[winner])]
		player[(winner+1)%2] = player[(winner+1)%2][1:len(player[(winner+1)%2])]
		if len(player[0]) == 0 || len(player[1]) == 0 {
			break
		}
	}

	result := 0
	for i := range player[winner] {
		mul := (len(player[winner]) - i)
		result += mul * player[winner][i]
	}

	fmt.Println(result)
}
