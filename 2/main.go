package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	resultOne := 0
	resultTwo := 0
	for scanner.Scan() {
		var lowest, highest int
		var letter []byte
		var password []byte
		fmt.Sscanf(scanner.Text(), "%d-%d %1s:  %s", &lowest, &highest, &letter, &password)

		// puzzle one
		hist := make(map[byte]int)
		for _, v := range password {
			hist[v]++
		}
		if hist[letter[0]] >= lowest && hist[letter[0]] <= highest {
			resultOne++
		}

		// puzzle two
		// Range check
		if lowest-1 >= 0 &&
			lowest-1 < len(password) &&
			highest-1 >= 0 &&
			highest-1 < len(password) {
			if (password[lowest-1] == letter[0] &&
				password[highest-1] != letter[0]) ||
				(password[lowest-1] != letter[0] &&
					password[highest-1] == letter[0]) {
				resultTwo++
			}
		}

	}

	fmt.Printf("result one: %d\n", resultOne)
	fmt.Printf("result two: %d\n", resultTwo)
}
