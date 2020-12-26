package main

import "fmt"

const (
	pubSubNr  = 7
	pubSubDiv = 20201227
)

func main() {

	calcVal := 1

	// cardPub := 5764801
	// doorPub := 17807724
	cardPub := 17607508
	doorPub := 15065270

	cardLoopSize := 0
	for {
		calcVal *= pubSubNr
		calcVal %= pubSubDiv
		cardLoopSize++
		if calcVal == cardPub {
			break
		}
	}

	calcVal = 1
	doorLoopSize := 0
	for {
		calcVal *= pubSubNr
		calcVal %= pubSubDiv
		doorLoopSize++
		if calcVal == doorPub {
			break
		}
	}

	calcVal = 1
	for i := 0; i < cardLoopSize; i++ {
		calcVal *= doorPub
		calcVal %= pubSubDiv
	}

	fmt.Println(calcVal)
}
