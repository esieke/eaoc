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

	result := 0
	idx := 1
	for s.Scan() {
		// packet pairs
		pL := s.Bytes()
		s.Scan()
		pR := s.Bytes()
		// empty line
		if s.Scan() {
			s.Bytes()
		}
		// fmt.Println(string(pL))
		// fmt.Println(string(pR))

		iL := 0
		iR := 0
		for true {
			var cmpCL byte
			cmpIL := -1
			if iL >= len(pL) {
				break
			}
			if pL[iL] >= '0' && pL[iL] <= '9' {
				ns := []byte{pL[iL]}
				for ; pL[iL+1] >= '0' && pL[iL+1] <= '9'; iL++ {
					ns = append(ns, pL[iL+1])
				}
				cmpIL, _ = strconv.Atoi(string(ns))
			} else {
				cmpCL = pL[iL]
			}
			iL++

			var cmpCR byte
			cmpIR := -1
			if iR >= len(pR) {
				break
			}
			if pR[iR] >= '0' && pR[iR] <= '9' {
				ns := []byte{pR[iR]}
				for ; pR[iR+1] >= '0' && pR[iR+1] <= '9'; iR++ {
					ns = append(ns, pR[iR+1])
				}
				cmpIR, _ = strconv.Atoi(string(ns))
			} else {
				cmpCR = pR[iR]
			}
			iR++

			if cmpCL == '[' && cmpIR > -1 {
				p := []byte{'['}
				p = append(p, []byte(strconv.Itoa(cmpIR))...)
				p = append(p, []byte{']'}...)
				pR = append(p, pR[iR:]...)
				cmpCR = '['
				cmpIR = -1
				iR = 1
				// fmt.Printf("expand pR to: %s\n", string(pR))
			}
			if cmpIL > -1 && cmpCR == '[' {
				p := []byte{'['}
				p = append(p, []byte(strconv.Itoa(cmpIL))...)
				p = append(p, []byte{']'}...)
				pL = append(p, pL[iL:]...)
				cmpCL = '['
				cmpIL = -1
				iL = 1
				// fmt.Printf("expand pL to: %s\n", string(pL))
			}

			if cmpIL < cmpIR {
				// fmt.Println("I: in the right order")
				result += idx
				break
			}
			if cmpIL > cmpIR {
				// fmt.Println("I: not in the right order")
				break
			}
			if cmpCL == ']' && cmpCR != ']' {
				// fmt.Println("A: in the right order")
				result += idx
				break
			}
			if cmpCL != ']' && cmpCR == ']' {
				// fmt.Println("A: not in the right order")
				break
			}
		}
		idx++
	}
	fmt.Println(result)
}
