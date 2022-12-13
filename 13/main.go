package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Packages [][]byte

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	var packages Packages
	packages = append(packages, []byte("[[2]]"))
	packages = append(packages, []byte("[[6]]"))
	for s.Scan() {
		packages = append(packages, []byte(s.Text()))
		s.Scan()
		packages = append(packages, []byte(s.Text()))
		// empty line
		if s.Scan() {
			s.Text()
		}
	}
	sort.Sort(packages)
	result := 1
	for i, v := range packages {
		if string(v) == "[[2]]" || string(v) == "[[6]]" {
			result *= (i + 1)
		}
	}
	fmt.Println(result)
}

func (p Packages) Len() int      { return len(p) }
func (p Packages) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Packages) Less(i, j int) bool {

	//func less(left, right []byte) bool {
	pL := make([]byte, len(p[i]))
	pR := make([]byte, len(p[j]))
	for i, v := range p[i] {
		pL[i] = v
	}
	for i, v := range p[j] {
		pR[i] = v
	}
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
			if iR+1 >= len(pR) {
				fmt.Println("panic")
			}
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
			return true

		}
		if cmpIL > cmpIR {
			// fmt.Println("I: not in the right order")
			return false
		}
		if cmpCL == ']' && cmpCR != ']' {
			// fmt.Println("A: in the right order")
			return true

		}
		if cmpCL != ']' && cmpCR == ']' {
			// fmt.Println("A: not in the right order")
			return false
		}
	}
	return false
}
