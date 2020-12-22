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

	player := make([][]byte, 2)
	player[0] = []byte{}
	player[1] = []byte{}
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
				player[0] = append(player[0], byte(num))
			}
			if parserState == 2 {
				player[1] = append(player[1], byte(num))
			}
		}
	}

	winner, player := recursiveCombat(player)

	result := 0
	for i := range player[winner] {
		mul := (len(player[winner]) - i)
		result += mul * int(player[winner][i])
	}

	fmt.Println(result)
}

func recursiveCombat(player [][]byte) (int, [][]byte) {
	// puzzle one
	mem := make([]map[string]bool, 2)
	mem[0] = make(map[string]bool)
	mem[1] = make(map[string]bool)
	winner := 1
	for {
		winner = -1

		// infinite game prevention rule
		if mem[0][string(player[0])] ||
			mem[1][string(player[1])] {
			winner = 0
			break
		}
		// store in hash map
		mem[0][string(player[0])] = true
		mem[1][string(player[1])] = true

		// check for sub game
		//                                    -1 or 0?
		if int(player[0][0]) <= len(player[0])-1 &&
			int(player[1][0]) <= len(player[1])-1 &&
			winner < 0 {
			// copy the deck for sub game
			p := make([][]byte, 2)
			p[0] = []byte{}
			p[1] = []byte{}
			for i, num := range player[0] {
				if i == 0 {
					continue
				}
				if i > int(player[0][0]) {
					break
				}
				p[0] = append(p[0], num)
			}
			for i, num := range player[1] {
				if i == 0 {
					continue
				}
				if i > int(player[1][0]) {
					break
				}
				p[1] = append(p[1], num)
			}
			winner, _ = recursiveCombat(p)
		}

		// 3. rule
		if winner < 0 {
			winner = 1
			if player[0][0] > player[1][0] {
				winner = 0
			}
		}

		// calculate decks
		player[winner] = append(player[winner], player[winner][0], player[(winner+1)%2][0])
		player[winner] = player[winner][1:len(player[winner])]
		player[(winner+1)%2] = player[(winner+1)%2][1:len(player[(winner+1)%2])]

		if len(player[0]) == 0 || len(player[1]) == 0 {
			break
		}
	}
	return winner, player
}
